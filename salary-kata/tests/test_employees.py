import pytest
from fastapi.testclient import TestClient

from salarykata.main import app

client = TestClient(app)


def create_employee(**overrides):
    payload = {
        "full_name": "Alice Smith",
        "job_title": "Engineer",
        "country": "India",
        "salary": 100000.0,
        **overrides,
    }
    return client.post("/employees/", json=payload)


# ── CREATE ────────────────────────────────────────────────────────────────


class TestCreateEmployee:
    def test_create_employee_success(self):
        resp = create_employee()
        assert resp.status_code == 201

        data = resp.json()
        assert data["full_name"] == "Alice Smith"
        assert data["job_title"] == "Engineer"
        assert data["country"] == "India"
        assert data["salary"] == 100000.0
        assert isinstance(data["id"], int)

    def test_create_employee_assigns_unique_ids(self):
        id1 = create_employee(full_name="Alice").json()["id"]
        id2 = create_employee(full_name="Bob").json()["id"]
        assert id1 != id2

    @pytest.mark.parametrize(
        "missing_field,payload",
        [
            (
                "full_name",
                {"job_title": "Engineer", "country": "India", "salary": 50000},
            ),
            ("job_title", {"full_name": "Alice", "country": "India", "salary": 50000}),
            (
                "country",
                {"full_name": "Alice", "job_title": "Engineer", "salary": 50000},
            ),
            (
                "salary",
                {"full_name": "Alice", "job_title": "Engineer", "country": "India"},
            ),
        ],
    )
    def test_create_employee_missing_required_field(self, missing_field, payload):
        resp = client.post("/employees/", json=payload)
        assert resp.status_code == 422

        data = resp.json()
        assert "errors" in data
        assert missing_field in data["errors"]

    @pytest.mark.parametrize("salary", [-100, 0])
    def test_create_employee_invalid_salary(self, salary):
        resp = create_employee(salary=salary)
        assert resp.status_code == 422

        data = resp.json()
        assert "salary" in data["errors"]

    def test_create_employee_invalid_salary_type(self):
        resp = create_employee(salary="not-a-number")
        assert resp.status_code == 422

        data = resp.json()
        assert "salary" in data["errors"]

        error = data["errors"]["salary"][0]
        assert "message" in error
        assert "type" in error


# ── GET ───────────────────────────────────────────────────────────────────


class TestGetEmployee:
    def test_get_employee_success(self):
        created = create_employee().json()

        resp = client.get(f"/employees/{created['id']}")
        assert resp.status_code == 200

        assert resp.json() == created

    def test_get_employee_not_found(self):
        assert client.get("/employees/99999").status_code == 404


# ── LIST ──────────────────────────────────────────────────────────────────


class TestListEmployees:
    def test_list_employees_empty(self):
        resp = client.get("/employees/")

        assert resp.status_code == 200
        assert resp.json() == []

    def test_list_employees_returns_all(self):
        create_employee(full_name="Alice")
        create_employee(full_name="Bob")

        data = client.get("/employees/").json()
        names = {e["full_name"] for e in data}

        assert len(data) == 2
        assert names == {"Alice", "Bob"}


# ── UPDATE / PATCH ────────────────────────────────────────────────────────


class TestUpdateEmployee:
    def test_update_employee_success(self):
        emp_id = create_employee().json()["id"]

        resp = client.put(
            f"/employees/{emp_id}",
            json={
                "full_name": "Updated",
                "job_title": "Senior",
                "country": "India",
                "salary": 120000.0,
            },
        )

        assert resp.status_code == 200
        data = resp.json()

        assert data["full_name"] == "Updated"
        assert data["salary"] == 120000.0

    def test_update_employee_not_found(self):
        resp = client.put(
            "/employees/99999",
            json={
                "full_name": "X",
                "job_title": "Y",
                "country": "Z",
                "salary": 1000,
            },
        )
        assert resp.status_code == 404

    def test_patch_employee_partial_update(self):
        emp_id = create_employee().json()["id"]

        client.patch(f"/employees/{emp_id}", json={"salary": 999999.0})

        data = client.get(f"/employees/{emp_id}").json()

        assert data["salary"] == 999999.0
        assert data["full_name"] == "Alice Smith"

    def test_patch_rejects_invalid_data(self):
        emp_id = create_employee().json()["id"]

        resp = client.patch(f"/employees/{emp_id}", json={"full_name": 123})

        assert resp.status_code == 422

        data = client.get(f"/employees/{emp_id}").json()
        assert data["full_name"] == "Alice Smith"

    def test_patch_employee_not_found(self):
        assert (
            client.patch("/employees/99999", json={"salary": 1000}).status_code == 404
        )


# ── DELETE ────────────────────────────────────────────────────────────────


class TestDeleteEmployee:
    def test_delete_employee(self):
        emp_id = create_employee().json()["id"]

        assert client.delete(f"/employees/{emp_id}").status_code == 204
        assert client.get(f"/employees/{emp_id}").status_code == 404

    def test_delete_employee_not_found(self):
        assert client.delete("/employees/99999").status_code == 404

    def test_delete_idempotency(self):
        emp_id = create_employee().json()["id"]

        client.delete(f"/employees/{emp_id}")
        resp = client.delete(f"/employees/{emp_id}")

        assert resp.status_code == 404

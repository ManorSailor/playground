"""
Tests for Salary Calculation and Salary Metrics endpoints.
DB wiring lives in conftest.py.
"""

from fastapi.testclient import TestClient
from salarykata.main import app

client = TestClient(app)


def make_employee(
    full_name="Test User", job_title="Engineer", country="India", salary=100000.0
):
    return client.post(
        "/employees/",
        json={
            "full_name": full_name,
            "job_title": job_title,
            "country": country,
            "salary": salary,
        },
    ).json()


class TestSalaryCalculation:
    def test_india_tds_is_10_percent(self):
        emp = make_employee(country="India", salary=100000)
        data = client.get(f"/salaries/{emp['id']}").json()
        assert data["tds"] == 10000.0
        assert data["net_salary"] == 90000.0

    def test_india_gross_is_preserved(self):
        emp = make_employee(country="India", salary=80000)
        data = client.get(f"/salaries/{emp['id']}").json()
        assert data["gross_salary"] == 80000.0

    def test_us_tds_is_12_percent(self):
        emp = make_employee(country="United States", salary=100000)
        data = client.get(f"/salaries/{emp['id']}").json()
        assert data["tds"] == 12000.0
        assert data["net_salary"] == 88000.0

    def test_other_country_no_deductions(self):
        emp = make_employee(country="Germany", salary=100000)
        data = client.get(f"/salaries/{emp['id']}").json()
        assert data["tds"] == 0.0
        assert data["net_salary"] == 100000.0

    def test_another_other_country_no_deductions(self):
        emp = make_employee(country="Canada", salary=50000)
        data = client.get(f"/salaries/{emp['id']}").json()
        assert data["tds"] == 0.0
        assert data["net_salary"] == 50000.0

    def test_salary_calc_nonexistent_employee_returns_404(self):
        assert client.get("/employees/99999").status_code == 404

    def test_salary_calc_response_has_all_fields(self):
        emp = make_employee(country="India", salary=100000)
        data = client.get(f"/salaries/{emp['id']}").json()
        assert "gross_salary" in data
        assert "tds" in data
        assert "net_salary" in data

    def test_india_fractional_salary(self):
        emp = make_employee(country="India", salary=33333.33)
        data = client.get(f"/salaries/{emp['id']}").json()
        assert round(data["tds"], 2) == round(33333.33 * 0.10, 2)
        assert round(data["net_salary"], 2) == round(33333.33 * 0.90, 2)

    def test_us_fractional_salary(self):
        emp = make_employee(country="United States", salary=33333.33)
        data = client.get(f"/salaries/{emp['id']}").json()
        assert round(data["tds"], 2) == round(33333.33 * 0.12, 2)


class TestSalaryMetricsByCountry:
    def test_country_metrics_returns_200(self):
        make_employee(country="India", salary=50000)
        assert client.get("/salaries/metrics/country/India").status_code == 200

    def test_country_metrics_single_employee(self):
        make_employee(country="India", salary=50000)
        data = client.get("/salaries/metrics/country/India").json()
        assert data["min_salary"] == 50000.0
        assert data["max_salary"] == 50000.0
        assert data["avg_salary"] == 50000.0

    def test_country_metrics_multiple_employees(self):
        make_employee(country="India", salary=40000)
        make_employee(country="India", salary=60000)
        make_employee(country="India", salary=80000)
        data = client.get("/salaries/metrics/country/India").json()
        assert data["min_salary"] == 40000.0
        assert data["max_salary"] == 80000.0
        assert data["avg_salary"] == 60000.0

    def test_country_metrics_only_counts_given_country(self):
        make_employee(country="India", salary=50000)
        make_employee(country="United States", salary=200000)
        data = client.get("/salaries/metrics/country/India").json()
        assert data["max_salary"] == 50000.0

    def test_country_metrics_unknown_country_returns_404(self):
        assert client.get("/salaries/metrics/country/Narnia").status_code == 404

    def test_country_metrics_response_has_country_field(self):
        make_employee(country="India", salary=50000)
        data = client.get("/salaries/metrics/country/India").json()
        assert data["country"] == "India"

    def test_country_metrics_url_encoded_country(self):
        make_employee(country="United States", salary=100000)
        resp = client.get("/salaries/metrics/country/United%20States")
        assert resp.status_code == 200
        assert resp.json()["country"] == "United States"


class TestSalaryMetricsByJobTitle:
    def test_job_title_metrics_returns_200(self):
        make_employee(job_title="Engineer", salary=80000)
        assert client.get("/salaries/metrics/job-title/Engineer").status_code == 200

    def test_job_title_metrics_single_employee(self):
        make_employee(job_title="Engineer", salary=80000)
        data = client.get("/salaries/metrics/job-title/Engineer").json()
        assert data["avg_salary"] == 80000.0

    def test_job_title_metrics_multiple_employees(self):
        make_employee(job_title="Engineer", salary=60000)
        make_employee(job_title="Engineer", salary=80000)
        make_employee(job_title="Engineer", salary=100000)
        data = client.get("/salaries/metrics/job-title/Engineer").json()
        assert data["avg_salary"] == 80000.0

    def test_job_title_metrics_only_counts_given_title(self):
        make_employee(job_title="Engineer", salary=60000)
        make_employee(job_title="Manager", salary=200000)
        data = client.get("/salaries/metrics/job-title/Engineer").json()
        assert data["avg_salary"] == 60000.0

    def test_job_title_metrics_unknown_title_returns_404(self):
        assert client.get("/salaries/metrics/job-title/Wizard").status_code == 404

    def test_job_title_metrics_response_has_job_title_field(self):
        make_employee(job_title="Engineer", salary=80000)
        data = client.get("/salaries/metrics/job-title/Engineer").json()
        assert data["job_title"] == "Engineer"

    def test_job_title_metrics_url_encoded_title(self):
        make_employee(job_title="Senior Engineer", salary=120000)
        resp = client.get("/salaries/metrics/job-title/Senior%20Engineer")
        assert resp.status_code == 200
        assert resp.json()["job_title"] == "Senior Engineer"

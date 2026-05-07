# Salary Management API

A REST API for employee salary management built with FastAPI, SQLAlchemy 2.0, and SQLite.

## Stack

- **FastAPI** — routing, validation, auto-docs
- **SQLAlchemy 2.0** — ORM (`select()`-style queries throughout)
- **Pydantic v2** — request/response schemas, `422` handling is automatic
- **uv** — dependency management and venv
- **ruff** — linting and formatting
- **ty** — type checking

## Structure

```
src/salarykata/
├── main.py
├── db/
│   └── database.py          # engine, session factory, get_db
├── modules/
│   ├── employees/
│   │   ├── entity.py        # SQLAlchemy model
│   │   ├── schemas.py       # Pydantic schemas
│   │   └── service.py       # CRUD + Query APIs
│   └── salary/
│       ├── constants.py     # TDS rates by country
│       ├── schemas.py       # SalaryBreakdown, CountryMetrics, JobTitleMetrics
│       └── service.py       # salary business logic, depends on EmployeeService
├── routers/
│   ├── employees.py
│   └── salary.py            
└── shared/
    └── di.py                # Dependency Injection graph. Provided by FastAPI.

tests/
├── conftest.py              # StaticPool test engine, get_db override, reset_db fixture
├── test_employees.py
└── test_salary.py
```

## Setup

```bash
python3 -m venv .venv
pip install uv
uv sync
uv run pip install -e . # Installs the project in editable mode, allowing it to run locally
uv run uvicorn salarykata.main:app --reload
```

Docs available at `http://localhost:8000/docs`.

## Tests

```bash
uv run pytest
```

Tests use an isolated in-memory SQLite database. Each test gets a clean slate via the `reset_db` autouse fixture in `conftest.py`.

## API

### Employees

| Method | Path | Description |
|---|---|---|
| `POST` | `/employees/` | Create employee |
| `GET` | `/employees/` | List all |
| `GET` | `/employees/{id}` | Get one |
| `PUT` | `/employees/{id}` | Full update |
| `PATCH` | `/employees/{id}` | Partial update |
| `DELETE` | `/employees/{id}` | Delete |

All four fields are required on create: `full_name`, `job_title`, `country`, `salary` (must be positive).

### Salary

| Method | Path | Description |
|---|---|---|
| `GET` | `/salaries/{id}` | Gross, TDS, net for an employee |
| `GET` | `/salaries/metrics/country/{country}` | Min, max, avg salary for a country |
| `GET` | `/salaries/metrics/job-title/{job_title}` | Avg salary for a job title |

**TDS rates:**

| Country | Rate |
|---|---|
| India | 10% |
| United States | 12% |
| All others | 0% |

Both metric endpoints return `404` when no employees match.

## Architecture

Modular monolith. `employees` and `salary` are independent modules — neither imports from the other - except for Type Annotations.

`SalaryService` depends on `EmployeeReaderPort` (defined in `shared/di.py`), not on `EmployeeService` directly. `EmployeeService` satisfies the dependency via duck typing. The di (`shared/di.py`) is the only place that knows both modules exist.

This keeps business logic framework-agnostic: services raise `ValueError`, routers translate that to `HTTPException`. `Depends(get_db)` never appears inside a service.
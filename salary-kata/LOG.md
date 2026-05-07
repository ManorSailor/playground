# Devlog

## Architecture

Went with a **modular monolith**. Aware that this pattern pairs naturally with DDD, but the problem doesn't have enough complexity to justify aggregates, value objects, and the rest of it.

**Repositories were skipped** for the same reason — too much scaffolding for the current scope. That said, the `db` session is injected into services rather than used globally, so dropping repositories in later is straightforward.

## Modules

`salary` is a separate module even though it's currently just a `float` field on the employee. The reasoning: salary will grow. Bonuses, CTC components, country-specific taxation policies — these belong together and they carry real business logic. Keeping salary collapsed inside the employees module would make that harder to untangle later.

`SalaryService` depends on `EmployeeService` satisfies it via duck typing. The router is the only place that knows both modules exist and wires them together.

## Validation and Error Handling

FastAPI + Pydantic handle payload validation automatically. Invalid requests get a `422 Unprocessable Entity` back with field-level error detail — no manual validation code needed in the routes.

Added a global error handler for `ValidationError` on top of that, which formats errors consistently for clients. Makes it straightforward to render errors under form fields or as toasts on the frontend.

## Testing

E2E testing has been done. There are no unit or integrations tests. 
FUN FACT: If you are merely building an API, then your API is the _End_ for your consumers. Hence, we can very clearly label our tests as E2E here. If you're a boring person, yes, they are just HTTP API tests.

Tests run against an in-memory SQLite database using `StaticPool` so the schema is shared across connections. A `reset_db` autouse fixture in `conftest.py` creates tables before each test and drops them after, keeping tests fully isolated.

## Tooling

- **uv** — dependency management and virtualenv
- **ruff** — linting and formatting
- **ty** — type checking
- **claude** — scaffolding & tests

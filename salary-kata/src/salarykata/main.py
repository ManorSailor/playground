from collections import defaultdict

from fastapi import FastAPI, status
from fastapi.exceptions import RequestValidationError
from fastapi.responses import JSONResponse

from salarykata.db import Base, engine
from salarykata.routers import EmployeeRouter, SalaryRouter

# Create all tables on startup (idempotent)
Base.metadata.create_all(bind=engine)

app = FastAPI(
    title="Salary Management API",
    description="Employee CRUD, salary deduction calculation, and salary metrics.",
    version="0.1.0",
)


@app.exception_handler(exc_class_or_status_code=RequestValidationError)
async def global_validation_error_handler(_, err: RequestValidationError):
    """
    Restructures validations errors, making it easier for the client to consume them.
    ```json
    {
      "status": "error",
      "errors": {
        "full_name": [
          {
            "message": "String should have at least 1 character",
            "type": "string_too_short"
          }
        ],
        ...
    }
    ```

    Note: If you want to provide custom validation error messages. Look here: https://docs.pydantic.dev/latest/errors/errors/
    """
    reformatted = defaultdict(list)

    for err in err.errors():
        loc = err.get("loc")
        msg = err.get("msg")
        data_type = err.get("type")

        filtered_loc = loc[1:] if loc and loc[0] in ("body", "query", "path") else loc

        field = ".".join(str(x) for x in filtered_loc)
        reformatted[field].append(
            {
                "message": msg,
                "type": data_type,
            }
        )

    return JSONResponse(
        status_code=status.HTTP_422_UNPROCESSABLE_ENTITY,
        content={
            "status": "error",
            "errors": reformatted,
        },
    )


app.include_router(router=EmployeeRouter)
app.include_router(router=SalaryRouter)

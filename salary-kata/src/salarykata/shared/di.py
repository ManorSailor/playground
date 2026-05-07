from fastapi import Depends
from sqlalchemy.orm import Session

from salarykata.db import get_db
from salarykata.modules.employees import EmployeeService
from salarykata.modules.salary import SalaryService


def get_employee_service(
    db: Session = Depends(get_db),
) -> EmployeeService:
    return EmployeeService(db=db)


def get_salary_service(
    employee_service: EmployeeService = Depends(get_employee_service),
) -> SalaryService:
    return SalaryService(employee_service=employee_service)

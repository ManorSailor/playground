from sqlalchemy import func, select
from sqlalchemy.orm import Session

from .entity import EmployeeEntity
from .schemas import (
    CountrySalaryStats,
    CreateEmployee,
    Employee,
    PatchEmployee,
    UpdateEmployee,
)


class EmployeeService:
    """
    Merely a CRUD service. NOT A DOMAIN SERVICE!
    """

    def __init__(self, db: Session) -> None:
        self._db = db

    def create_employee(self, payload: CreateEmployee) -> Employee:
        employee = EmployeeEntity(**payload.model_dump())

        self._db.add(employee)
        self._db.commit()
        self._db.refresh(employee)

        return self._to_employee_dto(employee)

    def list_employees(self) -> list[Employee]:
        employee_entities = self._db.query(EmployeeEntity).all()

        return list(map(lambda emp: self._to_employee_dto(emp), employee_entities))

    def get_employee(self, employee_id: int) -> Employee | None:
        employee = self._db.get(EmployeeEntity, employee_id)

        if employee:
            return self._to_employee_dto(employee)

    def patch_employee(self, employee_id: int, payload: PatchEmployee) -> Employee:
        employee = self._db.get(EmployeeEntity, employee_id)

        if not employee:
            raise ValueError("Employee not found")

        for field, value in payload.model_dump(
            exclude_none=True, exclude_unset=True
        ).items():
            setattr(employee, field, value)

        self._db.commit()
        self._db.refresh(employee)

        return self._to_employee_dto(employee)

    def delete_employee(self, employee_id: int) -> None:
        employee = self._db.get(EmployeeEntity, employee_id)

        if not employee:
            raise ValueError("Employee not found")

        self._db.delete(employee)
        self._db.commit()

    def update_employee(self, employee_id: int, payload: UpdateEmployee) -> Employee:
        return self.patch_employee(employee_id=employee_id, payload=payload)  # ty:ignore[invalid-argument-type]

    def get_salary_stats_by_country(self, country: str) -> CountrySalaryStats | None:
        stmt = stmt = select(
            func.min(EmployeeEntity.salary).label("min_salary"),
            func.max(EmployeeEntity.salary).label("max_salary"),
            func.avg(EmployeeEntity.salary).label("avg_salary"),
        ).where(func.lower(EmployeeEntity.country) == country.lower())

        row = self._db.execute(statement=stmt).one()

        if row.min_salary is None:
            return None

        return CountrySalaryStats(
            country=country,
            min_salary=row.min_salary,
            max_salary=row.max_salary,
            avg_salary=row.avg_salary,
        )

    def get_avg_salary_by_job_title(self, job_title: str) -> float | None:
        stmt = stmt = select(func.avg(EmployeeEntity.salary).label("avg_salary")).where(
            func.lower(EmployeeEntity.job_title) == job_title.lower()
        )

        row = self._db.execute(statement=stmt).one()

        return row.avg_salary

    def _to_employee_dto(self, employee: EmployeeEntity) -> Employee:
        return Employee(
            id=employee.id,
            full_name=employee.full_name,
            country=employee.country,
            job_title=employee.job_title,
            salary=employee.salary,
        )

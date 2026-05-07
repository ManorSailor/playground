from typing import List

from fastapi import APIRouter, Depends, status, HTTPException

from salarykata.modules.employees import (
    CreateEmployee,
    Employee,
    EmployeeService,
    PatchEmployee,
    UpdateEmployee,
)
from salarykata.shared.di import get_employee_service

EmployeeRouter = APIRouter(prefix="/employees", tags=["employees"])


@EmployeeRouter.post("/", response_model=Employee, status_code=status.HTTP_201_CREATED)
def create_employee(
    payload: CreateEmployee,
    employee_service: EmployeeService = Depends(get_employee_service),
):
    return employee_service.create_employee(payload=payload)


@EmployeeRouter.get("/", response_model=List[Employee])
def list_employees(employee_service: EmployeeService = Depends(get_employee_service)):
    return employee_service.list_employees()


@EmployeeRouter.get("/{employee_id}", response_model=Employee)
def get_employee(
    employee_id: int, employee_service: EmployeeService = Depends(get_employee_service)
):
    employee = employee_service.get_employee(employee_id=employee_id)

    if not employee:
        raise HTTPException(status_code=404, detail="Employee not found")

    return employee


@EmployeeRouter.put("/{employee_id}", response_model=Employee)
def update_employee(
    employee_id: int,
    payload: UpdateEmployee,
    employee_service: EmployeeService = Depends(get_employee_service),
):
    try:
        updated_employee = employee_service.update_employee(
            employee_id=employee_id, payload=payload
        )
    except ValueError:
        raise HTTPException(status_code=404, detail="Employee not found")

    return updated_employee


@EmployeeRouter.patch("/{employee_id}", response_model=Employee)
def patch_employee(
    employee_id: int,
    payload: PatchEmployee,
    employee_service: EmployeeService = Depends(get_employee_service),
):
    try:
        updated_employee = employee_service.patch_employee(
            employee_id=employee_id, payload=payload
        )
    except ValueError:
        raise HTTPException(status_code=404, detail="Employee not found")

    return updated_employee


@EmployeeRouter.delete("/{employee_id}", status_code=status.HTTP_204_NO_CONTENT)
def delete_employee(
    employee_id: int, employee_service: EmployeeService = Depends(get_employee_service)
):
    try:
        employee_service.delete_employee(employee_id=employee_id)
    except ValueError:
        raise HTTPException(status_code=404, detail="Employee not found")

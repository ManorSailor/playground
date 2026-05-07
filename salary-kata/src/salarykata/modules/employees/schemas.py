from typing import Optional

from pydantic import BaseModel, Field


class CreateEmployee(BaseModel):
    full_name: str = Field(
        min_length=1, description="Name should have at least 1 character"
    )
    job_title: str = Field(
        min_length=1, description="Job Title should have at least 1 character"
    )
    country: str = Field(
        min_length=1, description="Country should have at least 1 character"
    )
    salary: float = Field(gt=0, description="Gross salary — must be a positive number")


class UpdateEmployee(BaseModel):
    full_name: str = Field(
        min_length=1, description="Name should have at least 1 character"
    )
    job_title: str = Field(
        min_length=1, description="Job Title should have at least 1 character"
    )
    country: str = Field(
        min_length=1, description="Country should have at least 1 character"
    )
    salary: float = Field(gt=0, description="Gross salary — must be a positive number")


class PatchEmployee(BaseModel):
    full_name: Optional[str] = Field(
        default=None, min_length=1, description="Name should have at least 1 character"
    )
    job_title: Optional[str] = Field(
        default=None,
        min_length=1,
        description="Job Title should have at least 1 character",
    )
    country: Optional[str] = Field(
        default=None,
        min_length=1,
        description="Country should have at least 1 character",
    )
    salary: Optional[float] = Field(
        default=None, gt=0, description="Name should have at least 1 character"
    )


class Employee(BaseModel):
    id: int
    full_name: str
    job_title: str
    country: str
    salary: float


class CountrySalaryStats(BaseModel):
    country: str
    min_salary: float
    max_salary: float
    avg_salary: float

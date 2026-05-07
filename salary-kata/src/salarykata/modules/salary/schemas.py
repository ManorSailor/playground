from pydantic import BaseModel


class SalaryBreakdown(BaseModel):
    employee_id: int
    gross_salary: float
    tds: float
    net_salary: float


class CountryMetrics(BaseModel):
    country: str
    min_salary: float
    max_salary: float
    avg_salary: float


class JobTitleMetrics(BaseModel):
    job_title: str
    avg_salary: float

from salarykata.modules.employees import EmployeeService

from .constants import get_tds_rate
from .schemas import (
    CountryMetrics,
    JobTitleMetrics,
    SalaryBreakdown,
)


class SalaryService:
    """
    Handles Salary related business logic. Not a Domain Service.

    We have created a separate module because salary won't be a float in the future.
    Additionally, the salary breakdown logic may also grow to include country specific policies. Taxation etc
    """

    def __init__(self, employee_service: EmployeeService) -> None:
        self._employee_service = employee_service

    def get_salary_breakdown(self, employee_id: int) -> SalaryBreakdown:
        employee = self._employee_service.get_employee(employee_id=employee_id)

        if not employee:
            raise ValueError("Employee not found")

        gross = employee.salary
        tds = round(gross * get_tds_rate(employee.country), 10)
        net = round(gross - tds, 10)

        return SalaryBreakdown(
            employee_id=employee_id,
            gross_salary=gross,
            tds=tds,
            net_salary=net,
        )

    def get_country_metrics(self, country: str) -> CountryMetrics:
        stats = self._employee_service.get_salary_stats_by_country(country=country)

        if stats is None:
            raise ValueError(f"No employees found for country '{country}'")

        return CountryMetrics(
            country=country,
            min_salary=stats.min_salary,
            max_salary=stats.max_salary,
            avg_salary=stats.avg_salary,
        )

    def get_job_title_metrics(self, job_title: str) -> JobTitleMetrics:
        avg_salary = self._employee_service.get_avg_salary_by_job_title(
            job_title=job_title
        )

        if avg_salary is None:
            raise ValueError(
                f"No employees found with job title '{job_title}'",
            )

        return JobTitleMetrics(
            job_title=job_title,
            avg_salary=avg_salary,
        )

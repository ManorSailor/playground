from fastapi import APIRouter, Depends, HTTPException, status

from salarykata.modules.salary import (
    CountryMetrics,
    JobTitleMetrics,
    SalaryBreakdown,
    SalaryService,
)
from salarykata.shared.di import get_salary_service

SalaryRouter = APIRouter(prefix="/salaries", tags=["salary"])


@SalaryRouter.get("/{employee_id}", response_model=SalaryBreakdown)
def get_salary_breakdown(
    employee_id: int, salary_service: SalaryService = Depends(get_salary_service)
):
    try:
        breakdown = salary_service.get_salary_breakdown(employee_id=employee_id)
    except ValueError as e:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail=str(e))

    return breakdown


@SalaryRouter.get("/metrics/country/{country}", response_model=CountryMetrics)
def get_country_metrics(
    country: str, salary_service: SalaryService = Depends(get_salary_service)
):
    try:
        metrics = salary_service.get_country_metrics(country=country)
    except ValueError as e:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail=str(e))

    return metrics


# ── Metrics — by Job Title ────────────────────────────────────────────────────


@SalaryRouter.get("/metrics/job-title/{job_title}", response_model=JobTitleMetrics)
def get_job_title_metrics(
    job_title: str, salary_service: SalaryService = Depends(get_salary_service)
):
    try:
        metrics = salary_service.get_job_title_metrics(job_title=job_title)
    except ValueError as e:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail=str(e))

    return metrics

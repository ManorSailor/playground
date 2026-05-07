# Deduction rates keyed by country (all-lowercase for case-insensitive lookup)
_TDS_RATES: dict[str, float] = {
    "india": 0.10,
    "united states": 0.12,
}


def get_tds_rate(country: str) -> float:
    return _TDS_RATES.get(country.lower(), 0.0)

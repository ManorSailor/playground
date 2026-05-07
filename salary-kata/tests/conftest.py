"""
Shared pytest configuration.

A single in-memory SQLite engine (via StaticPool) is used for the whole
test session.  The `reset_db` fixture (autouse) creates tables before each
test and tears them down after, guaranteeing a clean slate.
"""
import pytest
from fastapi.testclient import TestClient
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from sqlalchemy.pool import StaticPool

from salarykata.db import Base, get_db
from salarykata.main import app

# One shared in-memory database for all test modules
_ENGINE = create_engine(
    "sqlite://",
    connect_args={"check_same_thread": False},
    poolclass=StaticPool,
)
_SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=_ENGINE)


def _override_get_db():
    db = _SessionLocal()
    try:
        yield db
    finally:
        db.close()


# Wire the override once, at import time
app.dependency_overrides[get_db] = _override_get_db


@pytest.fixture(autouse=True)
def reset_db():
    """Recreate tables before every test; drop them after."""
    Base.metadata.create_all(bind=_ENGINE)
    yield
    Base.metadata.drop_all(bind=_ENGINE)


# Expose a ready-made TestClient so individual test modules can import it
@pytest.fixture(scope="session")
def client():
    return TestClient(app)

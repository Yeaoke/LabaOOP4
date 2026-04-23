from contextlib import contextmanager

from sqlalchemy import Engine, create_engine
from sqlalchemy.orm import declarative_base, sessionmaker


DATABASE_URL = "postgresql://postgres:MaL8033038@localhost:5432/OOP2FFF"
engine = create_engine(DATABASE_URL)
Base = declarative_base()

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
# Также как в java create-update
Base.metadata.create_all(engine)

@contextmanager
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()    
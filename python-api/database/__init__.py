from sqlalchemy import create_engine
from sqlalchemy.orm import DeclarativeBase, Session

class Base(DeclarativeBase):
    pass

DATABASE_URL = "postgresql://postgres:password@localhost:5432/OOP2FFF"

engine = create_engine(DATABASE_URL)

Base.metadata.create_all(engine)

with Session(engine) as session:
    
    session.add()
    session.commit()
from sqlalchemy import Column, String, BigInteger, ForeignKey
from sqlalchemy.orm import declarative_base, relationship
from sqlalchemy.ext.declarative import DeclarativeMeta
from sqlalchemy.dialects.postgresql import UUID
from abc import ABC, abstractmethod, ABCMeta
import uuid


class DeclarativeABCMeta(DeclarativeMeta, ABCMeta):
    pass


Base = declarative_base(metaclass=DeclarativeABCMeta)


class IndustrialCompanies(Base, ABC):
    __tablename__ = 'industrial_companies'

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    company_name = Column('company_name', String(100), nullable=False)
    annual_turnover = Column('annual_turnover', BigInteger, nullable=True)
    company_type = Column('company_type', String(20), nullable=False)

    holding_companies_id = Column(
        String(36),
        ForeignKey('holding_companies.id', ondelete='SET NULL'),
        nullable=True
    )

    holding_company = relationship(
        "HoldingCompanies",
        back_populates="industrial_companies",
        foreign_keys=[holding_companies_id]
    )

    __mapper_args__ = {
        'polymorphic_identity': 'IndustrialCompanies',
        'polymorphic_on': company_type
    }

    def __init__(self, company_name: str, annual_turnover: int = None,
                 company_type: str = None, holding_companies_id=None):
        self.company_name = company_name
        self.annual_turnover = annual_turnover
        self.company_type = company_type
        self.holding_companies_id = holding_companies_id

    @abstractmethod
    def cost_of_time_production(self):
        pass

    @abstractmethod
    def show_info(self):
        pass

    def _normalized_type(self):
        """Возвращает тип в формате, понятном Go-серверу: 'coal' или 'oil'"""
        mapping = {
            'CoalCompany': 'coal',
            'OilCompany': 'oil',
            'coal': 'coal',
            'oil': 'oil',
        }
        return mapping.get(self.company_type, self.company_type)

    def to_dict(self):
        return {
            "id": str(self.id) if self.id else None,
            "company_name": self.company_name,
            "annual_turnover": self.annual_turnover,
            "company_type": self._normalized_type(),
            "holding_companies_id": str(self.holding_companies_id) if self.holding_companies_id else None,
            "holding_name": self.holding_company.holding_name if self.holding_company else None,
            "coal_volume": getattr(self, "coal_volume", None),
            "mineCount": getattr(self, "mine_count", None),
            "coal_action": getattr(self, "coal_action", None),
            "oil_volume": getattr(self, "oil_volume", None),
            "holeCount": getattr(self, "hole_count", None),
            "oil_action": getattr(self, "oil_action", None),
        }

    def __repr__(self):
        return f"{self.id} {self.company_name} {self.company_type}"
from sqlalchemy import Column, Integer, BigInteger, String, ForeignKey
from sqlalchemy.orm import relationship
from models.industrial_companies import IndustrialCompanies, Base


class OilCompany(IndustrialCompanies):
    __tablename__ = 'oil_company'

    id = Column(
        Integer, 
        ForeignKey('industrial_companies.id', ondelete='CASCADE'), 
        primary_key=True
    )
    
    oil_volume = Column(BigInteger, nullable=True)
    hole_count = Column(BigInteger, nullable=True)
    oil_action = Column(String(255), nullable=True)

    __mapper_args__ = {
        'polymorphic_identity': 'OilCompany',
    }

    def __init__(self, company_name: str, annual_turnover: int = None, 
                 oil_volume: int = None, hole_count: int = None,
                 holding_companies_id: int = None):
        super().__init__(
            company_name, 
            annual_turnover, 
            company_type='OilCompany',
            holding_companies_id=holding_companies_id
        )
        self.oil_volume = oil_volume
        self.hole_count = hole_count
        self.oil_action = None

    def show_info(self):
        print("=== Нефтяная компания ===")
        print(f"ID: {self.id}")
        print(f"Название: {self.company_name}")
        print(f"Оборот: {self.annual_turnover}")
        print(f"Объём нефти: {self.oil_volume}")
        print(f"Скважин: {self.hole_count}")
        if self.oil_action:
            print(f"Действие: {self.oil_action}")

    def cost_of_time_production(self):
        if self.oil_volume and self.oil_volume > 0 and self.annual_turnover:
            counter = self.annual_turnover // self.oil_volume
            self.oil_action = f"Стоимость за единицу: {counter}"
            return counter
        return None

    def add_oil_wells(self, count: int):
        if count and count > 0:
            current = self.hole_count or 0
            self.hole_count = current + count
            self.oil_action = f"Добавлено скважин: {count}"

    def check_resources(self):
        self.oil_action = "Ресурсы проверены"

    def to_dict(self):
        data = super().to_dict()
        data.update({
            "oil_volume": self.oil_volume,
            "holeCount": self.hole_count,
            "oil_action": self.oil_action,
        })
        return data
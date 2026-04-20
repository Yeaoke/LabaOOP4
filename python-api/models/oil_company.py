from sqlalchemy import Column, ForeignKey, String, BigInteger, Integer
from sqlalchemy.orm import relationship
from industrial_companies import IndustrialCompanies, Base

class OilCompany(IndustrialCompanies):
    __tablename__ = 'oil_company'

    # Первичный ключ также является внешним ключом к родительской таблице
    id = Column(Integer, ForeignKey('industrial_companies.id'), primary_key=True)
    
    oil_volume = Column(BigInteger, nullable=True)
    hole_count = Column('well_count', BigInteger, nullable=True) # mapped to well_count in DB if needed, but attr is holeCount in Java logic
    oil_action = Column(String, nullable=True)

    __mapper_args__ = {
        'polymorphic_identity': 'OilCompany',
    }

    def __init__(self, company_name: str, annual_turnover: int = None, 
        oil_volume: int = None, hole_count: int = None):
        super().__init__(company_name, annual_turnover)
        self.oil_volume = oil_volume
        self.hole_count = hole_count
        self.oil_action = None

    def show_info(self):
        print("Нефтяная компания")
        print(f"ID: {self.id}")
        print(f"Название: {self.company_name}")
        print(f"Оборот: {self.annual_turnover}")
        print(f"Объём нефти: {self.oil_volume}")
        print(f"Скважин: {self.hole_count}")
        if self.oil_action:
            print(f"Действие: {self.oil_action}")

    def cost_of_time_production(self):
        if (self.oil_volume is not None and self.oil_volume > 0 
                and self.annual_turnover is not None):
            counter = self.annual_turnover // self.oil_volume
            print(f"Стоимость за единицу: {counter}")
            self.oil_action = f"Стоимость за единицу: {counter}"

    def add_oil_wells(self, count: int):
        if count is not None and count > 0:
            current_holes = self.hole_count if self.hole_count is not None else 0
            self.hole_count = current_holes + count
            print(f"Добавлено скважин: {count}")
            self.oil_action = f"Добавлено скважин: {count}"

    def check_resources(self):
        self.oil_action = "CHECKED"
        print("Ресурсы проверены")
        self.oil_action = "Ресурсы проверены"

    def __repr__(self):
        return super().__repr__()
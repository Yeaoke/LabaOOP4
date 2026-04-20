from dataclasses import dataclass
from typing import List, Optional
from models.industrial_companies import IndustrialCompanies
from models.coal_company import CoalCompany
from models.oil_company import OilCompany

class IndustrialCompamniesRepository:
    def add(self, industrialCompany: IndustrialCompanies) -> None: raise NotImplementedError
    def delete(self, industrialCompany: IndustrialCompanies) -> None: raise NotImplementedError
    def update(self, industrialCompany: IndustrialCompanies) -> None: raise NotImplementedError
    def gel_all(self) -> List[IndustrialCompanies]: raise NotImplementedError

    
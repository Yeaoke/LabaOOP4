package cache

import (
	dto "LabaOOP4/go-server/dto"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/golang-lru/v2/expirable"
)

var cache = expirable.NewLRU[uuid.UUID, dto.IndustrialCompanies](100, nil, time.Minute*10)

func CacheAdd(id uuid.UUID, model dto.IndustrialCompanies) {
	cache.Add(id, model)

	if company, ok := cache.Get(id); ok {
		log.Println(company)
	}
}

func CacheRemove(id uuid.UUID) {
	cache.Remove(id)

	company, ok := cache.Get(id)
	if ok {
		log.Fatal("Error with deleting cache", company)
	}
}

func CacheCheck(id uuid.UUID) {
	cache.Peek(id)

	company, ok := cache.Get(id)
	if ok {
		log.Fatal("Error with find cache", company)
	}
}

package strategy

import (
	"context"
	"fmt"
	"github.com/google/wire"
)

type GetUserStrategyFactory interface {
	GetStrategy(ctx context.Context, strategyType string) (GetUserStrategy, error)
}

type GetUserStrategyFactoryImpl struct {
	dbStrategy    *DB
	cacheStrategy *Cache
}

var (
	FactoryWireSet = wire.NewSet(NewGetUserStrategyFactoryImpl, wire.Bind(new(GetUserStrategyFactory), new(*GetUserStrategyFactoryImpl)))
)

func NewGetUserStrategyFactoryImpl(dbStrategy *DB, cacheStrategy *Cache) *GetUserStrategyFactoryImpl {
	return &GetUserStrategyFactoryImpl{dbStrategy: dbStrategy,
		cacheStrategy: cacheStrategy,
	}
}

func (g *GetUserStrategyFactoryImpl) GetStrategy(ctx context.Context, strategyType string) (GetUserStrategy, error) {
	switch strategyType {
	case "db":
		return g.dbStrategy, nil
	case "cache":
		return g.cacheStrategy, nil
	default:
		return nil, fmt.Errorf("no such strategy exist %v", strategyType)
	}
}

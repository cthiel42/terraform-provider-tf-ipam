package storage

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

type Pool struct {
	Name  string   `json:"name"`
	CIDRs []string `json:"cidrs"`
}

type Allocation struct {
	ID            string `json:"id"`
	PoolName      string `json:"pool_name"`
	AllocatedCIDR string `json:"allocated_cidr"`
	PrefixLength  int    `json:"prefix_length"`
}

type Storage interface {
	// pool operations
	GetPool(ctx context.Context, name string) (*Pool, error)
	ListPools(ctx context.Context) ([]Pool, error)
	SavePool(ctx context.Context, pool *Pool) error
	DeletePool(ctx context.Context, name string) error

	// allocation operations
	GetAllocation(ctx context.Context, id string) (*Allocation, error)
	ListAllocations(ctx context.Context) ([]Allocation, error)
	ListAllocationsByPool(ctx context.Context, poolName string) ([]Allocation, error)
	SaveAllocation(ctx context.Context, allocation *Allocation) error
	DeleteAllocation(ctx context.Context, id string) error

	Close() error
}

type Config struct {
	Type string // "file", "azure", "dynamodb", etc.

	// File backend config
	FilePath string

	// Future: Azure Storage Table config
	// AzureAccountName string
	// AzureAccountKey  string
	// AzureTableName   string

	// Future: DynamoDB config
	// DynamoDBTableName string
	// DynamoDBRegion    string
}

func Factory(ctx context.Context, config *Config) (Storage, error) {
	switch config.Type {
	case "file", "": // default to file
		return NewFileStorage(config.FilePath)
	// Future backends:
	// case "azure":
	//     return NewAzureStorage(config)
	// case "dynamodb":
	//     return NewDynamoDBStorage(config)
	default:
		return nil, errors.New("unknown storage type")
	}
}

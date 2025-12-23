package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type FileStorage struct {
	filePath string
	mu       sync.RWMutex
	data     *fileData
}

type fileData struct {
	Pools       map[string]*Pool       `json:"pools"`
	Allocations map[string]*Allocation `json:"allocations"`
}

// Most methods make copies of data to avoid external mutation issues

func NewFileStorage(filePath string) (*FileStorage, error) {
	if filePath == "" {
		// default to .terraform directory in current working directory
		cwd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("failed to get working directory: %w", err)
		}
		terraformDir := filepath.Join(cwd, ".terraform")
		if err := os.MkdirAll(terraformDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create .terraform directory: %w", err)
		}
		filePath = filepath.Join(terraformDir, "ipam-storage.json")
	}

	fs := &FileStorage{
		filePath: filePath,
		data: &fileData{
			Pools:       make(map[string]*Pool),
			Allocations: make(map[string]*Allocation),
		},
	}

	// check if file already exists
	if err := fs.load(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to load storage file: %w", err)
	}

	return fs, nil
}

// reads storage from disk
func (fs *FileStorage) load() error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	data, err := os.ReadFile(fs.filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, fs.data)
}

// write storage to disk
func (fs *FileStorage) save() error {
	// make directory if it doesnt exist
	dir := filepath.Dir(fs.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	data, err := json.MarshalIndent(fs.data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal storage data: %w", err)
	}

	// Write to tmp file first, then rename for atomicity
	tempFile := fs.filePath + ".tmp"
	if err := os.WriteFile(tempFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write storage file: %w", err)
	}

	if err := os.Rename(tempFile, fs.filePath); err != nil {
		os.Remove(tempFile) // cleanup tmp file on error
		return fmt.Errorf("failed to rename storage file: %w", err)
	}

	return nil
}

func (fs *FileStorage) GetPool(ctx context.Context, name string) (*Pool, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	pool, exists := fs.data.Pools[name]
	if !exists {
		return nil, ErrNotFound
	}

	// returns copy
	poolCopy := *pool
	return &poolCopy, nil
}

func (fs *FileStorage) ListPools(ctx context.Context) ([]Pool, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	// return copies
	pools := make([]Pool, 0, len(fs.data.Pools))
	for _, pool := range fs.data.Pools {
		pools = append(pools, *pool)
	}

	return pools, nil
}

func (fs *FileStorage) SavePool(ctx context.Context, pool *Pool) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	// make a copy to store
	poolCopy := *pool
	fs.data.Pools[pool.Name] = &poolCopy

	return fs.save()
}

func (fs *FileStorage) DeletePool(ctx context.Context, name string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	if _, exists := fs.data.Pools[name]; !exists {
		return ErrNotFound
	}

	delete(fs.data.Pools, name)
	return fs.save()
}

func (fs *FileStorage) GetAllocation(ctx context.Context, id string) (*Allocation, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	allocation, exists := fs.data.Allocations[id]
	if !exists {
		return nil, ErrNotFound
	}

	// Return copy
	allocCopy := *allocation
	return &allocCopy, nil
}

func (fs *FileStorage) ListAllocations(ctx context.Context) ([]Allocation, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	// return copies
	allocations := make([]Allocation, 0, len(fs.data.Allocations))
	for _, alloc := range fs.data.Allocations {
		allocations = append(allocations, *alloc)
	}

	return allocations, nil
}

func (fs *FileStorage) ListAllocationsByPool(ctx context.Context, poolName string) ([]Allocation, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	allocations := make([]Allocation, 0)
	for _, alloc := range fs.data.Allocations {
		if alloc.PoolName == poolName {
			allocations = append(allocations, *alloc)
		}
	}

	return allocations, nil
}

func (fs *FileStorage) SaveAllocation(ctx context.Context, allocation *Allocation) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	allocCopy := *allocation
	fs.data.Allocations[allocation.ID] = &allocCopy

	return fs.save()
}

func (fs *FileStorage) DeleteAllocation(ctx context.Context, id string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	if _, exists := fs.data.Allocations[id]; !exists {
		return ErrNotFound
	}

	delete(fs.data.Allocations, id)
	return fs.save()
}

func (fs *FileStorage) Close() error {
	// file storage doesn't need any cleanup
	return nil
}

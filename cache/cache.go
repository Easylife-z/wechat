package cache

import (
	"context"
	"time"
)

type Cache interface {
	Set(key string, val interface{}, timeout time.Duration) error
	Get(key string) interface{}
	IsExist(key string) bool
	Delete(key string) error
}

// ContextCache interface
type ContextCache interface {
	Cache
	GetContext(ctx context.Context, key string) interface{}
	SetContext(ctx context.Context, key string, val interface{}, timeout time.Duration) error
	IsExistContext(ctx context.Context, key string) bool
	DeleteContext(ctx context.Context, key string) error
}

// GetContext get value from cache
func GetContext(ctx context.Context, cache Cache, key string) interface{} {
	if cache, ok := cache.(ContextCache); ok {
		return cache.GetContext(ctx, key)
	}
	return cache.Get(key)
}

// SetContext set value to cache
func SetContext(ctx context.Context, cache Cache, key string, val interface{}, timeout time.Duration) error {
	if cache, ok := cache.(ContextCache); ok {
		return cache.SetContext(ctx, key, val, timeout)
	}
	return cache.Set(key, val, timeout)
}

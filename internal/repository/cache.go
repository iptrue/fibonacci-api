package repository

// Cache - интерфейс для кэширования
type Cache interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	Exists(key string) (bool, error)
	Clear() error
	GetWithFallback(key string) (string, error)
}

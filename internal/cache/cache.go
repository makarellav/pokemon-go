package cache

type Cache interface {
	Get(key string) ([]byte, bool)
	Add(key string, data []byte)
}

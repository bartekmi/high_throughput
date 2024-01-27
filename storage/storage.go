package storage

type KVPair struct {
	ID      string
	Title   string
	Content string
}

type Storage interface {
	Write(data KVPair) error
	Read(id string) (KVPair, bool, error)
}

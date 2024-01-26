package storage

type KVPair struct {
	ID      string
	Title   string
	Content string
}

type Storage interface {
	Write(content KVPair) error
	Read(guid string) (KVPair, bool, error)
}

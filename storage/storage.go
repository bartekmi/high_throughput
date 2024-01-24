package storage

type KVPair struct {
	Guid    string
	Content string
}

type Storage interface {
	Write(guid, content string) error
	Read(guid string) (KVPair, bool, error)
}

package storage

type KVPair struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Storage interface {
	Write(data KVPair) error
	Read(id string) (KVPair, bool, error)
	Count() (int, error)
}

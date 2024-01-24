package storage

type dummy struct {
	data map[string]string
}

func NewDummy() *dummy {
	return &dummy{data: make(map[string]string)}
}

func (sd *dummy) Write(key, content string) {
	sd.data[key] = content
}

func (sd *dummy) Read(key string) (KVPair, bool) {
	content, ok := sd.data[key]
	if ok {
		return KVPair{Guid: key, Content: content}, true
	}
	return KVPair{}, false
}

func (sd *dummy) Count() int {
	return len(sd.data)
}

package storage

import "fmt"

type dummy struct {
	data  map[string]KVPair
	Error string
}

func NewDummy() *dummy {
	return &dummy{data: make(map[string]KVPair)}
}

func (sd *dummy) Write(data KVPair) error {
	if sd.Error != "" {
		return fmt.Errorf(sd.Error)
	}

	fmt.Printf("Dummy Storage: %s => %s\n", data.ID, data.Content)
	sd.data[data.ID] = data
	return nil
}

func (sd *dummy) Read(key string) (KVPair, bool, error) {
	if sd.Error != "" {
		return KVPair{}, false, fmt.Errorf(sd.Error)
	}

	content, ok := sd.data[key]
	if ok {
		fmt.Printf("Dummy Storage: %s => %s\n", key, content)
		return content, true, nil
	}
	fmt.Printf("Dummy Storage: %s does not exist\n", key)
	return KVPair{}, false, nil
}

func (sd *dummy) Count() int {
	return len(sd.data)
}

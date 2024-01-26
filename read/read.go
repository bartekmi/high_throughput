package main

import (
	"snippetstore/storage"
)

type Reader struct {
	storage storage.Storage
}

func New(s storage.Storage) *Reader {
	return &Reader{storage: s}
}

func (r *Reader) Read(key string) (storage.KVPair, bool, error) {
	return r.storage.Read(key)
}

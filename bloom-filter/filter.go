package main

import (
	"hash/maphash"
)

type BloomFilter struct {
	bitSize uint64
	bits    []bool
	hasher  *maphash.Hash
}

func NewBloomFilter(bitSize uint64) *BloomFilter {
	return &BloomFilter{
		bitSize: bitSize,
		hasher:  &maphash.Hash{},
		bits:    make([]bool, bitSize),
	}
}

func (filter *BloomFilter) getIndex(value []byte) uint64 {
	filter.hasher.Reset()
	filter.hasher.Write(value)

	hash := filter.hasher.Sum64()

	return hash % filter.bitSize
}

func (filter *BloomFilter) Insert(value []byte) {
	index := filter.getIndex(value)
	filter.bits[index] = true
}

func (filter *BloomFilter) Check(value []byte) bool {
	index := filter.getIndex(value)
	return filter.bits[index]
}

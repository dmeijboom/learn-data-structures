package main

import "fmt"

func main() {
	filter := NewBloomFilter(100)
	keys := []string{"whale", "dog", "giant-blue-lion"}
	keysNotInFilter := []string{"cat", "small-red-panda"}

	for _, key := range keys {
		filter.Insert([]byte(key))
	}

	for _, key := range keys {
		if filter.Check([]byte(key)) {
			fmt.Printf("ok: %s matches\n", key)
		} else {
			fmt.Printf("fail: %s doesn't match\n", key)
		}
	}

	for _, key := range keysNotInFilter {
		if filter.Check([]byte(key)) {
			fmt.Printf("fail: %s matches\n", key)
		} else {
			fmt.Printf("ok: %s doesn't match\n", key)
		}
	}
}
package main

import (
	"go-cache/internal"
	"log"
)

func main() {
	log.Println("starting the cache")

	cache := internal.NewCacheCreation()

	for _, word := range []string{"abc", "bed", "fdsfs"} {
		cache.Check(word)
		cache.Display()
	}

}

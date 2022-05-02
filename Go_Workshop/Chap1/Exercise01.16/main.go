package main

import "fmt"

// const declare
const GlobalLimit = 100
const MaxCacheSzie = GlobalLimit * 2
const (
	KeyBook = "book_"
	KeyCD   = "cd_"
)

// cache is key-value pair
var cache map[string]string

func getCache(key string) string {
	return cache[key]
}

func setCache(key string, val string) {
	if len(cache) > MaxCacheSzie {
		// cache is full ... do nothing
		return
	}
	cache[key] = val
}

func setBook(isbn string, name string) {
	key := KeyBook + isbn
	setCache(key, name)
}

func getBook(isbn string) string {
	key := KeyBook + isbn
	return getCache(key)
}

func setCD(sku string, name string) {
	key := KeyCD + sku
	setCache(key, name)
}

func getCD(sku string) string {
	key := KeyCD + sku
	return getCache(key)
}

func main() {
	fmt.Println("GlobalLimit:", GlobalLimit)
	fmt.Println("MaxCacheSzie:", MaxCacheSzie)
	fmt.Println("KeyBook:", KeyBook)
	fmt.Println("KeyCD:", KeyCD)

	cache = make(map[string]string)

	// set data in cache
	setBook("123456", "Ready to Go")
	setCD("123456", "Fly to the moon")

	fmt.Println("Book:", getBook("123456"))
	fmt.Println("CD:", getCD("123456"))
}

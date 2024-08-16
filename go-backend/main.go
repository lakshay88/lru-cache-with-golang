package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"myproject/cache" // Import the cache package correctly
)

var lruCache *cache.LRUCache

// Enable cors
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	duration, _ := strconv.Atoi(r.URL.Query().Get("duration"))

	lruCache.Set(key, value, time.Duration(duration)*time.Second)
	w.Write([]byte("Set Success"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	key := r.URL.Query().Get("key")
	value, found := lruCache.Get(key)

	if found {
		json.NewEncoder(w).Encode(map[string]string{"value": value})
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Key Not Found or Expired"))
	}
}

func main() {
	lruCache = cache.NewLRUCache(1024)

	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/get", getHandler)

	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	fmt.Println("Server running at :1234")
	http.ListenAndServe(":1234", nil)
}

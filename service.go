package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<body> %s</body>", "This is a test")
}

func initRedisContent() {

}

func createRedisPool() *redis.Pool {
	return redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", "0.0.0.0:6379")
		if err != nil {
			return nil, err
		}

		return c, err
	}, 10)
}

func main() {
	log.Println("Starting service on port 80")

	redisPool = createRedisPool()
	defer redisPool.Close()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

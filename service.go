package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

func handler(w http.ResponseWriter, r *http.Request) {
	content := getContent()
	fmt.Fprintf(w, "<body>content: %s</body>", content)
}

func initRedisContent() {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", "content", "this is some content")

	if err != nil {
		log.Println(err)
	}
}

func getContent() string {
	conn := redisPool.Get()
	defer conn.Close()

	result, err := redis.String(conn.Do("GET", "content"))
	if err != nil {
		log.Println(err)
	}
	return result
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

	log.Println("Initiating content")
	initRedisContent()

	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:80", nil)
}

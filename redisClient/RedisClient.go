package redisClient

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

type RedisClient struct {
	network 	string
	address 	string
	conn 	redis.Conn
}

func (c *RedisClient) ConnectionToServer() {
	// Connect to the Redis server (default port is 6379)
	conn, err := redis.Dial("tcp", "localhost:6379")
	c.conn
	if err != nil {
		log.Fatal(err)
	}
}

func main()  {

	// Connect to the Redis server (default port is 6379)
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	// Ensure the connection is always closed before exiting
	err
	defer conn.Close()

	r, err := conn.Do("SET", "A", "ONE")
	if err != nil {
		log.Fatal(err)
	}

	r, err = conn.Do("GET", "A")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("r : %s (type %T) \n", r, r)

	r, err = redis.String(conn.Do("GET", "A"))

}
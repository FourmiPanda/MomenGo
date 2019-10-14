package redisMqtt

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"myproject/internal/entities"
)

type RedisClient struct {
	network string
	address string
	conn    redis.Conn
}

func CreateARedisClient(network string, address string) RedisClient {
	return RedisClient{network: network, address: address, conn: nil}
}

func (c *RedisClient) connectionToServer() redis.Conn {
	// Connect to the Redis server (default port is 6379)
	// conn, err := redis.Dial("tcp", "localhost:6379")
	var err error
	c.conn, err = redis.Dial(c.network, c.address)
	if err != nil {
		log.Fatal(err)
	}
	// Ensure the connection is always closed before exiting
	//defer c.conn.Close()
	return c.conn
}

func (c *RedisClient) CreateAnEntry(entry *entities.RedisEntry) *redis.Conn {
	c.connectionToServer()
	_, err := c.connectionToServer().Do("RPUSH", "captor", entry.RedisEntryTOString())
	if err != nil {
		log.Fatal(err)
	}
	//defer c.conn.Close()
	defer c.conn.Close()
	return &c.conn
}

func (c *RedisClient) GetAnEntry(entry string) string {
	r, err := redis.String(c.conn.Do("GET", entry))
	if err != nil {
		log.Fatal(err)
	}
	defer c.conn.Close()
	return r
}

func main() {

}

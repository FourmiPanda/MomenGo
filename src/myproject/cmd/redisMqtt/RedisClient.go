package redisMqtt

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"myproject/internal/entities"
	"strings"
)

type RedisClient struct {
	config 	entities.RedisDB
	conn    redis.Conn
}

func CreateARedisClient(network string, address string) *RedisClient {
	return &RedisClient{config: entities.RedisDB{Network: network, Address: address}, conn: nil}
}
func CreateARedisClientFromConfig(config entities.RedisDB) *RedisClient {
	return &RedisClient{config: config, conn: nil}
}

func (r *RedisClient) connectionToServer() redis.Conn {
	// Connect to the Redis server (default port is 6379)
	// conn, err := redis.Dial("tcp", "localhost:6379")
	var err error
	r.conn, err = redis.Dial(r.config.Network, r.config.Address)
	if err != nil {
		log.Fatal(err)
	}
	// Ensure the connection is always closed before exiting
	//defer c.conn.Close()
	return r.conn
}
func (r *RedisClient) doesKeysExists (tabKeys []string) bool{
	res, err := redis.Bool(
		r.connectionToServer().Do("EXISTS", strings.Join(tabKeys, " ")))
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (r *RedisClient) AddCaptorEntryToDB (entry *entities.RedisEntry)  {
	key := []string{entry.CaptorKey()}
	values := entry.GetCaptorValues()
	r.connectionToServer()
	defer r.conn.Close()
	// Ensure there is already a table for this Captor and this CaptorValues
	if r.doesKeysExists(key) {
		for i := 0; i < len(values); i++ {
			//For DEBUG purpose
			//println("DEBUG :","RPUSH", entry.CaptorValuesKey(), values[i])
			_, err := r.connectionToServer().Do("RPUSH", entry.CaptorValuesKey(), values[i])
			if err != nil {
				log.Fatal(err)
			}
		}
	} else if entry.Captor.IsEmpty(){
		log.Fatal("WARNING : There is no value in this RedisEntry")
	} else {
		//For DEBUG purpose
		//println("DEBUG :","HMSET", entry.CaptorKey(),
		//	"idAirport", entry.Captor.GetIdAirportToString(),
		//	"idCaptor", entry.Captor.GetIdCaptorToString(),
		//	"measure", entry.Captor.GetMeasureToString(),
		//	"day", "lol",
		//	"values", entry.CaptorValuesKey())
		_, err := r.connectionToServer().Do("HMSET", entry.CaptorKey(),
			"idAirport", entry.Captor.GetIdAirportToString(),
			"idCaptor", entry.Captor.GetIdCaptorToString(),
			"measure", entry.Captor.GetMeasureToString(),
			"day", "lol",
			"values", entry.CaptorValuesKey())
		if err != nil {
			log.Fatal(err)
		}
		r.AddCaptorEntryToDB(entry)
	}
}
//func (r *RedisClient) CreateAnEntry(entry *entities.RedisEntry) *redis.Conn {
//	r.connectionToServer()
//	// Ensure there is already an entry for this Captor
//	_, err := r.connectionToServer().Do("RPUSH", "captor", entry.RedisEntryToString())
//	if err != nil {
//		log.Fatal(err)
//	}
//	//defer r.conn.Close()
//	defer r.conn.Close()
//	return &r.conn
//}
//
func (r *RedisClient) GetACaptorAttributeEntry(key string, value string) string {
	r.connectionToServer()
	res, err := redis.Strings(r.conn.Do("HMGET", key, value))
	if err != nil {
		log.Fatal(err)
	}
	defer r.conn.Close()
	return res[0]
}

func (r *RedisClient) GetAllCaptorValuesEntry(key string) []string {
	r.connectionToServer()
	res, err := redis.Strings(r.conn.Do("LRANGE", key, "0","-1"))
	if err != nil {
		log.Fatal(err)
	}
	defer r.conn.Close()
	return res
}

func main() {

}

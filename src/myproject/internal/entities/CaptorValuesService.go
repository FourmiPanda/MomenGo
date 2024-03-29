/**
 * CaptorValues Service
 *
 * @description ::  Manage CaptorValues.
 */

package entities

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
	"strings"
	"time"
)

type RedisClient struct {
	config RedisDB
	conn   redis.Conn
}

func CreateARedisClient(network string, address string) *RedisClient {
	return &RedisClient{config: RedisDB{Network: network, Address: address}, conn: nil}
}

func CreateARedisClientFromConfig(config *Configuration) *RedisClient {
	return &RedisClient{config: config.Redis, conn: nil}
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

func (r *RedisClient) doesKeysExists(tabKeys []string) bool {
	res, err := redis.Bool(
		r.connectionToServer().Do("EXISTS", strings.Join(tabKeys, " ")))
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (r *RedisClient) AddCaptorEntryToDB(entry *RedisEntry) error {
	values := entry.GetCaptorValues()
	r.connectionToServer()
	defer r.conn.Close()
	// Ensure there is already a hashes for this Captor and this CaptorValues
	for i := 0; i < len(values); i++ {
		//For DEBUG purpose
		//println("DEBUG :",
		//	"ZADD",
		//	entry.CaptorKey(),
		//	entry.GetDayDateAsInt(i),
		//	entry.GetDayDate(i))
		res, err := r.connectionToServer().Do("ZADD", entry.CaptorKey(),
			entry.GetDayDateAsInt(i),
			entry.GetDayDate(i))
		if err != nil {
			return err
		}
		if res.(int64) == 1 {
			log.Println("New entry added to the Captor sorted set", entry.RedisEntryToString())
		}
		//For DEBUG purpose
		//println("DEBUG :",
		//	"ZADD",
		//	entry.CaptorValuesKey(i),
		//	entry.GetTimestampAsInt(i),
		//	entry.GetCaptorValueAsJson(i))
		res, err = r.connectionToServer().Do("ZADD", entry.CaptorValuesKey(i),
			entry.GetTimestampAsInt(i),
			entry.GetCaptorValueAsJson(i))
		if err != nil {
			return err
		}
		if res.(int64) == 1 {
			log.Println("New entry added to the CaptorValues sorted set", values[i])
		}
	}
	return nil
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

func (r *RedisClient) GetCaptorValuesKeysInInterval(keys []string, start time.Time, end time.Time) ([]string, error) {
	r.connectionToServer()
	var res []string
	var err error
	for i := 0; i < len(keys); i++ {
		res, err = r.GetACaptorValuesKeyInInterval(keys[i], start, end)
		if err != nil {
			break
		}
	}
	return res, err
}

func (r *RedisClient) GetACaptorValuesKeyInInterval(key string, start time.Time, end time.Time) ([]string, error) {
	r.connectionToServer()
	res, err := redis.Strings(r.conn.Do("ZRANGEBYSCORE", key, start.Unix(), end.Unix()))
	if err != nil {
		log.Println(err)
	}
	defer r.conn.Close()
	s := strings.Split(key, ":")
	s[0] = "CaptorValues"
	key = strings.Join(s, ":")
	for i := 0; i < len(res); i++ {
		res[i] = key + ":" + res[i]
	}
	return res, err
}

func (r *RedisClient) GetAllCaptorValuesOfPresForADay(dayDate time.Time) ([]Captor, error) {
	return r.GetAllCaptorValuesOfMeasureForADay("PRES", dayDate)
}

func (r *RedisClient) GetAllCaptorValuesOfWindForADay(dayDate time.Time) ([]Captor, error) {
	return r.GetAllCaptorValuesOfMeasureForADay("WIND", dayDate)
}

func (r *RedisClient) GetAllCaptorValuesOfTempForADay(dayDate time.Time) ([]Captor, error) {
	return r.GetAllCaptorValuesOfMeasureForADay("TEMP", dayDate)
}

func (r *RedisClient) GetAllCaptorValuesOfMeasureForADay(measure string, dayDate time.Time) ([]Captor, error) {
	var res []Captor

	r.connectionToServer()

	y := strconv.Itoa(dayDate.Year())
	m := strconv.Itoa(int(dayDate.Month()))
	d := strconv.Itoa(dayDate.Day())

	keys, err := redis.Strings(r.conn.Do("keys", "*:"+measure+":*:"+y+":"+m+":"+d))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, key := range keys {
		q, err2 := r.GetAllCaptorValuesOfADay(key)
		if err2 != nil {
			log.Println(err2)
			continue
		}
		res = append(res, q)
	}

	res = MergeEqualsCaptors(res)

	return res, nil
}
func (r *RedisClient) GetAllCaptorValuesOfATypeInInterval(measure string, startDate time.Time,endDate time.Time) ([]Captor, error) {
	var res []Captor

	r.connectionToServer()
	defer r.conn.Close()

	// Get all keys for the "measure" type
	keysCaptors, err := redis.Strings(r.conn.Do("keys", "Captor:*:"+measure+":*"))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, keyCaptor := range keysCaptors {
		// keyCaptor = Captor[0]:IATA[1]:MEASURE[2]:ID[3]
		data := strings.Split(keyCaptor, ":")
		idCaptor, errId:= strconv.Atoi(data[3])
		if errId != nil {
			log.Println(err)
			continue
		}

		// Create a Captor from the key
		c := Captor{
			IdCaptor:  idCaptor,
			IdAirport: data[1],
			Measure:   data[2],
			Values:    nil,
		}

		// Get all the keys of the days with recorded values between "startDate" and "endDate"
		keysValues, err := redis.Strings(r.conn.Do("ZRANGEBYSCORE", keyCaptor, startDate.Unix(), endDate.Unix()))
		if err != nil {
			log.Println(err)
			continue
		}

		for _, keyValue := range keysValues {
			// keyValue = Captor[0]:IATA[1]:MEASURE[2]:ID[3]:YYYY[4]:MM[5]:DD[6]
			// Get all the values of this day between "startDate" and "endDate"
			currentKey := "CaptorValues:"+data[1]+":"+data[2]+":"+data[3]+":"+keyValue
			j, err := redis.ByteSlices(r.conn.Do("ZRANGEBYSCORE", currentKey, startDate.Unix(), endDate.Unix()))
			if err != nil {
				log.Println(err)
				continue
			}
			for _, i := range j {
				c.AddValuesFromJson(i)
			}
		}

		res = append(res, c)

	}

	return res, nil
}

func (r *RedisClient) GetACaptorValuesEntriesInInterval(key string, start time.Time, end time.Time) ([]string, error) {
	r.connectionToServer()
	res, err := redis.Strings(r.conn.Do("ZRANGEBYSCORE", key, start.Unix(), end.Unix()))
	if err != nil {
		log.Println(err)
	}
	return res, err
}

func (r *RedisClient) GetCaptorValuesEntriesInInterval(key string, start time.Time, end time.Time) ([]string, error) {
	r.connectionToServer()
	res, err := redis.Strings(r.conn.Do("ZRANGEBYSCORE", key, start.Unix(), end.Unix()))
	if err != nil {
		log.Println(err)
	}
	return res, err
}

func (r *RedisClient) GetAllCaptorValuesOfADay(key string) (Captor, error) {
	r.connectionToServer()
	q, err := redis.ByteSlices(r.conn.Do("ZRANGE", key, "0", "-1"))
	if err != nil {
		log.Println(err)
	}
	defer r.conn.Close()
	k := strings.Split(key, ":")
	idCaptor, err := strconv.Atoi(k[3])
	idAirport := k[1]
	measure := k[2]

	if err != nil {
		log.Println(err)
		return Captor{}, err
	}

	res := Captor{
		IdCaptor:  idCaptor,
		IdAirport: idAirport,
		Measure:   measure,
		Values:    nil,
	}

	for _, p := range q {
		_, err = res.AddValuesFromJson(p)
		if err != nil {
			log.Println(err)
			return Captor{}, err
		}
	}

	return res, err
}

func (r *RedisClient) Find(query string) ([]Captor, error) {
	var res []Captor

	r.connectionToServer()

	keys, err := redis.Strings(r.conn.Do("keys", query))
	if err != nil {
		return nil, err
	}

	fmt.Println(keys)

	for _, key := range keys {
		q, err2 := r.GetAllCaptorValuesOfADay(key)
		if err2 != nil {
			log.Println(err2)
			continue
		}
		res = append(res, q)
	}

	res = MergeEqualsCaptors(res)

	return res, nil
}

package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

var client *redis.Client

func main() {
	var err error
	var index int
	var port int
	var host string
	var dataBase int
	var passWord string

	flag.IntVar(&port, "p", 6379, "specify port to use.  defaults to 6379.")
	flag.StringVar(&host, "hostip", "127.0.0.1", "specify port to use.  defaults to 127.0.0.1")
	flag.IntVar(&dataBase, "db", 0, "specify port to use.  defaults to 0")
	flag.StringVar(&passWord, "password", "", "specify port to use.  defaults to \"\"")
	flag.Parse()

	addr := host + ":" + strconv.Itoa(port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passWord, // no password set
		DB:       dataBase, // use default DB
	})

	_, err = client.Ping().Result()
	if err != nil {
		panic(err)
	}
	var key string
	var val int64

	for index = 100; index >= 1; index-- {
		key = "key" + strconv.Itoa(index)
		val, err = client.Get(key).Int64()
		fmt.Printf("\n%s %d", key, val)
		if err != nil {
			panic(err)
		}
	}

}

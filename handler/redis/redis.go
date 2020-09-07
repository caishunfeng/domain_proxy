package redis

import (
	"domain_proxy/config"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/redis.v4"
)

type RedisCfg struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

var redisCfg RedisCfg

func Init() {
	jsonFile, err := os.Open(path.Join(config.ConfigRoot, "redis.json"))
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &redisCfg)
	if err != nil {
		panic(err)
	}
}

func CreateClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.Db,
	})

	_, err := client.Ping().Result()

	return client, err
}

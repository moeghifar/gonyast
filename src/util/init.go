package util

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // this blank import required by `jmoiron/sqlx`
	mgo "gopkg.in/mgo.v2"
)

var (
	// DBCon ...
	DBCon *sqlx.DB
	// MongoCon ,,,
	MongoCon *mgo.Session
	// RedisCon ,,,
	RedisCon *redis.Pool
)

// Init will initiate all required initial configuration such db, redis, mongo
func Init() (err error) {
	// if DBCon, err = initDatabase(); err != nil {
	// 	log.Fatal("[ERROR] failed initiate database :", err)
	// }
	// if mongoCon, err = initMongo(); err != nil {
	// 	log.Fatal("[ERROR] failed initiate mongo :", err)
	// }
	if RedisCon, err = initRedis(); err != nil {
		log.Fatal("[ERROR] failed initiate redis :", err)
	}
	return err
}

func initDatabase() (DBConn *sqlx.DB, err error) {
	connection := Config.DBConfig["core"]
	DBConn, err = sqlx.Connect("postgres", connection)
	return DBConn, err
}

func initMongo() (session *mgo.Session, err error) {
	mongoConnection := Config.MongoConfig["core"]
	session, err = mgo.Dial(mongoConnection)
	return session, err
}

func initRedis() (Pool *redis.Pool, err error) {
	redisConnection := Config.RedisConfig["core"]
	Pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConnection)
			return c, err
		},
	}
	return Pool, err
}

// SetRedis ...
func SetRedis(value string, key string) (err error) {
	openCon := RedisCon.Get()
	_, err = openCon.Do("SET", key, value)
	if err != nil && err != redis.ErrNil {
		log.Println("[ERROR] ->", err)
	}
	return
}

// GetRedis ...
func GetRedis(key string) (data string, err error) {
	openCon := RedisCon.Get()
	data, err = redis.String(openCon.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		log.Println("[ERROR] ->", err)
	}
	return
}

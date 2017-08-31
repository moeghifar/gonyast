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
	dbCon    *sqlx.DB
	mongoCon *mgo.Session
	redisCon *redis.Pool
)

// Init will initiate all required initial configuration such db, redis, mongo
func Init() (err error) {
	if dbCon, err = initDatabase(); err != nil {
		log.Fatal("[ERROR] failed initiate database :", err)
	}
	// if mongoCon, err = initMongo(); err != nil {
	// 	log.Fatal("[ERROR] failed initiate mongo :", err)
	// }
	if redisCon, err = initRedis(); err != nil {
		log.Fatal("[ERROR] failed initiate redis :", err)
	}
	return err
}

func initDatabase() (DBConn *sqlx.DB, err error) {
	connection := "user='postgres'dbname='db_alpha_app'password='root'sslmode='disable'"
	DBConn, err = sqlx.Connect("postgres", connection)
	return DBConn, err
}

func initMongo() (session *mgo.Session, err error) {
	mongoConnection := "localhost:27017"
	session, err = mgo.Dial(mongoConnection)
	return session, err
}

func initRedis() (Pool *redis.Pool, err error) {
	redisConnection := "localhost:6379"
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

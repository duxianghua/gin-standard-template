package internal

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

const (
	redisPingKey = "ping"
)

// Service for the core of our service
type Service struct {
	rdb  *redis.Client
	host string
	port int
}

// NewService create new service with conf
func NewService(conf *Conf) *Service {
	rdb := newRedisClient(conf)

	return &Service{
		rdb:  rdb,
		host: conf.Host,
		port: conf.Port,
	}
}

// Run the service forever
func (svc *Service) Run() {
	log.WithFields(log.Fields{"host": svc.host, "port": svc.port}).Info("run")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		svc.rdb.Incr(c, redisPingKey)
	})

	r.GET("/count", svc.counter)
	r.Run(fmt.Sprintf("%s:%d", svc.host, svc.port))
}

func (svc *Service) counter(c *gin.Context) {
	val, err := svc.rdb.Get(c, redisPingKey).Result()
	if err == redis.Nil {
		val = "0"
	} else if err != nil {
		log.WithFields(log.Fields{"cmd": "get", "key": redisPingKey}).Error(err)
		val = "0"
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		log.WithFields(log.Fields{"val": val, "key": redisPingKey}).Error("need to be int")
		i = 0
	}

	c.JSON(200, gin.H{
		"counter": i,
	})
}

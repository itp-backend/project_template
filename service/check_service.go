package service

import (
	"fmt"
	"github.com/rysmaadit/go-template/external/minio"
	"github.com/rysmaadit/go-template/external/mysql"
	"github.com/rysmaadit/go-template/external/redis"
	log "github.com/sirupsen/logrus"
)

type checkService struct {
	redisClient redis.Client
	mysqlClient mysql.Client
	minioClient minio.Client
}

type CheckService interface {
	CheckRedis() ([]byte, error)
	CheckMysql() ([]byte, error)
	CheckMinio() ([]byte, error)
}

func NewCheckService(redisClient redis.Client, mysqlClient mysql.Client, minioClient minio.Client) *checkService {
	return &checkService{
		redisClient: redisClient,
		mysqlClient: mysqlClient,
		minioClient: minioClient,
	}
}

func (c *checkService) CheckRedis() ([]byte, error) {
	err := c.redisClient.Ping()
	if err != nil {
		log.Warning(fmt.Errorf("redis ping failed: %v", err))
		return nil, err
	}
	return []byte("Redis OK"), err
}

func (c *checkService) CheckMysql() ([]byte, error) {
	err := c.mysqlClient.Ping()
	if err != nil {
		log.Warning(fmt.Errorf("mysql ping failed: %v", err))
		return nil, err
	}
	return []byte("Mysql OK"), err
}

func (c *checkService) CheckMinio() ([]byte, error) {
	err := c.minioClient.Ping()
	if err != nil {
		log.Warning(fmt.Errorf("minio ping failed: %v", err))
		return nil, err
	}
	return []byte("Minio OK"), err
}

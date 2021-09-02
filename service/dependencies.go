package service

import (
	"github.com/rysmaadit/go-template/app"
	"github.com/rysmaadit/go-template/external/jwt_client"
	"github.com/rysmaadit/go-template/external/minio"
	"github.com/rysmaadit/go-template/external/mysql"
	"github.com/rysmaadit/go-template/external/redis"
)

type Dependencies struct {
	AuthService  AuthServiceInterface
	CheckService CheckService
}

func InstantiateDependencies(application *app.Application) Dependencies {
	jwtWrapper := jwt_client.New()
	authService := NewAuthService(application.Config, jwtWrapper)
	redisClient := redis.NewRedisClient(application.Config.RedisAddress)
	mysqlClient := mysql.NewMysqlClient(mysql.ClientConfig{
		Username: application.Config.DBUsername,
		Password: application.Config.DBPassword,
		Host:     application.Config.DBHost,
		Port:     application.Config.DBPort,
		DBName:   application.Config.DBName,
	})
	minioClient := minio.NewMinioClient(minio.ClientConfig{
		Endpoint:   application.Config.MinioEndpoint,
		AccessKey:  application.Config.MinioAccessKey,
		SecretKey:  application.Config.MinioSecretKey,
		Region:     application.Config.MinioRegion,
		BucketName: application.Config.MinioBucket,
	})
	checkService := NewCheckService(redisClient, mysqlClient, minioClient)

	return Dependencies{
		AuthService:  authService,
		CheckService: checkService,
	}
}

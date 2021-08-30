package mysql

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client interface {
	Ping() error
}

type client struct {
	dbConnection *gorm.DB
}

func (c *client) Ping() error {
	var result int64
	tx := c.dbConnection.Raw("select 1").Scan(&result)
	if tx.Error != nil {
		return fmt.Errorf("mysql unable to serve basic query. %v", tx.Error)
	}
	return nil
}

func NewMysqlClient(config ClientConfig) *client {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
	dbConn, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("unable to initiate mysql connection. %v", err)
	}
	return &client{
		dbConnection: dbConn,
	}
}

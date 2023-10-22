package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"portto-homework/internal/utils/logger"
	"portto-homework/service/scanner-service/config"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

type DBClient struct {
	client *gorm.DB
}

func NewDBClient() *DBClient {
	once.Do(func() {
		cfg := config.GetDBConfig()
		fmt.Printf("cfg: %+v\n", cfg)
		db = connectDB(cfg)
		logger.SysLog().Info(context.Background(), fmt.Sprintf("Database [%s] connected", cfg.Address))
	})

	return &DBClient{client: db}
}

// Session creates an original gorm.DB session.
func (d *DBClient) Session() *gorm.DB {
	return d.client
}

func connectDB(cfg config.DBConfig) *gorm.DB {
	connect := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		cfg.Username,
		cfg.Password,
		cfg.Address,
		cfg.Database,
	)

	fmt.Printf("connect: %s\n", connect)
	var err error
	client, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if cfg.LogMode {
		client = client.Debug()
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	_client, err := client.DB()
	if err != nil {
		panic(err)
	}

	_client.SetMaxIdleConns(cfg.MaxIdle)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	_client.SetMaxOpenConns(cfg.MaxOpen)
	// SetConnMaxLifetime sets the maximum amount of timeUtil a connection may be reused.
	_client.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifeMin) * time.Minute)

	return client
}

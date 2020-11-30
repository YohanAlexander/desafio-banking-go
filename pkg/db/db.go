package db

import (
	"github.com/yohanalexander/desafio-banking-go/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB armazena a conexão com o banco de dados
type DB struct {
	Client *gorm.DB
}

// GetDB retorna a conexão com o banco de dados
func GetDB(connStr string) (*DB, error) {
	db, err := getDB(connStr)
	if err != nil {
		logger.Info.Fatal(err.Error())
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

// CloseDB fecha a conexão com o banco de dados
func (db *DB) CloseDB() error {
	sqlDB, err := db.Client.DB()
	if err != nil {
		logger.Info.Fatal(err.Error())
		return err
	}
	return sqlDB.Close()
}

// getDB estabelece a conexão com o banco de dados
func getDB(connStr string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Info.Fatal(err.Error())
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Info.Fatal(err.Error())
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		logger.Info.Fatal(err.Error())
		return nil, err
	}

	return db, nil
}

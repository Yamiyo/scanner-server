package db

import "portto-homework/internal/model/po"

func (d *DBClient) InitDBTable() error {
	return d.Session().AutoMigrate(
		&po.Block{},
		&po.Transaction{},
	)
}

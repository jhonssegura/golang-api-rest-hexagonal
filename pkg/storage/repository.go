package storage

import (
	"fmt"
	"golang-api-rest-hexagonal/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func NewDb(c *config.Config) ([]*gorm.DB, error) {

	arrayConnections := make([]*gorm.DB, 0)
	var db *gorm.DB
	var err error

	if len(c.DB.Mssql) > 0 {
		for _, v := range c.DB.Mssql {
			connStr := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
				v.Username,
				v.Password,
				v.Host,
				v.Port,
				v.Database)
			db, err = gorm.Open("mssql", connStr)
			if err != nil {
				fmt.Printf(" err = %v", err)
			}
			arrayConnections = append(arrayConnections, db)
			db = &gorm.DB{}
		}
	}
	return arrayConnections, nil
}

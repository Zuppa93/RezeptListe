package DBCom

import (
	_ "fmt"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// This "class" will take care of MySQL Queries

type DBCom struct {
	entityManager
}
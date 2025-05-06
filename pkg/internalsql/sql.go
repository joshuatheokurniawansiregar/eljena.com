package internalsql

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func Connect(dataSourceName string)(*sql.DB, error){
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil{
		log.Fatalf("error connecting to database %+v\n",err)
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}
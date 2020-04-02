package CRUD

import (
	"RezeptListe/Data/Entity"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"RezeptListe/Data"

)

type CRUDRezept struct {
}

// create
func (e *CRUDRezept) CreateRezept(rezept Entity.Rezept) bool {
	db, err := sql.Open(Data.DatabaseType, Data.Database)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(rezept.ToJSON())

	result, err := db.Query(`insert into rezepte(rezept) VALUES ('` + rezept.ToJSON() + `')`)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	defer result.Close()

	return false
}

// read
func (e *CRUDRezept) GetAlleRezepte() []Entity.Rezept {
	db,err := sql.Open(Data.DatabaseType,Data.Database)

	if err != nil{
		panic(err.Error())
	}

	defer db.Close()

	results,err := db.Query("select * from rezepte")

	var rezepte []Entity.Rezept

	for results.Next(){
		rezepte = append(rezepte,results.Scan(&Entity.Rezept{}))
	}




	return nil
}
func (e *CRUDRezept) GetRezeptByName(name string) Entity.Rezept {
	return Entity.Rezept{}
}

// update
func (e *CRUDRezept) UpdateRezept(updatedRezept Entity.Rezept) bool {
	return false
}

// delete
func (e *CRUDRezept) DeleteRezeptByName(name string) bool {
	return false
}

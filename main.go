package main

import (
	"RezeptListe/Data/CRUD"
	"RezeptListe/Data/Entity"
)

func main(){

	var CRUDrezept CRUD.CRUDRezept

	var rezept Entity.Rezept

	var zutat1,zutat2 Entity.Zutat

	zutat1.SetName("Unglaublicher Name")
	zutat1.SetMenge("24")

	zutat2.SetName("Schlechterer Name")
	zutat2.SetMenge("25")

	zutaten := []Entity.Zutat{zutat1,zutat2}

	rezept.SetName("Test")
	rezept.SetSchritte([]string{"schneiden","kleiden","schweigen"})
	rezept.SetZutaten(zutaten)

	CRUDrezept.CreateRezept(rezept)

}
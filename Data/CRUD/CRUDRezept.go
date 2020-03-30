package CRUD

import(
	"RezeptListe/Data/Entity"
)

type CRUDRezept struct {
	rezept Entity.Rezept
}

// create
func (e *CRUDRezept) createCRUDRezept(anleitung Entity.Anleitung,zutaten []Entity.Zutat)CRUDRezept{
	var crudRezept CRUDRezept
	crudRezept.rezept.SetAnleitung(anleitung)
	crudRezept.rezept.SetZutaten(zutaten)
	return crudRezept
}

// read
func (e *CRUDRezept) getRezept()Entity.Rezept{
	return e.rezept
}
// update
// delete
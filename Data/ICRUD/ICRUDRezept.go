package ICRUD

import "RezeptListe/Data/Entity"

type CRUDRezept interface{
	// create
	CreateRezept(rezept Entity.Rezept) bool
	// read
	GetAlleRezepte() []Entity.Rezept
	GetRezeptByName(name string) Entity.Rezept
	// update
	UpdateRezept(updatedRezept Entity.Rezept) bool
	// delete
	DeleteRezeptByName(name string) bool
}

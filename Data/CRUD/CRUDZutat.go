package CRUD

import(
	"RezeptListe/Data/Entity"
)

type CRUDZutat struct{
	zutat Entity.Zutat
}

// create

func (e *CRUDZutat) createCRUDZutat(name, menge string, tag []string)CRUDZutat{
	var crudZutat CRUDZutat

	crudZutat.zutat.SetName(name)
	crudZutat.zutat.SetMenge(menge)
	crudZutat.zutat.SetTag(tag)

	return crudZutat
}

// read
func (e *CRUDZutat) getZutat()Entity.Zutat{
	return e.zutat
}
// update
// delete
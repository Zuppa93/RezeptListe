package CRUD

import "RezeptListe/Data/Entity"

type CRUDRezept struct{

}
// create
func CreateRezept(rezept Entity.Rezept) bool{
	return false
}
// read
func GetAlleRezepte() []Entity.Rezept{
	return nil
}
func GetRezeptByName(name string) Entity.Rezept{
	return Entity.Rezept{}
}
// update
func UpdateRezept(updatedRezept Entity.Rezept) bool{
	return false
}
// delete
func DeleteRezeptByName(name string) bool{
	return false
}
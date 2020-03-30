package CRUD

import(
	"RezeptListe/Data/Entity"
)

type CRUDAnleitung struct{
	anleitung Entity.Anleitung
}

// Create
func createAnleitung (schritte []string) Entity.Anleitung{

	var anleitung Entity.Anleitung
	anleitung.SetSchritte(schritte)
	return anleitung

}
// Read
func (e *CRUDAnleitung)getAnleitung()Entity.Anleitung{
	return e.anleitung
}
// Update
func (e *CRUDAnleitung) updateSchritte(neueSchritte []string){
	e.anleitung.SetSchritte(neueSchritte)
}

// Wir nehmen an dass hier keine Fehler auftreten aufgrund des Filters
func (e *CRUDAnleitung) addSchrittAtIndex(schritt string, index int){
	// Man könnte das auch theoretisch in einem Statement zusammenfassen, aber das ist so schon kompliziert genug
	schritte := append(e.anleitung.GetSchritte()[:index],schritt)
	schritte = append(schritte,e.anleitung.GetSchritte()[index:]...)
	e.anleitung.SetSchritte(schritte)
}

// Delete


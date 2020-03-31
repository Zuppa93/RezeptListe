package Entity

/*
Just Getters and Setters.
Nothing extraordinary
I set zutat as a private structure. We won't need it except in Rezept
*/

type zutat struct {
	name string
	menge string
}

func (e *zutat) setName(name string){
	e.name = name
}

func (e *zutat) setMenge(menge string){
	e.menge = menge
}

func (e *zutat) getName()string{
	return e.name
}

func (e *zutat) getMenge()string{
	return e.menge
}

type Rezept struct {
	name string
	schritte []string
	zutaten []zutat
}

func (e *Rezept) setName(name string){
	e.name = name
}

func (e *Rezept) setSchritte(schritte []string){
	e.schritte = schritte
}

func (e *Rezept) setZutaten(zutaten []zutat){
	e.zutaten = zutaten
}

func (e *Rezept) getName()string{
	return e.name
}

func (e * Rezept) getSchritte()[]string{
	return e.schritte
}

func (e *Rezept) getZutaten()[]zutat{
	return e.zutaten
}

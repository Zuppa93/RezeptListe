package Entity

// Entity "Class"

type Rezept struct{
	anleitung Anleitung `json:anleitung`
	zutaten []Zutat     `json":zutaten`
}

func (e *Rezept) GetAnleitung() Anleitung {
	return e.anleitung
}

func (e *Rezept) GetZutaten() []Zutat {
	return e.zutaten
}

func (e *Rezept) SetAnleitung(newAnleitung Anleitung){
	e.anleitung = newAnleitung
}

func (e *Rezept) SetZutaten(newZutaten []Zutat){
	e.zutaten = newZutaten
}

func (e *Rezept) AddZutat(newZutat Zutat){
	e.zutaten = append(e.zutaten,newZutat)
}

func (e *Rezept) AddSchrittToAnleitung(schritt string){
	e.anleitung.AddSchritt(schritt)
}
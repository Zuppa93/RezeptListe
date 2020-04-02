package Entity

/*
Just Getters and Setters.
Nothing extraordinary
I set zutat as a private structure. We won't need it except in Rezept
*/

type Zutat struct {
	name string		`json:"name"`
	menge string	`json:"menge"`
}

func (e *Zutat) InitZutat(name,menge string){
	e.name = name
	e.menge = menge
}

func (e *Zutat) SetName(name string){
	e.name = name
}

func (e *Zutat) SetMenge(menge string){
	e.menge = menge
}

func (e *Zutat) GetName()string{
	return e.name
}

func (e *Zutat) GetMenge()string{
	return e.menge
}

type Rezept struct {
	name string			`json:"name"`
	schritte []string	`json:"schritte"`
	zutaten []Zutat		`json:"zutaten"`
}

func (e *Rezept) SetName(name string){
	e.name = name
}

func (e *Rezept) SetSchritte(schritte []string){
	e.schritte = schritte
}

func (e *Rezept) SetZutaten(zutaten []Zutat){
	e.zutaten = zutaten
}

func (e *Rezept) GetName()string{
	return e.name
}

func (e * Rezept) GetSchritte()[]string{
	return e.schritte
}

func (e *Rezept) GetZutaten()[]Zutat{
	return e.zutaten
}

func (e *Rezept) ToJSON()string{
	var json string = `{"name":"`
	json = json+e.name+`","schritte":[`

	for i := 0;i<len(e.schritte);i++{
		json = json+`{"schritt":"`+e.schritte[i]+`"}`
		if i+1 <len(e.schritte){
			// Gibt es einen nächsten Schritt fügen wir ein Komma hinzu
			json = json+`,`
		}
	}
	json = json+`],"zutaten":[`
	for i := 0;i<len(e.zutaten);i++{
		json = json+`{"name":"`+e.zutaten[i].name+`",`
		json = json+`"menge":"`+e.zutaten[i].menge+`"}`
		if i+1 <len(e.zutaten){
			// Gibt es einen nächsten Schritt fügen wir ein Komma hinzu
			json = json+`,`
		}
	}
	json = json+`]}`


	return json
}

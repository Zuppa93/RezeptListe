package Entity

// Entity "Class"

type Zutat struct {
	name string
	menge string
}

func (e *Zutat) setName(name string){
	e.name = name
}

func (e *Zutat) setMenge(menge string){
	e.menge = menge
}

func (e *Zutat) getName() string{
	return e.name
}

func (e *Zutat) getMenge() string {
	return e.menge
}
package Entity

// Entity "Class"

type Zutat struct {
	name  string `json:name`
	menge string `json:menge`
	tag []string `json:tag`
}

func (e *Zutat) SetName(name string){
	e.name = name
}

func (e *Zutat) SetMenge(menge string){
	e.menge = menge
}

func (e *Zutat) SetTag(tag []string ){
	e.tag = tag
}

func (e *Zutat) GetName() string{
	return e.name
}

func (e *Zutat) GetMenge() string {
	return e.menge
}

func (e *Zutat) GetTag() []string{
	return e.tag
}
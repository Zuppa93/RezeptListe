package Entity

// Entity "Class"

type Anleitung struct {
	schritte []string `json:"schritte"`
}

func (e *Anleitung) SetSchritte(schritte []string){
	e.schritte = schritte
}

func (e *Anleitung) GetSchritte() []string{
	return e.schritte
}
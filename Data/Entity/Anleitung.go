package Entity

// Entity "Class"

type Anleitung struct {
	schritte []string `json:"schritte"`
}

func (e *Anleitung) AddSchritt(newSchritt string){
	e.schritte = append(e.schritte,newSchritt)
}

func (e *Anleitung) DelLastSchritt(){
	e.schritte = append(e.schritte[:0],e.schritte[:len(e.schritte)-1]...)
}

func (e *Anleitung) DelSchrittByIndex(index int){
	e.schritte = append(e.schritte[:index],e.schritte[index+1:]...)
}

func (e *Anleitung) SetSchritte(schritte []string){
	e.schritte = schritte
}

func (e *Anleitung) GetSchritte() []string{
	return e.schritte
}
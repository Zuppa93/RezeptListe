package main

import (
	"RezeptListe/Data"
	"fmt"
)

func main(){

	index := 3

	testText := []string{"0","1","2","3","4","5","6","7","8","9",}

	anleitung := Data.Anleitung{}

	anleitung.SetSchritte(testText)

	fmt.Println(anleitung)
	fmt.Println("Nun löschen wir Index ",index)
	anleitung.DelSchrittByIndex(index)
	fmt.Println(anleitung)

	fmt.Println("Lösche letzten Schritt")
	anleitung.DelLastSchritt()
	fmt.Println(anleitung)


}
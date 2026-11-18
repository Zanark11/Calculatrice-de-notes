package main

import "fmt"

func main() {
	var nombreNotes int

	fmt.Println("Combien de notes souhaites-tu saisir ?")
	fmt.Scan(&nombreNotes)

	var somme float64 = 0

	for i := 0; i < nombreNotes; i++ {
	var note float64 = 0

	fmt.Println("Entre une note :")
	fmt.Scan(&note) 
	somme = somme + note 
}
	var moyenne float64
	
if nombreNotes == 0 {
    fmt.Println("Erreur : tu dois saisir au moins une note.")
} else {
	moyenne = somme / float64(nombreNotes)
	fmt.Println("La moyenne est :", moyenne)
}
}
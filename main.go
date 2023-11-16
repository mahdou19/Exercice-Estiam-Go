package main

import (
	"bufio"
	"estiam/dictionary"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	dict := dictionary.NewDictionary()

	for {
		fmt.Print("Entrez quelque chose : (add, define, list ou remove) : ")
		scanner.Scan()
		entreeUtilisateur := scanner.Text()

		if entreeUtilisateur == "add" {
			actionAdd(dict, scanner)

		} else if entreeUtilisateur == "define" {
			fmt.Println(entreeUtilisateur)
		} else if entreeUtilisateur == "list" {
			fmt.Println(entreeUtilisateur)
		} else if entreeUtilisateur == "remove" {
			fmt.Println(entreeUtilisateur)
		} else {
			fmt.Println("Entrez Non valide ! Choisisser dans cette liste: (add, define, list ou remove) : ")
		}
	}

}

func actionAdd(d *dictionary.Dictionary, scanner *bufio.Scanner) {
	fmt.Print("Entrez un mot : ")
	scanner.Scan()
	word := scanner.Text()
	fmt.Print("Entrez une definition du mot entrer précédemment : ")
	scanner.Scan()
	definition := scanner.Text()
	fmt.Println("La definition du mot ", word, "est : ", definition)

	d.Add(word, definition)
	fmt.Println("le mot < ", word, " > est ajouté.")
}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {

}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {

}

func actionList(d *dictionary.Dictionary) {

}

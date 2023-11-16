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
			actionList(dict)
		} else if entreeUtilisateur == "remove" {
			actionRemove(dict, scanner)
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

func actionRemove(d *dictionary.Dictionary, scanner *bufio.Scanner) {
	fmt.Print("Entrer le mot à supprimé du dictionnaire : ")
	scanner.Scan()
	word := scanner.Text()

	d.Remove(word)
	fmt.Printf("%s estsupprimé du dictionnaire.\n", word)
}

func actionList(d *dictionary.Dictionary) {
	entries := d.List()
	if len(entries) == 0 {
		fmt.Println("Le dictionaire est vide !")
		return
	}

	fmt.Println("Liste du dictionaire")
	for _, entry := range entries {
		fmt.Printf("- %s: %s \n", entry.Word, entry.Definition)
	}
}

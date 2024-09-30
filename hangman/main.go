package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type HangmanData struct {
	mot_brouillé  string
	mot_a_trouver string
	nombre_essais int
	coordonnées   []int
}

func hangmanreader(ent1 int, ent2 int) {
	readFile, err := os.Open("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lignes []string
	for fileScanner.Scan() {
		lignes = append(lignes, fileScanner.Text())
	}
	readFile.Close()
	for i := ent1; i <= ent2; i++ {
		fmt.Printf(lignes[i] + "\n")
	}
}

func RandomWord() string {
	data, err := os.ReadFile("words.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	var tab_words []string
	for i := 0; i < len(data); i++ {
		tab_words = append(tab_words, string(data[i]))
	}
	data, err = os.ReadFile("words2.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	for i := 0; i < len(data); i++ {
		tab_words = append(tab_words, string(data[i]))
	}
	data, err = os.ReadFile("words3.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	for i := 0; i < len(data); i++ {
		tab_words = append(tab_words, string(data[i]))
	}
	words := ""
	for i := 0; i < len(tab_words); i++ {
		words += string(tab_words[i][0])
	}
	var tab_Words []string
	Words := ""
	for i := 0; i < len(words); i++ {
		if words[i] != '\n' {
			Words += string(words[i])
		} else {
			tab_Words = append(tab_Words, Words)
			Words = ""
		}
	}
	randomisation := rand.Intn(len(tab_Words))
	return tab_Words[randomisation]
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func hideword(word string) string {
	var a []int
	var s []string
	for i := 0; i <= (len(word)/2)-1; i++ {
		b := rand.Intn(len(word) - 1)
		if i == 0 {
			a = append(a, b)
			continue
		} else {
			if contains(a, b) {
				i--
				continue
			} else {
				a = append(a, b)
			}
		}
	}
	for i := 0; i <= len(word)-1; i++ {
		s = append(s, "_")
		for j := 0; j <= len(a)-1; j++ {
			if i == a[j] {
				k := a[j]
				s[i] = string(word[k])
			}
		}
	}
	stringjoin := strings.Join(s, "")
	return string(stringjoin)
}

func main() {
	var userinput string
	var tempmotbrouillé string
	pendu := HangmanData{}
	var a []int
	pendu.mot_a_trouver = RandomWord()
	pendu.mot_brouillé = hideword(pendu.mot_a_trouver)
	pendu.nombre_essais = 0
	for i := 0; i <= 9; i++ {
		if i == 0 {
			a = append(a, 0)
			a = append(a, 8)
		} else {
			a = append(a, i*8)
			a = append(a, i*8+7)
		}
	}
	pendu.coordonnées = a
	tentative := 9
	for true {
		tempmotbrouillé = pendu.mot_brouillé
		word := ""
		for i := 0; i < len(tempmotbrouillé); i++ {
			if string(tempmotbrouillé[i]) == "_" {
				word += string(tempmotbrouillé[i]) + string(" ")
			} else {
				word += string(rune(tempmotbrouillé[i]-32)) + string(" ")
			}

		}
		fmt.Println("Veuillez entrer une lettre")
		fmt.Println(word)
		fmt.Scanln(&userinput)
		for i := 0; i < len(pendu.mot_a_trouver); i++ {
			if userinput == string(pendu.mot_a_trouver[i]) {
				if pendu.mot_a_trouver[i] == pendu.mot_brouillé[i] {
					continue
				} else {
					tempmotbrouilléRunes := []rune(pendu.mot_brouillé)
					tempmotbrouilléRunes[i] = rune(userinput[0])
					pendu.mot_brouillé = string(tempmotbrouilléRunes)
				}
			}
		}
		if pendu.mot_brouillé != tempmotbrouillé {
			fmt.Println("\nvotre lettre fait partie du mot")
		} else {
			fmt.Println("\nvotre lettre ne fait pas partie du mot")
			pendu.nombre_essais++
			tentative--
			fmt.Println("Il vous reste " + string(rune(tentative+48)) + " tentative(s)")
		}
		if pendu.nombre_essais == 0 {
			hangmanreader(pendu.coordonnées[0], pendu.coordonnées[1])
		} else {
			if pendu.nombre_essais*2 < len(pendu.coordonnées) && pendu.nombre_essais*2+1 < len(pendu.coordonnées) {
				hangmanreader(pendu.coordonnées[pendu.nombre_essais*2], pendu.coordonnées[pendu.nombre_essais*2+1])
			}
		}
		if pendu.nombre_essais == 9 {
			fmt.Println("Vous avez perdu")
			fmt.Println("Le mot était : " + string(pendu.mot_a_trouver))
			break
		}
		if pendu.mot_a_trouver == pendu.mot_brouillé {
			fmt.Println("Vous avez gagné")
			break
		}
	}
}

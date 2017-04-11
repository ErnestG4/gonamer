package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
	"log"
)

func randColor() string {
	var colors []string

	colorfile, err := os.Open("colors.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(colorfile)

	for scanner.Scan() {
		colors = append(colors, scanner.Text())
	}

	var colcount = rand.Intn(len(colors))
	var color = colors[colcount]
	return color
}

func randNoun() string {
	var nouns []string

	nounfile, err := os.Open("nouns.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(nounfile)

	for scanner.Scan() {
		nouns = append(nouns, scanner.Text())
	}

	var nouncount = rand.Intn(len(nouns))
	var noun = nouns[nouncount]
	return noun
}

func combineName(Color string, Noun string) string {
	var combinedNames = Color + Noun
	return combinedNames
}

func main() {

	rand.Seed(time.Now().UnixNano())

	var Color = randColor()
	var Noun = randNoun()

	var fullName = combineName(Color, Noun)

	fmt.Println(fullName)
//	fmt.Println("Combination #", colcount * nouncount, " of ", len(colors)*len(nouns))
}

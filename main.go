package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	characterSets  [5]string
	configInputs   [8]string //indexes are linked to keywordsConfig array!
	keywordsConfig [8]string //LINKED TO configInputs by index!
)

const (
	totalKeywords int = len(keywordsConfig)
)

func main() {
	keywordsConfig[0] = "passwordlength"
	keywordsConfig[1] = "passwordamount"
	keywordsConfig[2] = "uppercaseletters"
	keywordsConfig[3] = "uppercaseamount"
	keywordsConfig[4] = "numbers"
	keywordsConfig[5] = "numbersamount"
	keywordsConfig[6] = "specialcharacters"
	keywordsConfig[7] = "specialcharamount"

	characterSets[0] = "abcdefghijklmnopqrstuvwxyz"
	characterSets[1] = strings.ToUpper(characterSets[0])
	characterSets[2] = "0123456789"
	characterSets[3] = "!@#$%&*"
	characterSets[4] = characterSets[0] //setting the default

	checkConfig()
	InitializeGeneration()

	fmt.Scanln()
}

func checkConfig() {
	f, err := os.Open("config.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	lineByLine := bufio.NewScanner(f)
	a := 0
	for lineByLine.Scan() {
		if strings.Contains(lineByLine.Text(), "#") || lineByLine.Text() == "" {
			fmt.Println("Line Ignored")
		} else {
			if strings.Contains(lineByLine.Text(), (keywordsConfig[a] + " = ")) {
				var restant1 = strings.ReplaceAll(lineByLine.Text(), (keywordsConfig[a] + " = "), "")
				dummy, err := strconv.Atoi(restant1)
				if err != nil {
					fmt.Print("Value is not a number: ", restant1+"\n")
					putStringsFromConfigIntoArray(restant1, keywordsConfig[a])
				} else {
					fmt.Println("Value is a number!", dummy)
					putNumberFromConfigIntoArray(restant1, keywordsConfig[a])
				}
			}
			if a < totalKeywords {
				a++
			}
		}
	}
}

func putNumberFromConfigIntoArray(restant1, corrKeyword string) { //This is where the int values go
	for i := 0; i < totalKeywords; i++ {
		if corrKeyword == keywordsConfig[i] {
			configInputs[i] = restant1
		}
	}
}

func putStringsFromConfigIntoArray(restant1, corrKeyword string) { //This is where the string (to be converted to bools) go
	for i := 0; i < totalKeywords; i++ {
		if corrKeyword == keywordsConfig[i] {
			configInputs[i] = strings.ToLower(restant1)
		}
	}
}

func InitializeGeneration() {
	var (
		totalSet   []rune
		usableInts [8]int
	)
	rand.Seed(time.Now().Unix())

	if configInputs[2] == "true" {
		characterSets[4] += characterSets[1]
	}
	if configInputs[4] == "true" {
		characterSets[4] += characterSets[2]
	}
	if configInputs[6] == "true" {
		characterSets[4] += characterSets[3]
	}
	totalSet = []rune(characterSets[4])
	for i := 0; i < totalKeywords; i++ {
		refinedVal, err := strconv.Atoi(configInputs[i])
		if err != nil {
		} else {
			usableInts[i] = refinedVal
		}
	}
	fmt.Println()
	for x := 0; x < usableInts[1]; x++ {
		password := make([]rune, usableInts[0])
		for i := range password {
			password[i] = totalSet[rand.Intn(len(totalSet))]
		}
		fmt.Println(string(password))
	}

}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	allExercises := map[string]int{
		"displayfirstparam":    1,
		"displaylastparam":     1,
		"displayz":             1,
		"displaya":             1,
		"hello":                1,
		"onlya":                1,
		"onlyz":                1,
		"printdigits":          1,
		"strlen":               1,
		"paramcount":           1,
		"displayalpham":        1,
		"diplaylrevm":          1,
		"nrune":                2,
		"lastrune":             2,
		"printstr":             2,
		"strrev":               2,
		"firstrune":            2,
		"printreversealphabet": 2,
		"printalphabet":        2,
		"wdmatch":              2,
		"ispowerof2":           2,
		"rot13":                2,
		"lastword":             2,
		"rot14":                2,
		"max":                  2,
		"reduceint":            2,
		"switchcase":           3,
		"swap":                 3,
		"compare":              3,
		"expandstr":            3,
		"tabmult":              3,
		"searchreplace":        3,
		"alphamirror":          3,
		"doop":                 3,
		"findrevprime":         3,
		"reversebits":          3,
		"chunks":               3,
		"foldint":              3,
		"swapbits":             4,
		"capitalize":           4,
		"repeatalpha":          4,
		"atoi":                 4,
		"reversestrcap":        4,
		"printbits":            4,
		"inter":                4,
		"union":                4,
		"piglatin":             4,
		"romannumbers":         4,
		"firstword":            5,
		"sortwordarr":          5,
		"cleanstr":             5,
		"addprimesum":          5,
		"printhex":             5,
		"gcd":                  5,
		"printcomb":            6,
		"split":                6,
		"hiddenp":              6,
		"rostring":             6,
		"revwstr":              6,
		"reverserange":         6,
		"slice":                6,
		"itoa":                 7,
		"atoibase":             7,
		"foreach":              7,
		"fprime":               7,
		"printrevcomb":         7,
		"listsize":             8,
		"rpncalc":              8,
		"brackets":             8,
		"options":              8,
		"grouping":             8,
		"printmemory":          8,
		"listremoveif":         9,
		"itoabase":             9,
		"brainfuck":            9,
	}
	fmt.Println(selectRandomExercises(allExercises))
}

func selectRandomExercises(exercises map[string]int) []string {

	levels := make([]int, 0)
	for _, level := range exercises {
		if !contains(levels, level) {
			levels = append(levels, level)
		}
	}

	selectedLevels := make([]int, 0)
	for len(selectedLevels) < 3 {
		randomLevel := levels[rand.Intn(len(levels))]
		if !contains(selectedLevels, randomLevel) {
			selectedLevels = append(selectedLevels, randomLevel)
		}
	}

	selectedExercises := make([]string, 0)
	for exercise, exerciseLevel := range exercises {
		for _, level := range selectedLevels {
			if exerciseLevel == level && len(selectedExercises) < 3{
				selectedExercises = append(selectedExercises, exercise)
				break
			}
		}
	}

	return selectedExercises
}

// Fonction utilitaire pour vérifier si une valeur existe dans une liste
func contains(list []int, value int) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Exercise struct {
	ID       int                      `json:"id"`
	Name     string                   `json:"name"`
	Type     string                   `json:"type"`
	Attrs    Attrs                    `json:"attrs"`
	Children map[string]ChildExercise `json:"children"`
}

type Attrs struct {
	Allowed []string `json:"allowedFunctions"`
}

type ChildExercise struct {
	Name  string `json:"name"`
	Attrs Attrs  `json:"attrs"`
}

func main() {
	// Lire le fichier JSON
	data, err := os.ReadFile("exam4.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	// Struct pour stocker les données du fichier JSON
	var exercise Exercise

	// Décoder le fichier JSON dans la structure Exercise
	err = json.Unmarshal(data, &exercise)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}

	// Extraire les noms d'exercices
	exerciseNames := make([]string, 0)
	for _, childExercise := range exercise.Children {
		exerciseNames = append(exerciseNames, childExercise.Name)
	}

	str := "{"

	for _, exerciseName := range exerciseNames {
		allowed := exercise.Children[exerciseName].Attrs.Allowed
		str += fmt.Sprintf(`"%v":[%v],`, exerciseName, allowed)
	}

	str += "}"

	str = strings.ReplaceAll(str, "github.com/01-edu/z01.PrintRune", `"github.com/01-edu/z01.PrintRune",`)
	str = strings.ReplaceAll(str, "--allow-builtin", `"--allow-builtin",`)
	str = strings.ReplaceAll(str, "fmt.*", `"fmt.*",`)
	str = strings.ReplaceAll(str, "fmt.Printf", `"fmt.Printf",`)
	str = strings.ReplaceAll(str, "fmt.Println", `"fmt.Println",`)
	str = strings.ReplaceAll(str, "os.Args", `"os.Args",`)
	str = strings.ReplaceAll(str, "os.*", `"os.*",`)
	str = strings.ReplaceAll(str, "strconv.Atoi", `"strconv.Atoi",`)
	str = strings.ReplaceAll(str, "strconv.Itoa", `"strconv.Itoa",`)
	str = strings.ReplaceAll(str, "unicode.IsGraphic", `"unicode.IsGraphic",`)
	str = strings.ReplaceAll(str, "strconv.*", `"strconv.*",`)
	str = strings.ReplaceAll(str, "strings.Split", `"strings.Split",`)

	os.WriteFile("r4.json", []byte(str), 0777)

	fmt.Println(exercise.Children)
}

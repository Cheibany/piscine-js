package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	requirements := []string{`"fmt"`, `"os"`}
	Split("../printalphabet/main.go", requirements)
}

func Split(filename string, requirements []string) {
	_content, err := os.ReadFile(filename)
	numbersOfpackageAllowed := len(requirements) - 1
	if err != nil {
		return
	}
	content := strings.Split(string(_content), "\n")
	packageLine := content[0]
	packageContent := strings.Split(packageLine, " ")

	if len(packageContent) != 2 {
		fmt.Println("package line must content two elements: \n-the keyword package and the name of the package")
	} else {
		if !(packageContent[0] == "package" && packageContent[1] == requirements[0]) {
			fmt.Println("package line must content two elements: \n\t-the keyword package and \n\tthe name of the package")
		}
	}

	if content[1] != "" {
		fmt.Println("After the package you should have an empty line")
	}

	if strings.Contains(content[2], "import") {
		if strings.Contains(content[2], "(") {
			if strings.Contains(content[numbersOfpackageAllowed+3], ")") {
				_importedPackage := content[3:numbersOfpackageAllowed+3]
				importedPackage := []string{}
				for _, imported := range _importedPackage {
					importedPackage = append(importedPackage, strings.TrimSpace(imported))
				}
				fmt.Println(importedPackage)
				for _, imported := range importedPackage {
					if !Contains(imported, importedPackage) {
						fmt.Println("Cheating...")
					}
				}
				fmt.Println("Package ok")
			}
		} else {
			if strings.Contains(content[2], "\"") {
				if imported := strings.Split(content[2], " "); imported[1] != requirements[1] {
					fmt.Println(requirements[1], imported[1])
					fmt.Println("Cheating")
				} else {
					fmt.Println("Package ok")
				}
			}
		}
	}

}

func Contains(str string, tab []string) bool {
	for _, v := range tab {
		if v == str {
			return true
		}
	}
	return false
}
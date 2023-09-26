package main

type Data struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Attrs    interface{} `json:"attrs"`
	Children Child       `json:"children"`
}

type Child struct {
	Signup      interface{} `json:"signup"`
	Onboarding  interface{} `json:"onboarding"`
	PiscineGo   Piscine     `json:"piscine-go"`
	Div01       interface{} `json:"interface"`
	Rattrapages interface{} `json:"rattrapages"`
}

type Piscine struct {
	Atrr     interface{} `json:"attrs"`
	Children interface{} `json:"children"`
}

// func main() {

// 	filename := "all01.json"

// 	content, err := os.ReadFile(filename)
// 	if err != nil {
// 		fmt.Println("Erreur :", err)
// 		return
// 	}

// 	var data Data

// 	err = json.Unmarshal(content, &data)
// 	if err != nil {
// 		fmt.Println("Erreur :", err)
// 		return
// 	}

// 	fmt.Println(data.Children.PiscineGo.Atrr)

// }

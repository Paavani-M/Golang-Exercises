package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`paavani`] = []string{`Charlie`, `jacky`, `sunny`}

	delete(m, `bond_james`)

	for k, v := range m {
		fmt.Println(k)
		for i, v1 := range v {
			fmt.Println(i, v1)
		}
	}
}

// Using the code from the previous example, delete a record from your map. Now print the map
// out using the “range” loop

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var p []person

	err := json.Unmarshal([]byte(s), &p)

	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println()

	for _, v := range p {
		fmt.Println("First Name, Lastname, Age:")
		fmt.Println(v.First, v.Last, v.Age)
		fmt.Println("Sayings:")
		for _, v1 := range v.Sayings {
			fmt.Println(v1)
		}
		fmt.Println()
	}

}

// Starting with this code, unmarshal the JSON into a Go data structure.

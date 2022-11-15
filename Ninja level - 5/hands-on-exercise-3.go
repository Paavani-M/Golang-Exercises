package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "brown",
		},
		fourwheel: true,
	}

	t2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "silver",
		},
		luxury: true,
	}

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println()

	fmt.Println("truck")
	fmt.Printf("Doors: %d\n", t1.doors)
	fmt.Printf("Color: %s\n", t1.color)
	fmt.Println(t1.fourwheel)
	fmt.Println()

	fmt.Println("sedan")
	fmt.Printf("Doors: %d\n", t2.doors)
	fmt.Printf("color: %s\n", t2.color)
	fmt.Println(t2.luxury)

}

// ● Create a new type: vehicle.
// ○ The underlying type is a struct.
// ○ The fields:
// ■ doors
// ■ color
// ● Create two new types: truck & sedan.
// ○ The underlying type of each of these new types is a struct.
// ○ Embed the “vehicle” type in both truck & sedan.
// ○ Give truck the field “fourWheel” which will be set to bool.
// ○ Give sedan the field “luxury” which will be set to bool. solution
// ● Using the vehicle, truck, and sedan structs:
// ○ using a composite literal, create a value of type truck and assign values to the
// fields;
// ○ using a composite literal, create a value of type sedan and assign values to the
// fields.
// ● Print out each of these values.
// ● Print out a single field from each of these values.

package types

import "fmt"

type vertex2 struct {
	Lat, Long float64
}

var m map[string]vertex2

/*
HelloMaps - Go maps

A map maps keys to values.
The zero value of a map is nil. A nil map has no keys, nor can keys be added.
The make function returns a map of the given type, initialized and ready for use.
*/
func HelloMaps() {
	fmt.Printf("\n****Running types.HelloMaps(), go maps \n")

	m = make(map[string]vertex2)
	m["Bell Labs"] = vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

/*
MapLiterals - Map literals are like struct literals, but the keys are required.
*/
func MapLiterals() {
	fmt.Printf("\n****Running types.MapLiterals(), map literals \n")
	var m = map[string]vertex2{
		"Bell Labs": vertex2{
			40.68433, -74.39967,
		},
		"Google": vertex2{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)
}

// MapLiteralsWithNoTypeName - If the top-level type is just a type name, you can omit it from the elements of the literal.
func MapLiteralsWithNoTypeName() {
	fmt.Printf("\n****Running types.MapLiteralsWithNoTypeName(), map literals with no type name \n")
	var m = map[string]vertex2{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

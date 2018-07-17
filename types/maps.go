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

/*
MapMutation - Mutating Maps.

Insert or update an element in map m:
	m[key] = elem
Retrieve an element:
	elem = m[key]
Delete an element:
	delete(m, key)
Test that a key is present with a two-value assignment:
	elem, ok = m[key]
If key is in m, ok is true. If not, ok is false.
If key is not in the map, then elem is the zero value for the map's element type.

Note: if elem or ok have not yet been declared you could use a short declaration form:
elem, ok := m[key]
*/
func MapMutation() {

	fmt.Printf("\n****Running types.MapMutation(), Insert or update an element in map  \n")
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

package interfaces

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

/*
HelloStringers - One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}

A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
*/
func HelloStringers() {
	fmt.Printf("****Running interfaces.HelloStringers(), Go stringers interface example. \n")
	a := person{"Arthur Dent", 42}
	z := person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

type ipAddr [4]byte

func (p ipAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", p[0], p[1], p[2], p[3])
}

//StringersExcercise - Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
//For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
func StringersExcercise() {
	fmt.Printf("****Running interfaces.StringersExcercise(), Running go stringers excercise. \n")
	hosts := map[string]ipAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

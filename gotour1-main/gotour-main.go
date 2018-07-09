package main

import (
	"fmt"
	"go-learn/gotour1"
)

func main() {
	fmt.Printf("============== START - Running gotour chapter 1 codes =============\n")
	gotour1.TestPackages()
	gotour1.TestImport()
	gotour1.Add(2, 5)
	gotour1.AddOmitTypeDeclaration(6464666646, 8794984945)
	gotour1.Swap("6", "87")
	gotour1.Split(17)
	gotour1.VaribalesDeclaration()
	gotour1.VaribalesWithInitializers()
	gotour1.ShortVaribaleDeclaration()
	gotour1.BasicTypes()
	gotour1.ZeroValueVariable()
	gotour1.ConvertTypes()
	gotour1.TypeInterference()
	fmt.Printf("\n\n=============== END =============\n")
}

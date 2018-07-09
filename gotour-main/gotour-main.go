package main

import (
	"fmt"
	"go-learn/basic"
)

func main() {
	fmt.Printf("============== START - Running gotour chapter 1 codes =============\n")
	basic.TestPackages()
	basic.TestImport()
	basic.Add(2, 5)
	basic.AddOmitTypeDeclaration(6464666646, 8794984945)
	basic.Swap("6", "87")
	basic.Split(17)
	basic.VaribalesDeclaration()
	basic.VaribalesWithInitializers()
	basic.ShortVaribaleDeclaration()
	basic.DataTypes()
	basic.ZeroValueVariable()
	basic.ConvertTypes()
	basic.TypeInterference()
	fmt.Printf("\n\n=============== END =============\n")
}

package main

import (
	"fmt"
	"go-learn/basic"
	"go-learn/flowcontrol"
)

func main() {
	fmt.Printf("============== START - Running gotour Basic codes =============\n")
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
	basic.Constants()
	basic.NumericConstant()
	fmt.Printf("\n=============== END - Basic =============\n")

	fmt.Printf("\n============== START - Running gotour FlowControl codes =============\n")
	flowcontrol.SimpleForLoop()
	fmt.Printf("\n\n=============== END - FlowControl =============\n")
}

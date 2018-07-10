package main

import (
	"fmt"
	"go-learn/basic"
	"go-learn/flowcontrol"
	"math"
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
	flowcontrol.ForOptionalInitAndPost()
	//flowcontrol.ForeverLoop()
	fmt.Println(flowcontrol.Sqrt(-4), flowcontrol.Sqrt(2))
	fmt.Println(flowcontrol.IfWithShortStatement(3, 2, 10))
	fmt.Println(flowcontrol.IfWithElse(3, 2, 20))
	//flowcontrol.SqrtMatching(1)
	flowcontrol.SqrtMatching(25000)
	fmt.Println("actual sqrt", math.Sqrt(25000))
	//flowcontrol.SqrtMatching(3)

	fmt.Printf("\n\n=============== END - FlowControl =============\n")
}

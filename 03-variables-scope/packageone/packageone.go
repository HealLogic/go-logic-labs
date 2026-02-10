package packageone

import "fmt"

//declare a package-level variable named PackageVar of type string
var PackageVar string = "this is package variable"

func PrintMe(myVar, blockVar string) {
	fmt.Println(myVar, blockVar, PackageVar)
}

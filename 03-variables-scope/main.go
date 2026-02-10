package main

import "myapp/packageone"

// Declare a package-level variable named myVar of type string
var myVar string = "this is package Main variable"

func main() {
	// Declare a block-level variable named blockVar of type string
	var blockVar string = "this is block variable"

	// Call the PrintMe function from packageone, passing myVar, blockVar, and PackageVar
	packageone.PrintMe(myVar, blockVar, packageone.PackageVar)
}

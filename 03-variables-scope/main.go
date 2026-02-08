package main

import (
    "fmt"
    "myapp/packageone" 
)

// Package-level variable:
// Available to every function in this file (and package).
var one = "One"

func main() {
    // Block-level variable:
    // Only available inside this main function.
    // Other functions (like myFunc) cannot see this.
    var somethingElse = "This is a block level variable"

    fmt.Println(somethingElse)

    myFunc()

    // Accessing an exported variable from another package
    // Notice the Capital Letter 'P' in PublicVar
    newString := packageone.PublicVar
    fmt.Println("From packageone:", newString)

    // Accessing an exported function from another package
    packageone.Exported()
}

func myFunc() {
    // This function can see 'one' because it is package-level.
    // But it CANNOT see 'somethingElse'.
    fmt.Println(one)
}
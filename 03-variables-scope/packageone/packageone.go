package packageone

import "fmt"

// Variables/Functions starting with a Capital Letter are Exported (Public).
// They can be seen by other packages (like main).
var PublicVar = "I am a public variable"

// Variables starting with a lowercase letter are Private.
var privateVar = "I am hidden"

func Exported() {
    fmt.Println("I am an exported function from packageone")
}
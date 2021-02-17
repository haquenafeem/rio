# Rio
Rio, is a beautiful console logging package named after my pet bird.
You can use this package to colorfully log texts, or any data types in the console.
Using color to print things, will make your developing experience easier and clearer.

# Example
### Code (cmd/main.go) : 

```golang

package main

import (
	"fmt"

	"github.com/haquenafeem/rio"
)

func main() {
	fmt.Println("Rio is beautiful.....Check her swag")
	fmt.Printf("%v %v %v %v %v %v %v %v",
		rio.Ignore("Ignore"),
		rio.Error("Error"),
		rio.Success("Success"),
		rio.Warn("Warn"),
		rio.Log("Log"),
		rio.Info("Info"),
		rio.Question("Question"),
		rio.Quote(`"Quote"`))
	fmt.Println()
	fmt.Println()
	fmt.Println("I'm sad because they", rio.Ignore("ignored"), "me.")
	fmt.Println("This could be an", rio.Error("error."))
	fmt.Println("We", rio.Success("successfully"), "did it.")
	fmt.Println("This", rio.Warn("warning"), "should not be avoided.")
	fmt.Println("Let me", rio.Log("log"), "it to the console.")
	fmt.Println("Here is an information:", rio.Info("Rio is my pet parrot."))
	fmt.Println("Hey!", rio.Question("How you doin?."))
	fmt.Println(rio.Quote(`"Love the life you live, live the life you love"`), "- Bob Marley.")

	fmt.Println()
}


```
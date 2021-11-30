//Some info about the package
package main

import (
	"fmt"

	"github.com/josepht96/learning/Go/printer"
)

func main() {
	fmt.Printf("Hello go\n")
}

//TestPublic doesnt do much
func TestPublic() {
	printer.PrinterRun()
	fmt.Printf("Hello\n")
}

//Library should be consumed by another package
//Name of library must match directory name

//Main package - application entry point
//Needs to contain main() func
//Can be in any directory
//App setup - not so much business logic

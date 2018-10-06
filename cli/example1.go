/*package main

import (
    "fmt"
    "os"
)

func main() {
    for i, a := range os.Args[1:] {
        fmt.Printf("Argument %d is %s\n", i+1, a)
    }

}
*/

package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8000, "specify port to use.  defaults to 8000.")
	flag.Parse()

	fmt.Printf("port = %d\n", port)
	fmt.Printf("other args: %+v\n", flag.Args())
}

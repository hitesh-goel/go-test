package injection

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Greet(os.Stdout, "Chris")
}

//Greet greet with name
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!", name)
}

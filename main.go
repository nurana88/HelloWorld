package main

import (
	"fmt"

	"github.com/jonbodner/language/mapper"
)

func main() {
	fmt.Println(mapper.Greet("How you doing?"))
	fmt.Println(mapper.Greet("Wie geht's dir?"))
	fmt.Println(mapper.Greet("Kak ti"))

}

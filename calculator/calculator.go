package calculator

import (
	"flag"
	"fmt"
	"github.com/mnogu/go-calculator"
	"log"
)

func Run() {
	expression := flag.String("e", "", "Expression to calculate")

	flag.Parse()

	if *expression == "" {
		log.Fatal("Expression is required")
	}

	val, err := calculator.Calculate(*expression)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
}

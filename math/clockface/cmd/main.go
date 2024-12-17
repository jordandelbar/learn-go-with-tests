package main

import (
	"learn_go_with_tests/math/clockface"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}

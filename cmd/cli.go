package main

import (
	"fmt"
	"log"

	"github.com/8qfx1ai5/viewcrypt/internal/encode"
)

func main() {
	in := `<!-- this is a comment -->
	<p>Lorem ipsum dolor sit amet, consectetur <b>adipisicing</b> elit. Repellat, deleniti!</p>`
	xCfg := encode.XsltConfig{}
	out, err := encode.HTML(in, xCfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(out)
}

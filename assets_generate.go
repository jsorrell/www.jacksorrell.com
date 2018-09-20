// +build ignore

package main

import (
	"log"

	"github.com/shurcooL/vfsgen"

	"github.com/jsorrell/www.jacksorrell.com/data"
)

func main() {
	err := vfsgen.Generate(data.Assets, vfsgen.Options{
		PackageName:  "data",
		BuildTags:    "!dev",
		VariableName: "Assets",
		Filename:     "data/assets_vfsdata.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}

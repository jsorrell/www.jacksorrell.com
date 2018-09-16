// +build ignore

package main

import (
	"log"

	"github.com/jsorrell/www.jacksorrell.com/data"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(data.Assets, vfsgen.Options{
		PackageName:  "data",
		BuildTags:    "!dev",
		VariableName: "Assets",
		Filename: "data/assets_vfsdata.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}

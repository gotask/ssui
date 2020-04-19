// vfsgen.go

// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

//uilib
func VfsGen(dir string) error {
	var fs http.FileSystem = http.Dir(dir)

	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "ssui",
		BuildTags:    "",
		VariableName: "UILib",
	})
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func main() {
	VfsGen("uilib")
}

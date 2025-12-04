package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run scripts/newday.go <dayXX>  (e.g. day03)")
		os.Exit(1)
	}
	day := os.Args[1]        // "day03"
	pkg := day               // package name
	srcDir := "days/template"
	dstDir := filepath.Join("days", day)

	// make sure template exists
	if _, err := os.Stat(srcDir); err != nil {
		panic(fmt.Errorf("template folder missing at %s", srcDir))
	}

	must(os.MkdirAll(dstDir, 0o755))

	must(filepath.WalkDir(srcDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(srcDir, path)
		dstPath := filepath.Join(dstDir, rel)

		if d.IsDir() {
			return os.MkdirAll(dstPath, 0o755)
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if strings.HasSuffix(d.Name(), ".go") {
			// remove ignore tags
			b = bytes.ReplaceAll(b, []byte("//go:build ignore\n"), []byte(""))
			b = bytes.ReplaceAll(b, []byte("// +build ignore\n\n"), []byte(""))

			// package rename
			b = bytes.ReplaceAll(b, []byte("package template"), []byte("package "+pkg))

			// import rewrite for template package
			b = bytes.ReplaceAll(b, []byte(`/days/template"`), []byte(`/days/`+day+`"`))
			b = bytes.ReplaceAll(b, []byte(`/days/_template"`), []byte(`/days/`+day+`"`))

			// input path rewrite
			b = bytes.ReplaceAll(b, []byte("days/template/input.txt"), []byte("days/"+day+"/input.txt"))
			b = bytes.ReplaceAll(b, []byte("days/_template/input.txt"), []byte("days/"+day+"/input.txt"))

			// input path rewrite
			b = bytes.ReplaceAll(b, []byte("template.Part"), []byte(day+".Part"))
		}

		return os.WriteFile(dstPath, b, 0o644)
	}))

	fmt.Println("created", dstDir)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

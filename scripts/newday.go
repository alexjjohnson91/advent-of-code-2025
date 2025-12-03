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

	entries, err := os.ReadDir(srcDir)
	must(err)

	for _, e := range entries {
		srcPath := filepath.Join(srcDir, e.Name())
		dstPath := filepath.Join(dstDir, e.Name())

		if e.IsDir() {
			continue // ignore nested dirs for simplicity
		}

		b, err := os.ReadFile(srcPath)
		must(err)

		// replace "package template" -> "package dayNN"
		if strings.HasSuffix(e.Name(), ".go") {
			b = bytes.ReplaceAll(b, []byte("package template"), []byte("package "+pkg))
		}

		must(os.WriteFile(dstPath, b, 0o644))
	}

	fmt.Println("created", dstDir)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zyy17/go-editor"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("editor <file>\n")
		os.Exit(2)
	}

	filePath := os.Args[1]
	defaultEditor := editor.NewDefaultEditor([]string{})

	if err := defaultEditor.Launch(filePath); err != nil {
		log.Fatalf("edit error: %v\n", err)
	}
}

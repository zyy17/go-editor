## go-editor

The source code of go-editor comes from [Kubernetes](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/kubectl/pkg/cmd/util/editor/editor.go) and refractor as the clean Go module.

You can embed go-editor in your command-line tool like kubectl (kubectl edit <your_resource>).

For example (examples/editor/main.go):

```go
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
```

The `go-editor` will invoke the local editor to edit your file and save back.
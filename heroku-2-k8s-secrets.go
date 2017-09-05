package main

import (
	"bufio"
	b64 "encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/bgentry/heroku-go"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	userName := flag.String("u", "foo", "Heroku Username")
	pass := flag.String("p", "foo", "Heroku Password/API key")
	appName := flag.String("a", "foo", "Heroku App Name")
	secretName := flag.String("s", "my-k8s-secret", "Kubernetes Secret and File Name")

	flag.Parse()

	client := heroku.Client{Username: *userName, Password: *pass}

	configVars, err := client.ConfigVarInfo(*appName)
	check(err)

	fileName := *secretName + ".yaml"

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("File does not exist or cannot be created")
		os.Exit(1)
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	templateTop := `apiVersion: v1
kind: Secret
metadata:
`

	templateBot := `type: Opaque
data:
`

	fmt.Fprint(w, templateTop)
	fmt.Fprintf(w, "  name: %v\n", *secretName)
	fmt.Fprint(w, templateBot)

	for k, v := range configVars {
		fmt.Fprintf(w, "  %v: %v\n", strings.ToLower(k), b64.StdEncoding.EncodeToString([]byte(v)))
	}
	w.Flush()
}

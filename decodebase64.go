package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// mirip seperti while read di bash
func dec(name string) {
	var decodedBy, _ = base64.StdEncoding.DecodeString(name)
	var decodedStr = string(decodedBy)
	djson := decodedStr
	res, _ := pretty(djson)
	fmt.Println(res)

}

func decode() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		data := sc.Text()
		clean := strings.Split(data, ".")
		colorize(Red, "\t\tHEADER:")
		colorize(Yellow, "")
		dec(clean[0])
		fmt.Println("\n")
		colorize(Red, "\t\tPAYLOAD:")
		colorize(Blue, "")
		dec(clean[1])
		colorize(Reset, "")
	}
}

func pretty(str string) (string, error) {
	var prettyJson bytes.Buffer
	if err := json.Indent(&prettyJson, []byte(str), "", "   "); err != nil {
		return "", err
	}
	return prettyJson.String(), nil
}

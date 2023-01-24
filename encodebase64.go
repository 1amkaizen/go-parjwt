package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

// mirip seperti while read di bash
func encode() {
	sca := bufio.NewScanner(os.Stdin)
	for sca.Scan() {
		enc := sca.Text()
		var dataenc = (enc)
		var encode = base64.StdEncoding.EncodeToString([]byte(dataenc))
		fmt.Println(encode)

	}
}

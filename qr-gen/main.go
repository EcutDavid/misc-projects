package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func encode(url string) string {
	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}
	sEnc := base64.StdEncoding.EncodeToString([]byte(png))
	return fmt.Sprintf("data:image/png;base64,%s", string(sEnc))
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need tell me what to encode")
	}
	b64 := encode(os.Args[1])
	fmt.Print(b64)
}

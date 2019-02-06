package main

import (
	"encoding/base64"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

var lightBlue = color.RGBA{0, 199, 206, 255}

func encode(url string, newStyle bool) string {
	var png []byte
	var err error
	if newStyle {
		// github.com/skip2/go-qrcode has to write the QR code to a file if non-default colors are needed.
		if err = qrcode.WriteColorFile(url, qrcode.Medium, 256, color.White, lightBlue, "qr.png"); err != nil {
			log.Fatal(err)
		}
		png, err = ioutil.ReadFile("qr.png")
	} else {
		png, err = qrcode.Encode(url, qrcode.Medium, 256)
	}

	if err != nil {
		log.Fatal(err)
	}

	sEnc := base64.StdEncoding.EncodeToString([]byte(png))
	return fmt.Sprintf("data:image/png;base64,%s", string(sEnc))
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please tell me what to encode")
	}
	b64 := encode(os.Args[1], len(os.Args) > 2)
	fmt.Print(b64)
}

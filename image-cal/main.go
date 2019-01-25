package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)

func encode(url string, newStyle bool) string {
	var png []byte
	var err error
	if newStyle {
		err = qrcode.WriteColorFile(url, qrcode.Medium, 256, color.White, color.RGBA{0, 199, 206, 255}, "qr.png")
		if err != nil {
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

var numImage []image.Image
var recSize image.Point

func init() {
	numImage = make([]image.Image, 10)
	for i := 0; i < 10; i++ {
		key := strconv.Itoa(i)
		raw, _ := ioutil.ReadFile(fmt.Sprintf("./num/%s.png", key))

		image, _, err := image.Decode(strings.NewReader(string(raw)))
		if err != nil {
			log.Fatal(err)
		}
		numImage[i] = image
		// buffer := bytes.NewBuffer([]byte{})
		// png.Encode(buffer, image)
		// ioutil.WriteFile(fmt.Sprintf("%s.png", key), buffer.Bytes(), 0666)

		// fmt.Println(format)
		// png.Encode(os.Stdout, image)
	}
	recSize = numImage[0].Bounds().Size()
}

func drawImg(num int) {
	if num >= 9999 || num < 0 {
		num = 42
	}
	str := strconv.Itoa(num)
	nums := []int{}
	for i := 0; i < len(str); i++ {
		dig, _ := strconv.Atoi(str[i : i+1])
		nums = append(nums, dig)
	}

	m := image.NewRGBA(image.Rect(0, 0, 360, 200))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	for i := 0; i < len(nums); i++ {
		draw.Draw(m,
			image.Rectangle{image.Point{20 + recSize.X*i, 20}, image.Point{20 + recSize.X*(i+1), 20 + recSize.Y}},
			numImage[nums[i]],
			image.ZP,
			draw.Over,
		)
	}

	buffer := bytes.NewBuffer([]byte{})
	png.Encode(buffer, m)
	sEnc := base64.StdEncoding.EncodeToString(buffer.Bytes())
	fmt.Printf("data:image/png;base64,%s", sEnc)
}

func parseNum(str string) int {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return 42
	}
	parts := strings.Split(str, "*")
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.Trim(parts[i], " ")
	}

	result, err := strconv.Atoi(parts[0])
	if err != nil {
		return 42
	}
	for i := 1; i < len(parts); i++ {
		num, err := strconv.Atoi(parts[i])
		if err != nil {
			return 42
		}
		result *= num
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need tell me what to calc")
	}
	num := parseNum(os.Args[1])
	drawImg(num)
}

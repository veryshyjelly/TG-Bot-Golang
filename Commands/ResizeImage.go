package Commands

import (
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"github.com/sunshineplan/imgconv"
	"image"
	"image/draw"
	"image/png"
	"io"
	"net/http"
	"net/url"
)

func ResizeImage(data io.Reader) (*bytes.Buffer, error) {

	myImage, err := imgconv.Decode(data)
	if err != nil {
		return nil, err
	}

	//var marks image.Image
	if myImage.Bounds().Dx() >= myImage.Bounds().Dy() {
		myImage = imgconv.Resize(myImage, imgconv.ResizeOption{Width: 512})
	} else {
		myImage = imgconv.Resize(myImage, imgconv.ResizeOption{Height: 512})
	}

	res := new(bytes.Buffer)
	err = imgconv.Write(res, myImage, imgconv.FormatOption{Format: imgconv.PNG, EncodeOption: []imgconv.EncodeOption{imgconv.PNGCompressionLevel(png.BestSpeed)}})
	if err != nil {
		return nil, err
	}

	myImage, err = png.Decode(res)
	if err != nil {
		return nil, err
	}

	if myImage.Bounds().Dx() > 512 || myImage.Bounds().Dy() > 512 {
		dst := image.NewRGBA(image.Rect(0, 0, 512, 512))
		draw.Draw(dst, image.Rect(0, 0, 512, 512), myImage, myImage.Bounds().Min, draw.Src)
		err = png.Encode(res, dst)
	} else {
		err = png.Encode(res, myImage)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func SendResizeImage(baseUrl, apiToken string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	imagePath, err := Functions.GetFile(baseUrl, message.Document.FileId)
	imageLink := "https://api.telegram.org/file/bot" + apiToken + "/" + url.QueryEscape(imagePath.FileId)

	res, err := http.Get(imageLink)
	if err != nil {
		return nil, err
	}

	data := new(bytes.Buffer)
	_, err = io.Copy(data, res.Body)
	if err != nil {
		return nil, err
	}

	resImage, err := ResizeImage(data)
	if err != nil {
		return nil, err
	}

	return Functions.SendPhotoByReader(baseUrl, resImage, message, "here", false)
}

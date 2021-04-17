package main

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

import (
	"github.com/bep/gowebp/libwebp"
	"github.com/bep/gowebp/libwebp/webpoptions"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/bmp"
)

func webpEncoder(p1, p2 string, quality int, Log bool, c chan int) (err error) {
	// if convert fails, return error; success nil

	log.Debugf("target: %s with quality of %d", path.Base(p1), quality)
	//var buf bytes.Buffer
	var img image.Image

	data, err := ioutil.ReadFile(p1)
	if err != nil {
		chanErr(c)
		return
	}

	contentType := getFileContentType(data[:512])
	if strings.Contains(contentType, "jpeg") {
		img, _ = jpeg.Decode(bytes.NewReader(data))
	} else if strings.Contains(contentType, "png") {
		img, _ = png.Decode(bytes.NewReader(data))
	} else if strings.Contains(contentType, "bmp") {
		img, _ = bmp.Decode(bytes.NewReader(data))
	} else if strings.Contains(contentType, "gif") {
		// TODO: need to support animated webp
		log.Warn("Gif support is not perfect!")
		img, _ = gif.Decode(bytes.NewReader(data))
	}

	if img == nil {
		msg := "image file " + path.Base(p1) + " is corrupted or not supported"
		log.Debug(msg)
		err = errors.New(msg)
		chanErr(c)
		return
	}

	//
	output, err := os.Create(p2)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	options := webpoptions.EncodingOptions{
		Quality:        quality,
		EncodingPreset: webpoptions.EncodingPresetDefault,
	}
	if err = libwebp.Encode(output, img, options); err != nil {
		log.Error(err)
		chanErr(c)
		return
	}

	if Log {
		log.Info("Save to " + p2 + " ok!\n")
	}

	chanErr(c)

	return nil
}

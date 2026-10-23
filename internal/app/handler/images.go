package handler

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

var errBadImage = errors.New("image must be a base64 encoded jpeg or png")

func saveImageAsWebp(dir, dataURL string) (string, error) {
	data := dataURL
	if i := strings.Index(dataURL, ","); i >= 0 {
		data = dataURL[i+1:]
	}

	raw, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", errBadImage
	}

	img, _, err := image.Decode(bytes.NewReader(raw))
	if err != nil {
		return "", errBadImage
	}

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		return "", err
	}

	name := uuid.New().String() + ".webp"
	out, err := os.Create(filepath.Join(dir, name))
	if err != nil {
		return "", err
	}
	defer out.Close()

	if err := webp.Encode(out, img, options); err != nil {
		os.Remove(out.Name())
		return "", err
	}
	return name, nil
}

func removeImage(dir, url string) error {
	err := os.Remove(filepath.Join(dir, filepath.Base(url)))
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}

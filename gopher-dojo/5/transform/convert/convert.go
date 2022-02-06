package convert

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

type convert interface {
	Encode(io.Writer)
	Decode(io.Reader)
}

type Config struct {
	SrcDIR  string
	DestDIR string
	BaseExt string
	ConvExt string
}
type ConvertCmd struct {
	Setting *Config
}

func (c *ConvertCmd) Decode(r io.Reader) (image.Image, string, error) {
	return image.Decode(r)
}

func (c *ConvertCmd) Encode(w io.Writer, img image.Image) error {
	switch c.Setting.ConvExt {
	case "png":
		return png.Encode(w, img)
	case "jpeg":
		return jpeg.Encode(w, img, &jpeg.Options{Quality: 100})
	}
	return nil
}

func (c *ConvertCmd) Execute() error {
	var (
		srcdir  = flag.String("srcdir", "./", "string flag")
		destdir = flag.String("destdir", "./dest", "string flag")
		baseExt = flag.String("baseExt", "jpeg", "string flag")
		convExt = flag.String("convExt", "png", "string flag")
	)
	flag.Parse()

	c.Setting = &Config{
		SrcDIR:  *srcdir,
		DestDIR: *destdir,
		BaseExt: *baseExt,
		ConvExt: *convExt,
	}

	err := filepath.Walk(c.Setting.SrcDIR, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		ext := filepath.Ext(info.Name())
		if ext != "" {
			n := info.Name()
			baseName := n[:len(n)-len(filepath.Ext(n))]
			exFile, err := os.Open(c.Setting.SrcDIR + info.Name())
			if err != nil {
				fmt.Println(err)
			}
			img, _, Err := c.Decode(exFile)
			if Err != nil {
				return errors.New("Failed decode.")
			}
			f, err := os.Create(fmt.Sprintf("%s.%s", baseName, c.Setting.ConvExt))
			if err != nil {
				return errors.New("Failed Open file.")
			}
			err = c.Encode(f, img)
			if err != nil {
				return errors.New("Failed encode.")
			}
		}
		return nil
	})

	return err
}

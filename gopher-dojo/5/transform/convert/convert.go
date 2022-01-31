package convert

import (
	"errors"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

type convert interface {
	Encode(io.Writer)
	Decode(io.Reader)
}

type ConvertSetting struct {
	srcdir  string
	destdir string
	baseExt string
	convExt string
}
type ConvertCmd struct {
	Setting ConvertSetting
}

func (c ConvertCmd) Decode(r io.Reader) (image.Image, string, error) {
	return image.Decode(r)
}

// func init() {
// 	flag.BoolVar(&n, "n", false, "Output line")
// }

func Execute() {
	var (
		srcdir  = flag.String("srcdir", "./", "string flag")
		destdir = flag.String("destdir", "./dest", "string flag")
		baseExt = flag.String("baseExt", "jpeg", "string flag")
		convExt = flag.String("convExt", "png", "string flag")
	)
	flag.Parse()

	cs := ConvertSetting{
		srcdir:  *srcdir,
		destdir: *destdir,
		baseExt: *baseExt,
		convExt: *convExt,
	}

	cc := ConvertCmd{
		Setting: cs,
	}
	// if err := rootCmd.Execute(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	err := filepath.Walk(cc.Setting.srcdir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 特定のディレクトリを無視したい場合は `filepath.SkipDir` を返す
		// 例えば `AAA` という名前のディレクトリを無視する場合は以下のようにする
		// if info.IsDir() && info.Name() == "AAA" {
		// 	return filepath.SkipDir
		// }

		// fmt.Printf("path: %#v\n", path)
		fmt.Printf("ext: %#v\n", filepath.Ext(info.Name()))
		ext := filepath.Ext(info.Name())
		if ext != "" {
			n := info.Name()
			baseName := n[:len(n)-len(filepath.Ext(n))]
			exFile, err := os.Open(cc.Setting.srcdir + info.Name())
			if err != nil {
				fmt.Println(err)
			}
			// img, _, Err := image.Decode(exFile)
			img, _, Err := cc.Decode(exFile)
			if Err != nil {
				fmt.Println(Err)
				return errors.New("decode失敗")
			}
			f, err := os.Create(baseName + ".png")
			if err != nil {
				return errors.New("オープン失敗")
			}
			err = png.Encode(f, img)
			if err != nil {
				return errors.New("encode失敗")
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

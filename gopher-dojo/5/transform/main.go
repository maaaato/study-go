package main

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk("image", func(path string, info os.FileInfo, err error) error {
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
			fmt.Println("image/" + info.Name())
			exFile, err := os.Open("./image/" + info.Name())
			img, _, Err := image.Decode(exFile)
			if Err != nil {
				return errors.New("Decode失敗")
			}
			f, err := os.Create("./")
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

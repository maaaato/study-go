package convert

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func Do() {
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
			n := info.Name()
			baseName := n[:len(n)-len(filepath.Ext(n))]
			exFile, err := os.Open("image/" + info.Name())
			if err != nil {
				fmt.Println(err)
			}
			img, _, Err := image.Decode(exFile)
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

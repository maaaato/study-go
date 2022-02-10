package cmd

import (
	"flag"
	"os"

	convert "github.com/maaaato/study-go/gopher-dojo/5/transform/convert"
)

var (
	srcdir  = flag.String("srcdir", "./", "string flag")
	destdir = flag.String("destdir", "./dest", "string flag")
	baseExt = flag.String("baseExt", "jpeg", "string flag")
	convExt = flag.String("convExt", "png", "string flag")
)

func Execute() {
	var conv convert.Convert
	flag.Parse()
	conv.Setting = &convert.Config{
		SrcDIR:  *srcdir,
		DestDIR: *destdir,
		BaseExt: *baseExt,
		ConvExt: *convExt,
	}

	err := conv.Execute()
	if err != nil {
		os.Exit(1)
	}
}

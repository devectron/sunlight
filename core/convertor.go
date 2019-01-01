package core

import (
	"strings"

	"github.com/ConvertAPI/convertapi-go"
	"github.com/ConvertAPI/convertapi-go/config"
	"github.com/ConvertAPI/convertapi-go/param"
	"github.com/devectron/sunlight/log"
)

func Convertor(srcfile string, apisecret string, format string) (string, error) {
	f := strings.Split(format, "to")
	log.Inf("Converting %s to %s ...", srcfile, f[1])
	config.Default.Secret = apisecret
	fileParam := param.NewPath("file", srcfile, nil)
	pdfRes := convertapi.ConvDef(f[0], f[1], fileParam)
	files, err := pdfRes.ToPath("/tmp")
	if err != nil {
		return "", err[0]
	}
	// need to return file[0].ParamReader.reader
	log.Inf("File saved to: %s", files[0].Name())
	return files[0].Name(), nil
}

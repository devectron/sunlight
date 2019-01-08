package core

import (
	"io"
	"strings"

	"github.com/devectron/sunlight/log"
	"github.com/hihebark/convertapi-go"
	"github.com/hihebark/convertapi-go/config"
	"github.com/hihebark/convertapi-go/param"
)

func Convertor(srcfile string, apisecret string, format string) (*convertapi.Result, error) {
	f := strings.Split(format, "to")
	log.Inf("Converting %s to %s ...", srcfile, f[1])
	config.Default.Secret = apisecret
	fileParam := param.NewPath("file", srcfile, nil)
	pdfRes := convertapi.ConvDef(f[0], f[1], fileParam)
	return pdfRes, nil
}
func ConvertorR(srcfile io.Reader, name, apisecret, format string) (*convertapi.Result, error) {
	f := strings.Split(format, "to")
	log.Inf("Converting to %s ...", f[1])
	config.Default.Secret = apisecret
	fileParam := param.NewFileReader(name, srcfile, nil)
	pdfRes := convertapi.ConvDef(f[0], f[1], fileParam)
	return pdfRes, nil
}

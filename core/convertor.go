package core

import (
	"io"
	"strings"

	"github.com/ConvertAPI/convertapi-go"
	"github.com/ConvertAPI/convertapi-go/config"
	"github.com/ConvertAPI/convertapi-go/param"
	"github.com/devectron/sunlight/log"
)

func Convertor(srcfile string, apisecret string, format string) (*convertapi.Result, error) {
	f := strings.Split(format, "to")
	log.Inf("Converting %s to %s ...", srcfile, f[1])
	config.Default.Secret = apisecret
	fileParam := param.NewPath("file", srcfile, nil)
	pdfRes := convertapi.ConvDef(f[0], f[1], fileParam)
	return pdfRes, nil
}

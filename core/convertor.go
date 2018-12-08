package core

import (
	"fmt"
	"strings"

	"github.com/ConvertAPI/convertapi-go"
	"github.com/ConvertAPI/convertapi-go/config"
	"github.com/ConvertAPI/convertapi-go/param"
	"github.com/devectron/sunlight/log"
)

func Convertor(srcfile string, dstfile string, apisecret string, format string) (string, bool) {
	f := strings.Split(format, "to")
	log.Inf("Converting %s to %s ...", srcfile, f[1])
	config.Default.Secret = apisecret
	fileParam := param.NewPath("file", srcfile, nil)
	pdfRes := convertapi.ConvDef(f[0], f[1], fileParam)
	if files, err := pdfRes.ToPath("tmp"); err == nil {
		log.Inf("File saved to: ", files[0].Name())
		return "", true
	} else {
		log.Err("Error while converting %v", err)
		return fmt.Sprintf("%v", err), false
	}
}

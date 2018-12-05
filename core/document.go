package core

import (
	"path"
	"strings"

	"github.com/devectron/sunlight/log"
	//	unicommon "github.com/unidoc/unidoc/common"
	"github.com/unidoc/unidoc/pdf/creator"
)

//dep ensure -add github.com/unidoc/unidoc github.com/unidoc/unidoc

//ImagesToPdf convert image to a pdf file.
func ImagesToPdf(imgsrc string, pdfdst string) {
	name := strings.Split(imgsrc, path.Ext(imgsrc))
	log.Inf("Converting %s to %s.pdf", name[0], name[0])
	img, err := creator.NewImageFromFile(imgsrc)
	if err != nil {
		log.Err("Error while creating new image from file %v", err)
	}
	img.ScaleToWidth(612.0)
	height := 612.0 * img.Height() / img.Width()
	c := creator.New()
	c.SetPageSize(creator.PageSize{612, height})
	c.NewPage()
	img.SetPos(0, 0)
	_ = c.Draw(img)
	c.WriteToFile(pdfdst + "/" + name[0] + ".pdf")
	//if err != nil {
	//	log.Err("Error while writing to file %v", err)
	//}
}

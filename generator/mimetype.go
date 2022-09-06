package generator

import (
	"github.com/Kangrao0o/goepub/resource"
	"github.com/Kangrao0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

type MimeType struct{}

func (m *MimeType) Write(savePath string) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Errorf("mimetype write err: %v when os.Getwd", err)
		return err
	}
	if err := utils.CreateDir(savePath); err != nil {
		log.Errorf("mimetype write err: %v when create tmp dir", err)
		return err
	}

	dir = strings.TrimSuffix(dir, "generator")
	tplFilename := filepath.Join(dir, resource.MimeTypeEpub3Path)
	filename := filepath.Join(savePath, "mimetype")
	if _, err := utils.CopyFile(tplFilename, filename); err != nil {
		log.Errorf("mimetype write err: %v when copy file", err)
		return err
	}
	return nil
}

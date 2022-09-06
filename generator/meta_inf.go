package generator

import (
	"github.com/Kangrao0o/goepub/resource"
	"github.com/Kangrao0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

// MetaInf 元数据
type MetaInf struct{}

func (m *MetaInf) Write(savePath string) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Errorf("meta info write err: %v when os.Getwd", err)
		return err
	}
	if err := utils.CreateDir(savePath); err != nil {
		log.Errorf("meta info write err: %v when create tmp dir", err)
		return err
	}

	dir = strings.TrimSuffix(dir, "generator")
	metaInfSrcPath := filepath.Join(dir, resource.MetaInfEpub3Path)
	if err := utils.CopyDir(metaInfSrcPath, savePath); err != nil {
		log.Errorf("meta info write err: %v when copy dir", err)
		return err
	}
	return nil
}

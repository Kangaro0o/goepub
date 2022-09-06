package generator

import (
	"github.com/Kangaro0o/goepub/resource"
	"github.com/Kangaro0o/goepub/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

// Style CSS 样式
type Style struct{}

func (s *Style) Write(savePath string) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Errorf("style write err: %v when os.Getwd", err)
		return err
	}
	if err := utils.CreateDir(savePath); err != nil {
		log.Errorf("style write err: %v when create tmp dir", err)
		return err
	}

	dir = strings.TrimSuffix(dir, "generator")
	stylesSrcPath := filepath.Join(dir, resource.StyleEpub3Path)
	if err := utils.CopyDir(stylesSrcPath, savePath); err != nil {
		log.Errorf("style write err: %v when copy dir: %s", err, stylesSrcPath)
		return err
	}
	return nil
}

package utils

import (
	log "github.com/sirupsen/logrus"
	"io/fs"
	"os"
)

func CreateDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		// 路径不存在，则创建
		if err = os.MkdirAll(path, fs.ModePerm); err != nil {
			log.Errorf("create path %s err: %v", path, err)
			return err
		}
	}
	return nil
}

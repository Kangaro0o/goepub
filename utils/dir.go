package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// CreateDir 创建目录
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

// FileIsExisted 文件是否存在
func FileIsExisted(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// CopyFile 复制文件
func CopyFile(src, dest string) (int64, error) {
	// 获取源文件的权限
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	fi, err := srcFile.Stat()
	if err != nil {
		return 0, err
	}
	perm := fi.Mode()
	defer srcFile.Close()

	bytes, err := ioutil.ReadFile(src)
	if err != nil {
		return 0, err
	}

	if err := ioutil.WriteFile(dest, bytes, perm); err != nil {
		return 0, err
	}
	return int64(len(bytes)), nil
}

// CopyDir 复制文件夹
func CopyDir(srcPath, destPath string) error {
	// 检查源目录是否正确
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else if !srcInfo.IsDir() {
		return fmt.Errorf("srcPath isn't correct dir")
	}

	// 检查目的目录是否正确
	if destInfo, err := os.Stat(destPath); err != nil {
		return err
	} else if !destInfo.IsDir() {
		return fmt.Errorf("destPath isn't correct dir")
	}

	if strings.TrimSpace(srcPath) == strings.TrimSpace(destPath) {
		return fmt.Errorf("srcPath and destPath are the same")
	}

	err := filepath.Walk(srcPath, func(path string, f fs.FileInfo, err error) error {
		if f == nil {
			return err
		}

		// 生成新路径
		destNewPath := strings.Replace(path, srcPath, destPath, -1)
		if !f.IsDir() {
			CopyFile(path, destNewPath)
		} else if !FileIsExisted(destNewPath) {
			return CreateDir(destNewPath)
		}
		return nil
	})
	return err
}

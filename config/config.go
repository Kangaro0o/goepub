package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

type Epub3 struct {
	UID      string    `yaml:"uid"`
	BookName string    `yaml:"book-name"`
	Author   string    `yaml:"author"`
	Date     string    `yaml:"date"`
	Rights   string    `yaml:"rights"`
	Language string    `yaml:"language"`
	Chapters []Chapter `yaml:"chapters"`
}

type Chapter struct {
	ID        string `yaml:"id"`
	Href      string `yaml:"href"`
	PlayOrder int    `yaml:"play_order"`
	Name      string `yaml:"name"`
}

type Config struct {
	Epub3Conf Epub3 `yaml:"epub3"`
}

var (
	conf *Config
	once sync.Once
)

func loadConfig(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(bytes, &conf); err != nil {
		log.Fatal(err)
	}
}

// Get 获取 Epub3 文档的配置信息
func Get() *Config {
	once.Do(func() {
		loadConfig("./application.yml")
	})
	return conf
}

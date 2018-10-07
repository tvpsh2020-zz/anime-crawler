package config

import (
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type iConfig interface {
	loadServerConfig()
	loadSMTPConfig()
	loadAnimeConfig()
}

type config struct {
	configFile []byte
}

type serverConfig struct {
	Setting serverConfigDetail `yaml:"server"`
}
type serverConfigDetail struct {
	Port                          string        `yaml:"port"`
	CheckServerStatusTimeInMinute time.Duration `yaml:"checkServerStatusTimeInMinute"`
	FetchTimeInMinute             time.Duration `yaml:"fetchTimeInMinute"`
	DmhyWebsiteURL                string        `yaml:"dmhyWebsiteUrl"`
}

// Server config can be use by every module
var Server serverConfig

type smtpConfig struct {
	Setting smtpConfigDetail `yaml:"smtpMail"`
}

type smtpConfigDetail struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Server   string `yaml:"server"`
	Port     string `yaml:"port"`
	From     string `yaml:"from"`
	To       string `yaml:"to"`
}

// SMTP config can be use by every module
var SMTP smtpConfig

type animeConfig struct {
	Setting []animeConfigDetail `yaml:"anime"`
}
type animeConfigDetail struct {
	QueryString string   `yaml:"queryString"`
	Keywords    []string `yaml:"keywords"`
}

// Anime config can be use by every module
var Anime animeConfig

func newConfig(yamlFile []byte) iConfig {
	return &config{configFile: yamlFile}
}

func (c *config) loadServerConfig() {
	if err := yaml.Unmarshal(c.configFile, &Server); err != nil {
		panic("cannot read config file of server part: " + err.Error())
	}
}
func (c *config) loadSMTPConfig() {
	if err := yaml.Unmarshal(c.configFile, &SMTP); err != nil {
		panic("cannot read config file of smtp part: " + err.Error())
	}
}
func (c *config) loadAnimeConfig() {
	if err := yaml.Unmarshal(c.configFile, &Anime); err != nil {
		panic("cannot read config file of anime part: " + err.Error())
	}
}

// Initialize all config
func Initialize() {
	yamlFile, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		panic("cannot open config file: " + err.Error())
	}

	config := newConfig(yamlFile)
	config.loadServerConfig()
	config.loadSMTPConfig()
	config.loadAnimeConfig()
}

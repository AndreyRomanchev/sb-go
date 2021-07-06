package conf

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	Port      uint16 `yaml:"Port"`
	DbUrl     string `yaml:"DbUrl"`
	JaegerUrl string `yaml:"JaegerUrl"`
	AppID     string `yaml:"AppID"`
}

func NewConfig() (*Config, error) {

	configFile := flag.String("config", "config.yaml", "configuration file")
	flag.Parse()

	c := Config{}
	var err error
	var yamlFile []byte
	yamlFile, err = ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("yamlFile err #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	var portUint16 uint16
	p, ok := os.LookupEnv("PORT")
	if ok {
		port, err := strconv.ParseUint(p, 10, 16)
		if err != nil {
			return nil, fmt.Errorf("cannot convert port to uint16")
		}
		portUint16 = uint16(port)
	}
	if portUint16 != 0 {
		c.Port = portUint16
	}

	var dbUrl *url.URL
	u, ok := os.LookupEnv("DB_URL")
	if ok {
		dbUrl, err = url.Parse(u)
		if err != nil {
			return nil, fmt.Errorf("%s is invalid DB_URL", dbUrl)
		}
	}
	if dbUrl != nil {
		c.DbUrl = dbUrl.String()
	}

	var jaegerUrl *url.URL
	j, ok := os.LookupEnv("JAEGER_URL")
	if ok {
		jaegerUrl, err = url.Parse(j)
		if err != nil {
			return nil, fmt.Errorf("%s is invalid JAEGER_URL", jaegerUrl)
		}
	}
	if jaegerUrl != nil {
		c.JaegerUrl = jaegerUrl.String()
	}

	var appID string
	appID, ok = os.LookupEnv("APP_ID")
	if appID != "" {
		c.AppID = appID
	}

	return &c, nil
}

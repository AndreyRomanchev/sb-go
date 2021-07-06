package conf

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	Port      uint16
	DbUrl     *url.URL
	JaegerUrl *url.URL
	AppID     string
}

func NewConfig() (*Config, error) {

	var port_uint16 uint16
	p, ok := os.LookupEnv("PORT")
	if !ok {
		port_uint16 = 8080
	} else {
		port, err := strconv.ParseUint(p, 10, 16)
		if err != nil {
			return nil, fmt.Errorf("cannot convert port to uint16")
		}
		port_uint16 = uint16(port)
	}

	var dbUrl *url.URL
	var err error
	u, ok := os.LookupEnv("DB_URL")
	if !ok {
		dbUrl, _ = url.Parse("postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	} else {
		dbUrl, err = url.Parse(u)
		if err != nil {
			return nil, fmt.Errorf("%s is invalid DB_URL", dbUrl)
		}
	}

	var jaegerUrl *url.URL
	j, ok := os.LookupEnv("JAEGER_URL")
	if !ok {
		jaegerUrl, _ = url.Parse("http://jaeger:16686")
	} else {
		jaegerUrl, err = url.Parse(j)
		if err != nil {
			return nil, fmt.Errorf("%s is invalid JAEGER_URL", jaegerUrl)
		}
	}

	var appID string
	appID, ok = os.LookupEnv("APP_ID")
	if !ok {
		appID = "test"
	}

	return &Config{Port: port_uint16, DbUrl: dbUrl, JaegerUrl: jaegerUrl, AppID: appID}, nil
}

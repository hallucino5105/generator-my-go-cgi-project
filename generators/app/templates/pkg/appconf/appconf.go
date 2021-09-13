package appconf

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hallucino5105/glog"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Title string         `yaml:"title"`
	Serve AppConfigServe `yaml:"serve"`
}

type AppConfigServe struct {
	PublicPath string `yaml:"public_path"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
}

func LoadAppConfig() (*AppConfig, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	confpath := fmt.Sprintf("%s/config_init.yaml", cwd)
	glog.Debugf("load init config from \"%s\"", confpath)

	buf, err := ioutil.ReadFile(confpath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	conf := AppConfig{}
	if err := yaml.Unmarshal(buf, &conf); err != nil {
		return nil, errors.WithStack(err)
	}

	return &conf, nil
}

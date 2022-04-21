package service

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Config 站点配置
type Config struct {
	Debug               bool
	EnablePasswordLogin bool
	Oauth2              struct {
		Github struct {
			Enable       bool
			ClientID     string
			ClientSecret string
			RedirectURL  string
			Scopes       []string
		}
	}

	v *viper.Viper
}

// Read 从配置文件读取配置
func (c *Config) Read(path string) error {
	c.v = viper.New()
	c.v.SetConfigFile(path)
	if err := c.v.ReadInConfig(); err != nil {
		return err
	}
	if err := c.v.Unmarshal(c); err != nil {
		return err
	}
	return nil
}

// Save 将配置保存到配置文件
func (c *Config) Save() error {
	if data, err := yaml.Marshal(c); err != nil {
		return err
	} else {
		return ioutil.WriteFile(c.v.ConfigFileUsed(), data, os.ModePerm)
	}

}

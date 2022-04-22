package service

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// AgentConfig is the configuration for the agent
type AgentConfig struct {
	Token               string
	ZoneAccessWhitelist map[string]string // key: zone name, value: notes
	ZoneIDList          map[string]string // key: zone name, value: zone id

	v *viper.Viper
}

// Read reads the configuration from the given file
func (ac *AgentConfig) Read(path string) error {
	ac.v = viper.New()
	ac.v.SetConfigFile(path)
	if err := ac.v.ReadInConfig(); err != nil {
		return err
	}
	if err := ac.v.Unmarshal(ac); err != nil {
		return err
	}
	if ac.ZoneIDList == nil {
		ac.ZoneIDList = make(map[string]string)
	}
	return nil
}

// Save writes the configuration
func (ac *AgentConfig) Save() error {
	if data, err := yaml.Marshal(ac); err != nil {
		return err
	} else {
		return ioutil.WriteFile(ac.v.ConfigFileUsed(), data, os.ModePerm)
	}
}

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

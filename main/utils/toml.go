package utils

import "github.com/BurntSushi/toml"

type config struct {
	Mysql map[string]MysqlServer `toml:"mysql"`
	Host  string                 `toml:"host"`
	Port  string                 `toml:"port "`
}

// 解析toml文件
func UnmarshalConfig(tomlFile string) (*config, error) {
	config := &config{}
	if _, err := toml.DecodeFile(tomlFile, config); err != nil {
		return nil, err
	}
	return config, nil
}

// 获取mysql的配置
func (c *config) MysqlConf(key string) (MysqlServer, bool) {
	s, ok := c.Mysql[key]
	return s, ok
}

// 服务器监听地址
func (c *config) ListenAddr() string {
	return c.Host + ":" + c.Port
}

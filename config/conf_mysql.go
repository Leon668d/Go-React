package config

import "strconv"

type Mysql struct {
	Host         string `yaml:"host"`                                 // 服务器地址：端口
	Port         int    `yaml:"port"`                                 // 端口
	User         string `yaml:"user"`                                 // 数据库用户名
	Password     string `yaml:"password"`                             // 数据库密码
	Log_level    string `yaml:"log_level"`                            // 日志等级， debug就是输出全部sql,dev,release
	DB           string `yaml:"db"`                                   // 数据库名
	Config       string `yaml:"config"`                               // 高级配置，例如 charset
	MaxIdleConns int    `json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `json:"max-open-conns" yaml:"max-open-conns"` // 打开数据库的最大连接数
	LogMode      string `yaml:"log-mode"`                             // 是否开启Gorm全局日志
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config + "&parseTime=true"
}

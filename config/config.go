package config

type Server struct {
	Mail     Mail     `mapstructure:"mail" json:"mail" yaml:"mail"`
	Mailto   Mailto   `mapstructure:"mailto" json:"mailto" yaml:"mailto"`
	Dingding Dingding `mapstructure:"dingding" json:"dingding" yaml:"dingding"`
	Weixin   Weixin   `mapstructure:"weixin" json:"weixin" yaml:"weixin"`
	Image    Image    `mapstructure:"image" json:"image" yaml:"image"`
}

type Mail struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
}

type Mailto struct {
	Mails string `mapstructure:"mails" json:"mails" yaml:"mails"`
}

type Dingding struct {
	Webhook string `mapstructure:"webhook" json:"webhook" yaml:"webhook"`
}

type Weixin struct {
	Webhook string `mapstructure:"webhook" json:"webhook" yaml:"webhook"`
}

type Image struct {
	Base64 string `mapstructure:"base64" json:"base64" yaml:"base64"`
}

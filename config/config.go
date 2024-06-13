package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}
type Mysql struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Db        string `yaml:"db"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Log_level string `yaml:"og_level"` //日志等级，debug就是输出全部sql，dev，release
}
type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show-line"`      //是否显示行号
	LogInConsole bool   `yaml:"log-in-console"` //是否显示打印日志
}

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

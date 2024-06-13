package config

type Config struct {
	Mysql  Mysql  `json:"mysql"`
	Logger Logger `json:"logger"`
	System System `json:"system"`
}

type Mysql struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Db        string `json:"db"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Log_level string `json:"log_level"`
}

type Logger struct {
	Level        string `json:"level"`
	Prefix       string `json:"prefix"`
	Director     string `json:"director"`
	ShowLine     bool   `json:"show-line"`
	LogInConsole bool   `json:"log-in-console"`
}

type System struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Env  string `json:"env"`
}

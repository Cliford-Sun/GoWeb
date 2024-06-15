package config

type Config struct {
	Mysql  Mysql  `json:"mysql"`
	Logger Logger `json:"logger"`
	System System `json:"system"`
}

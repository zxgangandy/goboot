package logger

type Config struct {
	Level      string `mapstructure:"level"`
	Dir        string `mapstructure:"dir"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
}

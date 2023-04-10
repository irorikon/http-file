package config

type JWTConfig struct {
	Expire int `mapstructure:"expire" json:"expire" yaml:"expire"`
}

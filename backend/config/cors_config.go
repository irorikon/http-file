package config

type CORSConfig struct {
	Mode      string          `mapstructure:"mode" json:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whitelist" json:"whitelist" yaml:"whitelist"`
}
type CORSWhitelist struct {
	AllowedOrigin    string `mapstructure:"allowed_origin" json:"allowed_origin" yaml:"allowed_origin"`
	AllowedMethod    string `mapstructure:"allowed_method" json:"allowed_method" yaml:"allowed_method"`
	AllowedHeader    string `mapstructure:"allowed_header" json:"allowed_header" yaml:"allowed_header"`
	ExposeHeader     string `mapstructure:"expose_header" json:"expose_header" yaml:"expose_header"`
	AllowCredentials bool   `mapstructure:"allow_credentials" json:"allow_credentials" yaml:"allow_credentials"`
}

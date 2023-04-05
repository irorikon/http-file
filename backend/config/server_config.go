package config

type ServerConfig struct {
	Address  string `mapstructure:"address" json:"address" yaml:"address"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	TLS      bool   `mapstructure:"tls" json:"tls" yaml:"tls"`
	CertFile string `mapstructure:"cert_file" json:"cert_file" yaml:"cert_file"`
	KeyFile  string `mapstructure:"key_file" json:"key_file" yaml:"key_file"`
	Prefix   string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Search   bool   `mapstructure:"search" json:"search" yaml:"search"`
	Download bool   `mapstructure:"download" json:"download" yaml:"download"`
	Static   string `mapstructure:"static" json:"static" yaml:"static"`
	SiteURL  string `mapstructure:"site_url" json:"site_url" yaml:"site_url"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

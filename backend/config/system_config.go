package config

type SystemConfig struct {
	Title      string        `mapstructure:"title" json:"title" yaml:"title"`
	Logo       string        `mapstructure:"logo" json:"logo" yaml:"logo"`
	FooterText string        `mapstructure:"footer_text" json:"footer_text" yaml:"footer_text"`
	FooterLink string        `mapstructure:"footer_link" json:"footer_link" yaml:"footer_link"`
	MusicIMG   string        `mapstructure:"music_img" json:"music_img" yaml:"music_img"`
	Script     string        `mapstructure:"script" json:"script" yaml:"script"`
	AutoPlay   bool          `mapstructure:"autoplay" json:"autoplay" yaml:"autoplay"`
	Sort       string        `mapstructure:"sort" json:"sort" yaml:"sort"`
	Preview    PreviewConfig `mapstructure:"preview" json:"preview" yaml:"preview"`
	Expire     int           `mapstructure:"expire" json:"expire" yaml:"expire"`
}

type PreviewConfig struct {
	Text []string `mapstructure:"text" json:"text" yaml:"text"`
}

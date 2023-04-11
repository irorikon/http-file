/*
 * @Author: iRorikon
 * @Date: 2023-04-07 09:58:29
 * @FilePath: \http-file\backend\config\site_config.go
 */
/*
 * @Author: iRorikon
 * @Date: 2023-04-07 09:58:29
 * @FilePath: \http-file\backend\config\site_config.go
 */
package config

type SiteConfig struct {
	Title      string `mapstructure:"title" json:"title" yaml:"title"`
	Logo       string `mapstructure:"logo" json:"logo" yaml:"logo"`
	FooterText string `mapstructure:"footer_text" json:"footer_text" yaml:"footer_text"`
	FooterLink string `mapstructure:"footer_link" json:"footer_link" yaml:"footer_link"`
	Script     string `mapstructure:"script" json:"script" yaml:"script"`
	Sort       string `mapstructure:"sort" json:"sort" yaml:"sort"`
}

type PreviewConfig struct {
	Text []string `mapstructure:"text" json:"text" yaml:"text"`
}

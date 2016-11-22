package models

type SSLCerts struct {
	KeyFile    string `yaml:"key_file"`
	CertFile   string `yaml:"cert_file"`
	CACertFile string `yaml:"ca_file"`
}

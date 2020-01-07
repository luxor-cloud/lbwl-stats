package statsserver

type Config struct {
	Address   string    `json:"address"`
	SQLConfig SQLConfig `json:"sql,omitempty"`
	TLSKey    string    `json:"tls_key_path,omitempty"`
	TLSCert   string    `json:"tls_cert_path,omitempty"`
}

type SQLConfig struct {
	Address string `json:"address"`
}

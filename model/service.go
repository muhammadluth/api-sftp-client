package model

type (
	Properties struct {
		Port       string     `json:"port"`
		Timeout    string     `json:"timeout"`
		LogPath    string     `json:"log_path"`
		PoolSize   int        `json:"pool_size"`
		SftpConfig SftpConfig `json:"sftp_config"`
	}

	SftpConfig struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
	}
)

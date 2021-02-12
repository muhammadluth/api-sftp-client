package model

type (
	ResponseHttp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	ResponseSuccessWithData struct {
		TotalData int         `json:"total_data"`
		Data      interface{} `json:"data"`
	}
)

type (
	ResponseGetFileSFTP struct {
		Directory    string `json:"directory"`
		Filename     string `json:"filename"`
		Size         string `json:"size"`
		FileModified string `json:"file_modified"`
	}

	ResponseGetDirectorySFTP struct {
		Directory    string `json:"directory"`
		Size         string `json:"size"`
		FileModified string `json:"file_modified"`
	}
)

type (
	QueryParams struct {
		PathDirectory string `query:"path_directory"`
	}
)

// Code generated by goctl. DO NOT EDIT.
package types

type DataItem struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type DeleteRequest struct {
	Path string `json:"path"`
	Key  string `json:"key"`
}

type DeleteResponse struct {
	Successed bool `json:"successed"`
}

type QueryRequest struct {
	Path string `json:"path"`
}

type QueryResponse struct {
	Items []DataItem `json:"items"`
}

type UploadResponse struct {
	Path string `json:"path"`
}
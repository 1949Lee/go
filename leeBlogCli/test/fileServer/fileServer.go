package fileServer

// 文件信息格式
type File struct {
	// 文件id 由第三方库生成
	ID uint32 `json:"id"`

	// 文件名字
	Name string `json:"name"`

	// 文件大小，单位字节。
	Size int `json:"size"`

	// 文件最后修改时间，单位毫秒。
	LastModified int `json:"lastModified"`
}

// 接收到的文件的列表
type FileMap map[uint32]File

// TODO 前两个字节（16位）表示每个文件的第几个分片（最大共65536个分片）；后两个字节表示每个分片的大小（单位kb，默认64kb）。所以单个文件最大64kb * 2^16=4G

// 文件服务器：上传下载相关。
type FileServer struct {
	FileMap FileMap

	//FileFragment
}

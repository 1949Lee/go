package fileServer

// 文件信息格式
type File struct {
	Name         string
	Size         int
	LastModified int
}

// 接收到的文件的列表
type FileMap map[string]File

// TODO 前两个字节（16位）表示每个文件的第几个分片（最大共65536个分片）；后两个字节表示每个分片的大小（64kb）。所以单个文件最大64kb * 2^16=4G

// 文件服务器：上传下载相关。
type FileServer struct {
	FileMap FileMap

	//FileFragment
}

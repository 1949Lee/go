package definition

import (
	"bufio"
	"os"
)

type FileOptions struct {
	FileName  string
	ArticleID string
}
type FileStatusEnum struct {
	// 云端数据初始化失败
	InitFail ResponseCodeType

	// 文件上传进行中
	Uploading ResponseCodeType

	// 文件上传失败
	UploadFail ResponseCodeType

	// 文件上传完成
	UploadComplete ResponseCodeType

	// 云端文件保存中
	Processing ResponseCodeType

	// 云处理完成
	Success ResponseCodeType

	// 云端文件保存失败
	ProcessFail ResponseCodeType
}

var FileStatus = FileStatusEnum{
	InitFail:       301,
	Uploading:      302,
	UploadFail:     303,
	UploadComplete: 304,
	Processing:     305,
	ProcessFail:    306,
	Success:        0,
}

// 文件信息格式
type File struct {
	// 文件id 由第三方库生成
	ID uint32 `json:"id"`

	// 文件名字
	Name string `json:"name"`

	// 文件类型（扩展名）
	ExtType string `json:"extType"`

	// 文件大小，单位字节。
	Size int `json:"size"`

	// 文件最后修改时间，单位毫秒。
	LastModified int `json:"lastModified"`
}

// 文件片段
type FileFragment struct {
	// 文件ID
	FileID uint32

	// 文件片段编号
	FragmentIndex uint16

	// 文件片段数据
	FragmentData []byte

	// 是否是最后一个文件片段
	FileFragmentEnd bool
}

type FileInfo struct {
	// 文件ID
	ID uint32

	// 文件名
	Name string

	// 服务文件名 ***.png
	ServerName string

	// 文件类型（扩展名）
	ExtType string

	ServerFile *os.File

	BufIOWriter *bufio.Writer
}

// 接收到的文件的列表
type FileMap map[uint32]*FileInfo

type FileFragmentQueue []*FileFragment

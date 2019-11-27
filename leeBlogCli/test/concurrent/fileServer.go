package concurrent

import (
	"bufio"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

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
}

type FileInfo struct {
	// 文件ID
	ID uint32

	// 文件名称
	Name string

	// 文件类型（扩展名）
	ExtType string

	ServerFile *os.File

	BufIOWriter *bufio.Writer
}

// 接收到的文件的列表
type FileMap map[uint32]*FileInfo

// TODO 前两个字节（16位）表示每个文件的第几个分片（最大共65536个分片）；后两个字节表示每个分片的大小（单位kb，默认64kb）。所以单个文件最大64kb * 2^16=4G

// 文件服务器：上传下载相关。
type FileServer struct {
	// 处理中的文件映射
	FileMap FileMap

	// 文件片段管道
	FileFragmentChan chan *FileFragment

	//File
}

type FileFragmentQueue []*FileFragment

// 文件服务器
func (s *FileServer) Run(w *websocket.Conn) {
	go func() {
		//var fragmentQueue FileFragmentQueue
		for fragment := range s.FileFragmentChan {
			//if result.Type != 1 {
			//    if err := w.Conn.WriteJSON(result); err != nil {
			//        log.Printf("write err:%v", err)
			//    }
			//} else {
			//    fragmentQueue = append(fragmentQueue, fragment)
			//}

			info := s.FileMap[fragment.FileID]

			go func(fileInfo *FileInfo) {
				if fileInfo.ServerFile == nil {
					f, err := os.Create("./test.jpg")
					if err != nil {
						log.Printf("create file fail, the err:%v", err)
						if err := w.WriteJSON(ResponseResult{Type: 3, Code: FileStatus.InitFail}); err != nil {
							log.Printf("write err:%v", err)
						}
					}
					fileInfo.ServerFile = f
				} else {
					//if fragment. {
					//
					//}
					//defer f.Close()
					//
					//bw := bufio.NewWriter(f)
					//_, ok := bw.Write(p[4:])
					//if ok != nil {
					//	log.Println(ok)
					//}
					//yes := bw.Flush()
					//if yes != nil {
					//	log.Println(yes)
					//}
				}
			}(info)
		}
	}()
}

package concurrent

import (
	"bufio"
	"log"
	"os"
	"strings"
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

// TODO 上传的图片需要处理，渐进式显示图片，图片压缩。https://godoc.org/gopkg.in/gographics/imagick.v2/imagick、https://github.com/gographics/imagick、https://godoc.org/gopkg.in/gographics/imagick.v2/imagick、https://godoc.org/gopkg.in/gographics/imagick.v3/imagick
// TODO 前两个字节（16位）表示每个文件的第几个分片（最大共65536个分片）；后两个字节表示每个分片的大小（单位kb，默认64kb）。所以单个文件最大64kb * 2^16=4G

// 文件服务器：上传下载相关。
type FileServer struct {
	// 处理中的文件映射
	FileMap FileMap

	// 文件片段管道
	FileFragmentChan chan *FileFragment
}

type FileFragmentQueue []*FileFragment

// 文件服务器
func (s *FileServer) Run(w *Writer) {
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
				var builder strings.Builder
				builder.WriteString("./")
				builder.WriteString(fileInfo.ServerName)
				builder.WriteString(".lee")
				temName := builder.String()
				builder.Reset()
				builder.WriteString("./")
				builder.WriteString(fileInfo.ServerName)
				builder.WriteString(".")
				builder.WriteString(fileInfo.ExtType)
				serverName := builder.String()
				if fileInfo.ServerFile == nil {
					//filepath.Dir()
					f, err := os.Create(temName)
					if err != nil {
						log.Printf("create file fail, the err:%v", err)
						w.ResultChan <- &ResponseResult{Type: 3, Code: FileStatus.InitFail}
					}
					fileInfo.ServerFile = f
					fileInfo.BufIOWriter = bufio.NewWriter(f)
				}
				log.Printf("%d", fragment.FragmentIndex)
				_, ok := fileInfo.BufIOWriter.Write(fragment.FragmentData)
				if ok != nil {
					log.Printf("Write file fagment error, fileID:%d fileName:%s fragmentIndex:%d error is: %v \n", fileInfo.ID, fileInfo.Name, fragment.FragmentIndex, ok)
				}
				yes := fileInfo.BufIOWriter.Flush()
				if yes != nil {
					log.Println(yes)
				}
				// 表示文件数据到了最后。
				if fragment.FileFragmentEnd {
					err := os.Rename(temName, serverName)
					if err != nil {
						log.Printf("Rename file(%s) error, error is: %v \n", temName, err)
						delete(s.FileMap, fileInfo.ID)
						fileInfo.ServerFile.Close()
						if err := os.Remove(temName); err != nil {
							log.Printf("Remove file(%s) error, error is: %v \n", temName, err)
						}
						w.ResultChan <- &ResponseResult{Type: 3, Code: FileStatus.ProcessFail}
					}
					fileInfo.ServerFile.Close()
				}
			}(info)
		}
	}()
}

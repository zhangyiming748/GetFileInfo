package GetFileInfo

import (
	"github.com/zhangyiming748/log"
	"os"
	"path"
	"path/filepath"
)

type Info struct {
	FullPath string // 文件的绝对路径
	Size     int64  // 文件大小
	FullName string // 文件名
	ExtName  string // 扩展名
	Dir      string // 所在目录
}

const (
	MegaByte = 1000 * 1000 * 1000
)

/*
获取单个文件信息
*/
func GetFileInfo(absPath string) Info {
	mate, err := os.Stat(absPath)
	if err != nil {
		log.Warn.Printf("获取文件 %v 元数据发生错误 %v\n", absPath, err)
	}
	ext := path.Ext(absPath)
	dir, file := filepath.Split(absPath)
	i := Info{
		FullPath: absPath,
		Size:     mate.Size(),
		FullName: file,
		ExtName:  ext,
		Dir:      dir,
	}
	return i
}

/*
获取目录下符合条件的所有文件信息
*/
func GetAllFileInfo(dir, pattern string) {

}

/*
获取单个视频文件信息
*/
func GetVideoFileInfo(path string) {

}

/*
获取目录下符合条件的所有视频文件信息
*/
func GetAllVideoFileInfo(dir, pattern string) {

}

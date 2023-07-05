package GetFileInfo

import (
	"golang.org/x/exp/slog"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Info struct {
	FullPath string `json:"full_path,omitempty"` // 文件的绝对路径
	Size     int64  `json:"size,omitempty"`      // 文件大小
	FullName string `json:"full_name,omitempty"` // 文件名
	ExtName  string `json:"ext_name,omitempty"`  // 扩展名
	IsVideo  bool   `json:"is_video,omitempty"`  // 是否为视频文件
	Width    int    `json:"width,omitempty"`     // 视频宽度
	Height   int    `json:"height,omitempty"`    // 视频高度
	Code     string `json:"code,omitempty"`      // 视频编码
	VTag     string `json:"v_tag,omitempty"`     // 视频标签
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
		slog.Warn("获取文件元数据发生错误", absPath, err)
	}
	ext := path.Ext(absPath)
	_, file := filepath.Split(absPath)
	i := Info{
		FullPath: absPath,
		Size:     mate.Size(),
		FullName: file,
		ExtName:  ext,
		IsVideo:  false,
	}
	return i
}

/*
获取目录下符合条件的所有文件信息
*/
func GetAllFileInfo(dir, pattern string) []Info {
	var aim []Info
	files, err := os.ReadDir(dir)
	if err != nil {
		slog.Warn("出错", slog.Any("读取文件夹下内容", err))
		return nil
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			slog.Debug("获取文件信息", slog.String("跳过隐藏文件", file.Name()))
			continue
		}
		if file.IsDir() {
			slog.Debug("获取文件信息", slog.String("跳过文件夹", file.Name()))
			continue
		}
		currentExt := path.Ext(file.Name()) //当前文件的扩展名
		currentExt = strings.Replace(currentExt, ".", "", -1)
		if In(currentExt, strings.Split(pattern, ";")) {
			fullPath := strings.Join([]string{dir, file.Name()}, string(os.PathSeparator))
			mate, _ := os.Stat(fullPath)
			f := &Info{
				FullPath: fullPath,
				Size:     mate.Size(),
				FullName: file.Name(),
				ExtName:  currentExt,
			}
			aim = append(aim, *f)
		}
	}
	return aim
}

/*
获取目录下所有文件信息
*/
func GetEveryFileInfo(dir string) []Info {
	var aim []Info
	files, err := os.ReadDir(dir)
	if err != nil {
		slog.Warn("出错", slog.Any("读取文件夹下内容", err))
		return nil
	}
	for _, file := range files {
		if file.IsDir() {
			slog.Debug("获取文件信息", slog.String("跳过文件夹", file.Name()))
			continue
		}
		fullPath := strings.Join([]string{dir, file.Name()}, string(os.PathSeparator))
		f := &Info{
			FullPath: fullPath,
			FullName: file.Name(),
			ExtName:  path.Ext(fullPath),
		}
		aim = append(aim, *f)
	}
	return aim
}

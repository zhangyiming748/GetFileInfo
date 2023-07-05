package GetFileInfo

import (
	"golang.org/x/exp/slog"
	"os"
	"path"
	"path/filepath"
	"strings"
)

/*
获取单个视频文件信息
*/
func GetVideoFileInfo(absPath string) Info {
	mate, err := os.Stat(absPath)
	if err != nil {
		slog.Warn("获取文件元数据发生错误", absPath, err)
	}
	ext := path.Ext(absPath)
	_, file := filepath.Split(absPath)
	Code, Tag, Width, Height := getMediaInfo(absPath)
	i := Info{
		FullPath: absPath,
		Size:     mate.Size(),
		FullName: file,
		ExtName:  ext,
		IsVideo:  true,
		Code:     Code,
		Width:    Width,
		Height:   Height,
		VTag:     Tag,
	}
	return i
}

/*
获取目录下符合条件的所有视频文件信息
*/
func GetAllVideoFileInfo(dir, pattern string) []Info {
	files, err := os.ReadDir(dir)
	if err != nil {
		slog.Warn("错误", slog.Any("读取文件目录", err))
	}
	var aim []Info
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			slog.Debug("跳过隐藏文件", slog.Any("文件名", slog.AnyValue(file.Name())))
			continue
		}
		ext := path.Ext(file.Name())
		currentExt := path.Ext(file.Name()) //当前文件的扩展名
		currentExt = strings.Replace(currentExt, ".", "", -1)
		if In(currentExt, strings.Split(pattern, ";")) || currentExt == pattern {
			mate, _ := os.Stat(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
			Code, Tag, Width, Height := getMediaInfo(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
			f := &Info{
				FullPath: strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)),
				Size:     mate.Size(),
				FullName: file.Name(),
				ExtName:  ext,
				IsVideo:  true,
				Code:     Code,
				Width:    Width,
				Height:   Height,
				VTag:     Tag,
			}
			aim = append(aim, *f)
		}
	}
	return aim
}

/*
获取全部目录下符合条件的所有视频文件信息并生成报告
*/
type VideoReport struct {
	ref           string //文件名
	FileExtension string //扩展名
	container     string //容器
	VideoFormat   string //视频编码格式
	Width         string //视频宽度
	Height        string //视频高度
	AudioFormat   string //音频编码格式
}

package GetFileInfo

import (
	"fmt"
	"github.com/zhangyiming748/GetAllFolder"
	"golang.org/x/exp/slog"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

/*
获取单个视频文件信息
*/
func GetVideoFileInfo(absPath string) Info {
	mate, err := os.Stat(absPath)
	if err != nil {
		mylog.Warn("获取文件元数据发生错误", absPath, err)
	}
	ext := path.Ext(absPath)
	_, file := filepath.Split(absPath)
	Code, Width, Height := getMediaInfo(absPath)
	i := Info{
		FullPath: absPath,
		Size:     mate.Size(),
		FullName: file,
		ExtName:  ext,
		IsVideo:  true,
		Frame:    0,
		Code:     Code,
		Width:    Width,
		Height:   Height,
	}
	return i
}

/*
获取目录下符合条件的所有视频文件信息
*/
func GetAllVideoFileInfo(dir, pattern string) []Info {
	files, err := os.ReadDir(dir)
	if err != nil {
		mylog.Warn("错误", slog.Any("读取文件目录", err))
	}
	var aim []Info
	if strings.Contains(pattern, ";") {
		for _, file := range files {
			if strings.HasPrefix(file.Name(), ".") {
				mylog.Info("跳过隐藏文件", slog.Any("文件名", slog.AnyValue(file.Name())))
				continue
			}
			ext := path.Ext(file.Name())
			currentExt := path.Ext(file.Name()) //当前文件的扩展名
			currentExt = strings.Replace(currentExt, ".", "", -1)
			if In(currentExt, strings.Split(pattern, ";")) || currentExt == pattern {
				mate, _ := os.Stat(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
				Code, Width, Height := getMediaInfo(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
				var frame int
				f := &Info{
					FullPath: strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)),
					Size:     mate.Size(),
					FullName: file.Name(),
					ExtName:  ext,
					IsVideo:  true,
					Code:     Code,
					Width:    Width,
					Height:   Height,
					Frame:    frame,
				}
				aim = append(aim, *f)
			}
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

func (i *Info) SetFrame(frame string) {
	f, _ := strconv.Atoi(frame)
	i.Frame = f
	return
}

/*
视频文件的帧数有需要的情况下单独计算,默认为空
*/
func CountFrame(i *Info) {
	i.Frame = detectFrame(i.FullPath)
	return
}

/*
获取全部超过1080P的视频
*/
func GetOutOffFHD(dir, pattern string) (bigger []Info) {
	sum := 0
	infos := GetAllVideoFileInfo(dir, pattern)
	for _, info := range infos {
		if info.Width > 1920 && info.Height > 1920 {
			bigger = append(bigger, info)
			sum++
		}
	}
	slog.Info(fmt.Sprintf("共找到%d个大于FHD的视频", sum))
	return
}

/*
获取单个目录下全部非h265编码的视频
*/
func GetNotH265VideoFile(dir, pattern string) (h264 []Info) {
	sum := 0
	infos := GetAllVideoFileInfo(dir, pattern)
	for _, info := range infos {
		if info.Code != "HEVC" {
			sum++
			h264 = append(h264, info)
			mylog.Info("非H265视频", slog.Any("文件信息", h264))
		}
	}
	mylog.Info(fmt.Sprintf("共找到%d个非h265的视频", sum))
	return
}

/*
获取全部文件夹中非h265编码的视频
*/
func GetAllNotH265VideoFile(root, pattern string) (h264 []Info) {
	sum := 0
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		infos := GetNotH265VideoFile(folder, pattern)
		h264 = append(h264, infos...)
		sum++
	}
	return
}

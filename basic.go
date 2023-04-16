package GetFileInfo

import (
	"golang.org/x/exp/slog"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Info struct {
	FullPath string // 文件的绝对路径
	Size     int64  // 文件大小
	FullName string // 文件名
	ExtName  string // 扩展名
	IsVideo  bool   // 是否为视频文件
	Frame    int    // 视频帧数
	Width    int    // 视频宽度
	Height   int    // 视频高度
	Code     string // 视频编码
}

const (
	MegaByte = 1000 * 1000 * 1000
)

var mylog *slog.Logger

func SetLog(level string) {
	var opt slog.HandlerOptions
	switch level {
	case "Debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Info("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}

	}
	file := "GetFileInfo.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	mylog = slog.New(opt.NewJSONHandler(io.MultiWriter(logf, os.Stdout)))
}
func init() {
	l := os.Getenv("LEVEL")
	SetLog(l)
}

/*
获取单个文件信息
*/
func GetFileInfo(absPath string) Info {
	mate, err := os.Stat(absPath)
	if err != nil {
		mylog.Warn("获取文件元数据发生错误", absPath, err)
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
		mylog.Warn("出错", slog.Any("读取文件夹下内容", err))
		return nil
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			mylog.Info("跳过", slog.Any("隐藏文件", file.Name()))
			continue
		}
		if file.IsDir() {
			mylog.Info("跳过", slog.Any("文件夹", file.Name()))
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

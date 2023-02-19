package GetFileInfo

import (
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
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

/*
获取单个文件信息
*/
func GetFileInfo(absPath string) Info {
	mate, err := os.Stat(absPath)
	if err != nil {
		log.Warn.Printf("获取文件 %v 元数据发生错误 %v\n", absPath, err)
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
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Debug.Printf("读取文件目录产生的错误:%v\n", err)
	}
	var aim []Info
	if strings.Contains(pattern, ";") {
		exts := strings.Split(pattern, ";")
		for _, file := range files {
			if strings.HasPrefix(file.Name(), ".") {
				log.Info.Println("跳过隐藏文件", file.Name())
				continue
			}
			ext := path.Ext(file.Name())                 //文件扩展名
			justExt := strings.Replace(ext, ".", "", -1) //去掉点
			//log.Info.Printf("extname is %v\n", ext)
			for _, ex := range exts {
				if justExt == ex {
					//aim = append(aim, file.Name())
					mate, _ := os.Stat(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
					f := &Info{
						FullPath: strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)),
						Size:     mate.Size(),
						FullName: file.Name(),
						ExtName:  ext,
						IsVideo:  false,
					}
					aim = append(aim, *f)
				}
			}
		}
	} else {
		for _, file := range files {
			if strings.HasPrefix(file.Name(), ".") {
				log.Info.Println("跳过隐藏文件", file.Name())
				continue
			}
			ext := path.Ext(file.Name())
			justExt := strings.Replace(ext, ".", "", -1)
			//log.Info.Printf("extname is %v\n", ext)
			if justExt == pattern {
				//aim = append(aim, file.Name())
				mate, _ := os.Stat(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
				f := &Info{
					FullPath: strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)),
					Size:     mate.Size(),
					FullName: file.Name(),
					ExtName:  ext,
				}
				aim = append(aim, *f)
			}
		}
	}
	// log.Debug.Printf("有效的目标文件: %+v \n", aim)
	return aim
}

/*
获取单个视频文件信息
*/
func GetVideoFileInfo(absPath string) Info {
	mate, err := os.Stat(absPath)
	if err != nil {
		log.Warn.Printf("获取文件 %v 元数据发生错误 %v\n", absPath, err)
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
		log.Debug.Printf("读取文件目录产生的错误:%v\n", err)
	}
	var aim []Info
	if strings.Contains(pattern, ";") {
		exts := strings.Split(pattern, ";")
		for _, file := range files {
			if strings.HasPrefix(file.Name(), ".") {
				log.Info.Println("跳过隐藏文件", file.Name())
				continue
			}
			ext := path.Ext(file.Name())
			justExt := strings.Replace(ext, ".", "", -1)
			//log.Info.Printf("extname is %v\n", ext)
			for _, ex := range exts {
				if justExt == ex {
					//aim = append(aim, file.Name())
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
	} else {
		for _, file := range files {
			if strings.HasPrefix(file.Name(), ".") {
				log.Info.Println("跳过隐藏文件", file.Name())
				continue
			}
			ext := path.Ext(file.Name())
			justExt := strings.Replace(ext, ".", "", -1)
			//log.Info.Printf("extname is %v\n", ext)
			if justExt == pattern {
				//aim = append(aim, file.Name())
				mate, _ := os.Stat(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
				Code, Width, Height := getMediaInfo(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
				f := &Info{
					FullPath: strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)),
					Size:     mate.Size(),
					FullName: file.Name(),
					ExtName:  ext,
					IsVideo:  true,
					Code:     Code,
					Width:    Width,
					Height:   Height,
					Frame:    0,
				}
				aim = append(aim, *f)
			}
		}
	}
	// log.Debug.Printf("有效的目标文件: %+v \n", aim)
	return aim
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
		if info.Width > 1920 || info.Height > 1920 {
			bigger = append(bigger, info)
			sum++
		}
	}
	log.Debug.Printf("共找到%d个大于FHD的视频\n", sum)
	return
}
func MoveOutOffFHD(dir, pattern string) {
	target := strings.Join([]string{dir, "bigger"}, string(os.PathSeparator))
	os.Mkdir(target, 0777)
	files := GetOutOffFHD(dir, pattern)
	solve, err := os.OpenFile("solve.sh", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer solve.Close()
	if err != nil {
		return
	}
	for _, file := range files {
		src := strings.Join([]string{"\"", file.FullPath, "\""}, "")
		dst := strings.Join([]string{target, file.FullName}, string(os.PathSeparator))
		dst = strings.Join([]string{"\"", dst, "\""}, "")
		cmd := strings.Join([]string{"mv", src, dst}, " ")
		log.Debug.Printf("生成的单条命令:%s\n", cmd)
		solve.WriteString(cmd)
		solve.WriteString("\n")
	}
}
func MoveAllOutOffFHD(root, pattern string) {
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		MoveOutOffFHD(folder, pattern)
		runtime.GC()
	}
}

/*
获取全部非h265编码的视频
*/
func GetNotH265VideoFile(dir, pattern string) (h264 []Info) {
	sum := 0
	infos := GetAllVideoFileInfo(dir, pattern)
	for _, info := range infos {
		if info.Code != "HEVC" {
			sum++
			h264 = append(h264, info)
		}
	}
	log.Debug.Printf("共找到%d个非h265的视频\n", sum)
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
	log.Debug.Printf("共排查%d个文件夹\n", sum)
	return
}

func GetAllNotH265VideoFileReport(root, pattern string) {
	sum := 0
	var h264 []Info
	folders := GetAllFolder.ListFolders(root)
	for _, folder := range folders {
		infos := GetNotH265VideoFile(folder, pattern)
		h264 = append(h264, infos...)
		sum++
	}
	log.Debug.Printf("共排查%d个文件夹\n", sum)
	file, err := os.OpenFile("h264Report.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return
	}
	for _, v := range h264 {
		file.WriteString(strings.Join([]string{"\"", v.FullPath, "\"", "\n"}, ""))
	}
}

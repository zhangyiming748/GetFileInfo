package GetFileInfo

import (
	"github.com/zhangyiming748/log"
	"os"
	"path"
	"path/filepath"
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
			ext := path.Ext(file.Name())
			//log.Info.Printf("extname is %v\n", ext)
			for _, ex := range exts {
				if strings.Contains(ext, ex) {
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
			//log.Info.Printf("extname is %v\n", ext)
			if strings.Contains(ext, pattern) {
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
		Frame:    detectFrame(absPath),
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
			//log.Info.Printf("extname is %v\n", ext)
			for _, ex := range exts {
				if strings.Contains(ext, ex) {
					//aim = append(aim, file.Name())
					mate, _ := os.Stat(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
					Code, Width, Height := getMediaInfo(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
					var frame int
					go func() {
						// 随缘计算帧数,没时间等
						frame = detectFrame(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
					}()
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
			//log.Info.Printf("extname is %v\n", ext)
			if strings.Contains(ext, pattern) {
				//aim = append(aim, file.Name())
				mate, _ := os.Stat(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
				Code, Width, Height := getMediaInfo(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
				var frame int
				go func() {
					// 随缘计算帧数,没时间等
					frame = detectFrame(strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)))
				}()
				f := &Info{
					FullPath: strings.Join([]string{dir, file.Name()}, string(os.PathSeparator)),
					Size:     mate.Size(),
					FullName: file.Name(),
					ExtName:  ext,
					Code:     Code,
					Width:    Width,
					Height:   Height,
					Frame:    frame,
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

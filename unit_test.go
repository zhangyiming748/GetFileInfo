package GetFileInfo

import (
	"github.com/zhangyiming748/GetAllFolder"
	"github.com/zhangyiming748/pretty"
)

import "testing"

func TestGetFileInfo(t *testing.T) {
	absPath := "/Users/zen/Downloads/Telegram Desktop/Soul/1 (49).jpg"
	ret := GetFileInfo(absPath)
	t.Logf("%+v\n", ret)
}

func TestGetAllFileInfo(t *testing.T) {
	absPath := "D:\\甄嬛传\\h264"
	ret := GetAllFileInfo(absPath, "mp4")
	t.Logf("%+v\n", ret)
}
func TestGetVideoFileInfo(t *testing.T) {
	absPath := ""
	ret := GetVideoFileInfo(absPath)
	t.Logf("%+v\n", ret)
}
func TestGetAllVideoFileInfo(t *testing.T) {
	absPath := ""
	ret := GetAllVideoFileInfo(absPath, "avi")
	t.Logf("%+v\n", ret)
}

func TestGetNotH265VideoFile(t *testing.T) {
	dir := "/Users/zen/Movies"
	folders := GetAllFolder.ListFolders(dir)
	for _, folder := range folders {
		ret := GetNotH265VideoFile(folder, "mp4")
		pretty.P(ret)
	}

}

func TestGetAllNotH265VideoFile(t *testing.T) {
	root := "/Volumes/T7/slacking/Telegram/DOA"
	pattern := "webm;mkv;m4v;MP4;mp4;mov;avi;wmv;ts;TS;rmvb"
	GetAllNotH265VideoFileReport(root, pattern)

}

func TestGetAllOutOffFHDVideoFileReport(t *testing.T) {
	root := "/Volumes/T7/slacking/Telegram"
	pattern := "webm;mkv;m4v;MP4;mp4;mov;avi;wmv;ts;TS;rmvb"
	GetAllOutOffFHDVideoFileReport(root, pattern)
}
func TestGetGeneralMediaInfo(t *testing.T) {
	ret := getGeneralMediaInfo("/Users/zen/Downloads/整理/live/5_6086967526091131008.mp4")
	pretty.P(ret)
}
func TestGetAllVideoFilesInfoReport(t *testing.T) {
	root := "/Users/zen/Downloads/整理"
	pattern := "mp4"
	GetAllVideoFilesInfoReport(root, pattern)
}

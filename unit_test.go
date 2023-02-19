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
	absPath := ""
	ret := GetAllFileInfo(absPath, "jpg")
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
func TestMoveOutOffFHD(t *testing.T) {
	absPath := "/Volumes/T7/slacking/Telegram/Frozen/Elsa/h265"
	MoveOutOffFHD(absPath, "mp4")
}
func TestGetH265VideoFile(t *testing.T) {
	dir := "/Users/zen/Movies"
	folders := GetAllFolder.ListFolders(dir)
	for _, folder := range folders {
		ret := GetH265VideoFile(folder, "mp4")
		pretty.P(ret)
	}

}
func TestMoveAllOutOffFHD(t *testing.T) {
	root := "/Volumes/T7/slacking"
	pattern := "webm;mkv;m4v;MP4;mp4;mov;avi;wmv;ts;TS;rmvb"
	MoveAllOutOffFHD(root, pattern)
}

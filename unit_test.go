package GetFileInfo

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

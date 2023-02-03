package GetFileInfo

import "testing"

func TestGetFileInfo(t *testing.T) {
	absPath := "/Users/zen/Downloads/Telegram Desktop/Soul/1 (49).jpg"
	ret := GetFileInfo(absPath)
	t.Logf("%+v\n", ret)
}

func TestGetAllFileInfo(t *testing.T) {
	absPath := "/Users/zen/Downloads/Telegram Desktop/OverWatch"
	ret := GetAllFileInfo(absPath, "jpg")
	t.Logf("%+v\n", ret)
}
func TestGetVideoFileInfo(t *testing.T) {
	absPath := "/Users/zen/Downloads/Telegram Desktop/Mass Effect/HD liara (Miranda Lawson x Liara T'Soni) [1920 × 1080 - 60 FPS].mp4"
	ret := GetVideoFileInfo(absPath)
	t.Logf("%+v\n", ret)
}

package GetFileInfo

import "testing"

func TestGetNotH265VideoFile(t *testing.T) {
	absPath := "/Users/zen/Downloads/Telegram Desktop/tele"
	ret := GetAllNotH265VideoFile(absPath, "mp4;mp3")
	t.Logf("%+v\n", ret)
}

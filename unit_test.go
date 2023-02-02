package GetFileInfo

import "testing"

func TestGetFileInfo(t *testing.T) {
	absPath := "/Users/zen/Downloads/Telegram Desktop/Soul/1 (49).jpg"
	ret := GetFileInfo(absPath)
	t.Logf("%+v\n", ret)
}

package GetFileInfo

import (
	"github.com/zhangyiming748/pretty"
	"testing"
)

func TestGetNotH265VideoFile(t *testing.T) {
	absPath := "/Users/zen/Downloads"
	ret := GetAllFileInfo(absPath, "mp4;mp3;avif;png;jpg", "Debug")
	t.Logf("%+v\n", ret)
}
func TestGetAllVideoFileInfo(t *testing.T) {
	ret := GetAllVideoFileInfo("/Users/zen/Downloads/NecDaz/ff/tifa/resize", "mp4;avi", "Debug")
	pretty.P(ret)
}

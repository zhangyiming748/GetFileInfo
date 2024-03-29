package GetFileInfo

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	str1 := "mp4"
	str2 := "mp4;avi"
	s1 := strings.Split(str1, ";")
	s2 := strings.Split(str2, ";")
	for _, v1 := range s1 {
		t.Logf("str1 len is %v\tthis value is %v\n", len(s1), v1)
	}
	for _, v2 := range s2 {
		t.Logf("str2 len is %v\tthis value is %v\n", len(s2), v2)
	}
}
func TestFind(t *testing.T) {

}

func TestGetAllFast(t *testing.T) {
	fast, err := GetAllKindsOfFilesInfoFast("F:\\large\\GirlFriend4ever", []string{"jpg", "png"})
	if err != nil {
		return
	}

	for _, v := range fast {
		t.Logf("%v\n", v)
	}
}

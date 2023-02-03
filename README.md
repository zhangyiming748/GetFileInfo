# GetFileInfo
获取指定文件夹中的特定文件信息


# 返回结构

```go
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
```

# 入参
单个文件使用文件绝对路径
多个文件使用目录绝对路径和扩展名
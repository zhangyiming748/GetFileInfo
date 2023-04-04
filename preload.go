package GetFileInfo

import (
	"encoding/json"
	"fmt"
	"github.com/zhangyiming748/replace"
	"golang.org/x/exp/slog"
	"io"
	"os/exec"
	"strconv"
	"sync"
)

type MediaInfo struct {
	CreatingLibrary struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Url     string `json:"url"`
	} `json:"creatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type                  string `json:"@type"` // Video
			VideoCount            string `json:"VideoCount,omitempty"`
			FileExtension         string `json:"FileExtension,omitempty"`
			Format                string `json:"Format"`
			FormatProfile         string `json:"Format_Profile"`
			CodecID               string `json:"CodecID"`
			CodecIDCompatible     string `json:"CodecID_Compatible,omitempty"`
			FileSize              string `json:"FileSize,omitempty"`
			Duration              string `json:"Duration"`
			OverallBitRateMode    string `json:"OverallBitRate_Mode,omitempty"`
			OverallBitRate        string `json:"OverallBitRate,omitempty"`
			FrameRate             string `json:"FrameRate"`
			FrameCount            string `json:"FrameCount"`
			StreamSize            string `json:"StreamSize"`
			HeaderSize            string `json:"HeaderSize,omitempty"`
			DataSize              string `json:"DataSize,omitempty"`
			FooterSize            string `json:"FooterSize,omitempty"`
			IsStreamable          string `json:"IsStreamable,omitempty"`
			EncodedDate           string `json:"Encoded_Date"`
			TaggedDate            string `json:"Tagged_Date"`
			FileModifiedDate      string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal string `json:"File_Modified_Date_Local,omitempty"`
			Extra                 struct {
				TIM                   string `json:"TIM,omitempty"`
				TSC                   string `json:"TSC,omitempty"`
				TSZ                   string `json:"TSZ,omitempty"`
				CodecConfigurationBox string `json:"CodecConfigurationBox,omitempty"`
			} `json:"extra"`
			StreamOrder                    string `json:"StreamOrder,omitempty"`
			ID                             string `json:"ID,omitempty"`
			FormatLevel                    string `json:"Format_Level,omitempty"`
			FormatSettingsCABAC            string `json:"Format_Settings_CABAC,omitempty"`
			FormatSettingsRefFrames        string `json:"Format_Settings_RefFrames,omitempty"`
			BitRateMode                    string `json:"BitRate_Mode,omitempty"`
			BitRate                        string `json:"BitRate,omitempty"`
			Width                          string `json:"Width,omitempty"`
			Height                         string `json:"Height,omitempty"`
			StoredWidth                    string `json:"Stored_Width,omitempty"`
			SampledWidth                   string `json:"Sampled_Width,omitempty"`
			SampledHeight                  string `json:"Sampled_Height,omitempty"`
			PixelAspectRatio               string `json:"PixelAspectRatio,omitempty"`
			DisplayAspectRatio             string `json:"DisplayAspectRatio,omitempty"`
			Rotation                       string `json:"Rotation,omitempty"`
			FrameRateMode                  string `json:"FrameRate_Mode,omitempty"`
			ColorSpace                     string `json:"ColorSpace,omitempty"`
			ChromaSubsampling              string `json:"ChromaSubsampling,omitempty"`
			BitDepth                       string `json:"BitDepth,omitempty"`
			ScanType                       string `json:"ScanType,omitempty"`
			Language                       string `json:"Language,omitempty"`
			BufferSize                     string `json:"BufferSize,omitempty"`
			ColourDescriptionPresent       string `json:"colour_description_present,omitempty"`
			ColourDescriptionPresentSource string `json:"colour_description_present_Source,omitempty"`
			ColourRange                    string `json:"colour_range,omitempty"`
			ColourRangeSource              string `json:"colour_range_Source,omitempty"`
			ColourPrimaries                string `json:"colour_primaries,omitempty"`
			ColourPrimariesSource          string `json:"colour_primaries_Source,omitempty"`
			TransferCharacteristics        string `json:"transfer_characteristics,omitempty"`
			TransferCharacteristicsSource  string `json:"transfer_characteristics_Source,omitempty"`
			MatrixCoefficients             string `json:"matrix_coefficients,omitempty"`
			MatrixCoefficientsSource       string `json:"matrix_coefficients_Source,omitempty"`
		} `json:"track"`
	} `json:"media"`
}

func getGeneralMediaInfo(absPath string) MediaInfo {
	var mi MediaInfo
	cmd := exec.Command("mediainfo", absPath, "--Output=JSON")
	mylog.Info("生成的命令", slog.String("命令", fmt.Sprint(cmd)))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		mylog.Warn("cmd.StdoutPipe", slog.String("产生的错误", fmt.Sprint(err)))
		return MediaInfo{}
	}
	if err = cmd.Start(); err != nil {
		mylog.Warn("cmd.Run", slog.String("产生的错误", fmt.Sprint(err)))
		return MediaInfo{}
	}

	//读取所有输出
	bytes, err := io.ReadAll(stdout)
	if err != nil {
		mylog.Warn("ReadAll Stdout", slog.String("产生的错误", fmt.Sprint(err)))
		return MediaInfo{}
	} else {
		//log.Debug.Printf("命令输出内容:%v\n", string(bytes))
		if err := json.Unmarshal(bytes, &mi); err != nil {
			mylog.Warn("解析json", slog.String("产生的错误", fmt.Sprint(err)))
		}
	}

	if err = cmd.Wait(); err != nil {
		mylog.Warn("命令执行中", slog.String("产生的错误", fmt.Sprint(err)))
	}
	return mi
}
func getMediaInfo(absPath string) (Code string, Width, Height int) {
	var mi MediaInfo
	cmd := exec.Command("mediainfo", absPath, "--Output=JSON")
	mylog.Info("生成的命令", slog.String("命令", fmt.Sprint(cmd)))
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		mylog.Warn("cmd.StdoutPipe", slog.String("产生的错误", fmt.Sprint(err)))
	}
	if err = cmd.Start(); err != nil {
		mylog.Warn("cmd.Run", slog.String("产生的错误", fmt.Sprint(err)))
	}

	//读取所有输出
	bytes, err := io.ReadAll(stdout)
	if err != nil {
		mylog.Warn("ReadAll Stdout", slog.String("产生的错误", fmt.Sprint(err)))

	} else {
		//log.Debug.Printf("命令输出内容:%v\n", string(bytes))
		if err := json.Unmarshal(bytes, &mi); err != nil {
			mylog.Warn("解析json", slog.String("产生的错误", fmt.Sprint(err)))
		}
	}

	if err = cmd.Wait(); err != nil {
		mylog.Warn("命令执行中有错误产生", err)
	}
	for _, video := range mi.Media.Track {
		if video.Type == "Video" {
			w, _ := strconv.Atoi(video.Width)
			h, _ := strconv.Atoi(video.Height)
			Code = video.Format
			Width = w
			Height = h
		}
	}
	return
}

/*
输出文件全部帧数
*/
func detectFrame(absPath string) int {
	cmd := exec.Command("ffprobe", "-v", "error", "-count_frames", "-select_streams", "v:0", "-show_entries", "stream=nb_read_frames", "-of", "default=nokey=1:noprint_wrappers=1", absPath)
	/*
		> -v error:这隐藏了"info"输出(版本信息等),使解析更容易.
		> -count_frames:计算每个流的帧数,并在相应的流部分中报告.
		> -select_streams v:0 :仅选择视频流.
		> -show_entries stream = nb_read_frames :只显示读取的帧数.
		> -of default = nokey = 1:noprint_wrappers = 1 :将输出格式(也称为"writer")设置为默认值,不打印每个字段的键(nokey = 1),不打印节头和页脚(noprint_wrappers = 1).
	*/
	mylog.Info("生成的命令", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		mylog.Warn("cmd.StdoutPipe", slog.String("产生的错误", fmt.Sprint(err)))
		return 0
	}
	if err = cmd.Start(); err != nil {
		mylog.Warn("cmd.Run产生的错误", err)
		return 0
	}
	tmp := make([]byte, 1024)
	stdout.Read(tmp)
	t := string(tmp)
	t = replace.Replace(t)
	if atoi, err := strconv.Atoi(t); err == nil {
		mylog.Info("文件帧数", absPath, atoi)
		return atoi
	}
	mylog.Warn("读取文件帧数出错")
	return 0
}

/*
输出文件全部帧数,wg防止程序提前退出
*/
func detectFrameWithWaitGroup(absPath string, wg *sync.WaitGroup) int {
	cmd := exec.Command("ffprobe", "-v", "error", "-count_frames", "-select_streams", "v:0", "-show_entries", "stream=nb_read_frames", "-of", "default=nokey=1:noprint_wrappers=1", absPath)
	/*
		> -v error:这隐藏了"info"输出(版本信息等),使解析更容易.
		> -count_frames:计算每个流的帧数,并在相应的流部分中报告.
		> -select_streams v:0 :仅选择视频流.
		> -show_entries stream = nb_read_frames :只显示读取的帧数.
		> -of default = nokey = 1:noprint_wrappers = 1 :将输出格式(也称为"writer")设置为默认值,不打印每个字段的键(nokey = 1),不打印节头和页脚(noprint_wrappers = 1).
	*/
	mylog.Info("生成的命令", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		mylog.Warn("cmd.StdoutPipe产生的错误", slog.String("产生的错误", fmt.Sprint(err)))
		return 0
	}
	if err = cmd.Start(); err != nil {
		mylog.Warn("cmd.Run", slog.String("产生的错误", fmt.Sprint(err)))
		return 0
	}
	tmp := make([]byte, 1024)
	stdout.Read(tmp)
	t := string(tmp)
	t = replace.Replace(t)
	if atoi, err := strconv.Atoi(t); err == nil {
		wg.Done()
		return atoi
	}
	mylog.Warn("读取文件帧数出错", slog.String("产生错误的文件", absPath))
	wg.Done()
	return 0
}
func In(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

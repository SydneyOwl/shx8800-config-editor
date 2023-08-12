package filetools

import (
	"encoding/json"
	"errors"
	"github.com/sydneyowl/shx8800-config-editor/internal/radio"
	"os"
	"os/exec"
	"runtime"
)

func IsFileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
func GetCSPath() string {
	pathExe := ""
	switch runtime.GOOS {
	case "windows":
		pathExe = "./SHX8800.exe"
	case "linux", "darwin":
		pathExe = "./SHX8800"
	default:
		return ""
	}
	return pathExe
}

func ParseData(exePath string, configPath string) (*radio.ClassTheRadioData, error) {
	cmd := exec.Command(exePath, "transjson", configPath)
	out, err := cmd.CombinedOutput()
	if err != nil || len(out) < 10 {
		return nil, errors.New("执行失败！")
	}
	configs := &radio.ClassTheRadioData{
		ChannelData:  [128][]string{},
		FunCfgData:   radio.CreateFFCData(),
		DtmfData:     radio.CreateDTMFData(),
		OtherImfData: radio.CreateOIMFData(),
	}
	err = json.Unmarshal(out, configs)
	if err != nil {
		return nil, err
	}
	return configs, nil
}

func WriteBackData(exePath string, configPath string, jsonstring string) error {
	cmd := exec.Command(exePath, "transfile", configPath, jsonstring)
	out, err := cmd.CombinedOutput()
	if err != nil || len(out) != 0 {
		return errors.New("执行失败！")
	}
	return nil
}

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"github.com/sydneyowl/shx8800-config-editor/internal/radio"
	"github.com/sydneyowl/shx8800-config-editor/pkg/filetools"
	"os"
	"path/filepath"
)

func mainRunner() {
	var configPath string
	var configs *radio.ClassTheRadioData
	if ConfigPath == "" {
		fmt.Print("请输入.dat文件路径:")
		_, _ = fmt.Scanln(&configPath)
	} else {
		configPath = ConfigPath
	}
	slog.Info("检查文件完整性...")
	if !filetools.IsFileExist(configPath) {
		slog.Fatal("文件路径有误！")
		return
	}
	pathExe := filetools.GetCSPath()
	if DepPath == "" {
		if !filetools.IsFileExist(pathExe) {
			err := filetools.ReleaseFile()
			if err != nil {
				slog.Fatalf("无法释放依赖！")
				slog.Debug(err)
				return
			}
			slog.Info("释放成功！请重启程序，按下回车确认")
			_, _ = fmt.Scanln()
			return
		}
	} else {
		pathExe = DepPath
	}
	configs, err := filetools.ParseData(pathExe, configPath)
	if err != nil {
		slog.Errorf("解析失败：%v", err)
		return
	} else {
		slog.Info("解析成功")
	}
	for {
		//C  C
	swh:
		printMenu()
		choose := "q"
		_, _ = fmt.Scanln(&choose)
		switch choose {
		case "m":
			for {
			wait:
				printSubMenu()
				var choice string
				_, _ = fmt.Scanln(&choice)
				for {
					switch choice {
					case "p":
						displayChannels((*configs).ChannelData)
						goto wait
					case "b":
						goto swh
					case "c":
						fmt.Printf("请输入要清空的信道号")
						channel := "e"
						_, _ = fmt.Scanln(&channel)
						err = ClearChannel(configs, channel)
						if err != nil {
							slog.Warn("无法清空：", err)
						} else {
							slog.Info("清空成功！")
						}
						goto wait
					case "k":
						RemoveAllEmpty(configs)
						slog.Info("删除成功！")
						goto wait
					case "e":
						fmt.Print("请输入要修改的信道号：")
						chanNo := "e"
						_, _ = fmt.Scanln(&chanNo)
						displayModify()
						modNo := 0
						fmt.Print("请输入要修改的选项：")
						_, _ = fmt.Scanln(&modNo)
						fmt.Print("要改成什么？")
						rewriteTo := "e"
						_, _ = fmt.Scanln(&rewriteTo)
						err := ModifyChannel(configs, chanNo, modNo, rewriteTo)
						if err != nil {
							slog.Error(err)
						} else {
							slog.Info("成功！")
						}
						goto wait
					case "i":
						fmt.Printf("请输入要在哪一个信道前插入信道（插入信道号）")
						channel := "e"
						_, _ = fmt.Scanln(&channel)
						targetArr, err := AskForDetail()
						if err != nil {
							slog.Info("用户取消！")
							goto wait
						}
						targetArr[0] = channel
						err = AddChannel(configs, channel, *targetArr)
						if err != nil {
							slog.Warn("错误：", err)
						} else {
							slog.Info("插入成功！")
						}
						goto wait
					case "m":
						channel := "e"
						fmt.Print("请输入要复制的信道：")
						_, _ = fmt.Scanln(&channel)
						err = copyChannel(*configs, channel)
						if err != nil {
							slog.Warn("错误：", err)
						} else {
							slog.Info("复制成功！")
						}
						goto wait
					case "d":
						ShowCopiedChannel()
						goto wait
					case "v":
						fmt.Print("复制缓冲区中的哪个信道？")
						from := "e"
						_, _ = fmt.Scanln(&from)
						fmt.Print("要复制到哪个信道？")
						dst := "e"
						_, _ = fmt.Scanln(&dst)
						err = pasteChannelWithCheck(configs, from, dst)
						if err != nil {
							slog.Warn("错误：", err)
						} else {
							slog.Info("粘贴成功！")
						}
						goto wait
					case "a":
						fmt.Println(configs.ChannelData)
						goto wait
					default:
						fmt.Println("输入错误，请重新输入")
						_, _ = fmt.Scanln(&choice)
					}
				}
			}
		case "q":
			os.Exit(0)
		case "s":
			jsonstring, err := json.Marshal(configs)
			out := jsonstring
			//out, _ := simplifiedchinese.GB18030.NewEncoder().Bytes(jsonstring)
			if err != nil {
				slog.Errorf("无法反序列：%v", err)
				return
			}
			fmt.Print("请输入要保存到的目标目录")
			saveto := "./"
			_, _ = fmt.Scanln(&saveto)
			if !filetools.IsFileExist(saveto) {
				slog.Warn("文件路径有误！")
				continue
			}
			id := uuid.New().String()[0:12] + ".dat"
			saveto = filepath.Join(saveto, id)
			slog.Trace(saveto)
			err = filetools.WriteBackData(pathExe, saveto, string(out))
			if err != nil {
				slog.Errorf("无法保存：%v", err)
				return
			}
			slog.Infof("保存成功！文件名为：%s", saveto)
		default:
			fmt.Println("输入错误！")
		}
	}
}

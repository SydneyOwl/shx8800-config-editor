package entrance

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"github.com/sydneyowl/shx8800-config-editor/internal/radio"
	"github.com/sydneyowl/shx8800-config-editor/pkg/filetools"
	"os"
	"path/filepath"
	"strings"
)

func mainRunner() {
	var specifyConfigPath string
	var configs *radio.ClassTheRadioData
	pathExe := filetools.GetCSPath()
	if DepPath == "" {
		if filetools.Build == "dev" {
			slog.Fatalf("未指定依赖包位置！请使用--dep-path指定！")
			_, _ = fmt.Scanln()
			return
		}
		if !filetools.IsFileExist(pathExe) {
			err := filetools.ReleaseFile()
			if err != nil {
				slog.Fatalf("无法释放依赖！")
				_, _ = fmt.Scanln()
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
	if ConfigPath == "" {
		fmt.Print("请输入.dat文件路径:")
		reader := bufio.NewReader(os.Stdin)
		specifyConfigPath, _ = reader.ReadString('\n')
		specifyConfigPath = strings.TrimSpace(specifyConfigPath)
	} else {
		specifyConfigPath = ConfigPath
	}
	specifyConfigPath = strings.Trim(specifyConfigPath, "\"")
	slog.Info("检查文件完整性...")
	if !filetools.IsFileExist(specifyConfigPath) {
		slog.Fatal("文件路径有误！")
		_, _ = fmt.Scanln()
		return
	}
	if !filetools.IsFileExist(pathExe) {
		slog.Fatal("C#文件路径有误！")
		_, _ = fmt.Scanln()
		return
	}
	configs, err := filetools.ParseData(pathExe, specifyConfigPath)
	if err != nil {
		slog.Errorf("解析失败：%v", err)
		return
	} else {
		slog.Info("解析成功")
	}
	for {
	swh:
		lastStep := make([]radio.ClassTheRadioData, 0)
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
					case "r":
						if len(lastStep) == 0 {
							slog.Warn("无法撤销！")
							goto wait
						}
						backto := lastStep[len(lastStep)-1]
						*configs = backto
						lastStep = lastStep[:len(lastStep)-1]
						slog.Notice("已经撤销")
						goto wait
					case "p":
						displayChannels((*configs).ChannelData)
						goto wait
					case "b":
						goto swh
					case "c":
						backupConfig(configs, &lastStep)
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
					case "z":
						backupConfig(configs, &lastStep)
						fmt.Printf("请输入要清空的信道范围，例如1,2-4,9代表清空1,2,3,4,9: ")
						channel := "e"
						_, _ = fmt.Scanln(&channel)
						channelRangesd, err := parseChannel(channel)
						if err != nil {
							slog.Warn("无法解析：", err)
							goto wait
						}
						err = BatchClear(configs, channelRangesd)
						if err != nil {
							slog.Warn("无法清空：", err)
						} else {
							slog.Info("清空成功！")
						}
						goto wait
					case "k":
						backupConfig(configs, &lastStep)
						RemoveAllEmpty(configs)
						slog.Info("删除成功！")
						goto wait
					case "e":
						backupConfig(configs, &lastStep)
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
						backupConfig(configs, &lastStep)
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
						backupConfig(configs, &lastStep)
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
			fmt.Print("请输入要保存到的目标目录，输入英文句号（.）保存到当前目录：")
			saveto := "./"
			reader := bufio.NewReader(os.Stdin)
			saveto, _ = reader.ReadString('\n')
			saveto = strings.TrimSpace(saveto)
			saveto = strings.Trim(saveto, "\"")
			if !filetools.IsFileExist(saveto) {
				slog.Warn("文件路径有误！")
				continue
			}
			id := uuid.New().String()[0:12] + ".dat"
			saveto = filepath.Join(saveto, id)
			err = filetools.WriteBackData(pathExe, saveto, string(out))
			if err != nil {
				slog.Errorf("无法保存：%v", err)
				continue
			}
			slog.Infof("保存成功！文件名为：%s", saveto)
		default:
			fmt.Println("输入错误！")
		}
	}
}

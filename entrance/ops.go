package entrance

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sydneyowl/shx8800-config-editor/internal/radio"
	"os"
	"strconv"
	"strings"
)

var copiedChannel = make([][]string, 0)

func copyChannel(configs radio.ClassTheRadioData, copyChan string) error {
	for i, v := range configs.ChannelData {
		if v[0] == copyChan {
			cache := make([]string, 14)
			copy(cache, configs.ChannelData[i])
			cache[0] = uuid.New().String()[0:3]
			copiedChannel = append(copiedChannel, cache)
			return nil
		}
	}
	return errors.New("未找到指定信道！")
}
func pasteChannel(configs *radio.ClassTheRadioData, from string, dst string) error {
	for i, v := range configs.ChannelData {
		if v[0] == dst {
			for j, k := range copiedChannel {
				if k[0] == from {
					configs.ChannelData[i] = copiedChannel[j]
					configs.ChannelData[i][0] = dst
					return nil
				}
			}
			return errors.New("缓冲区不存在该信道！")
		}
	}
	return errors.New("未找到指定信道！")
}
func pasteChannelWithCheck(configs *radio.ClassTheRadioData, from string, dst string) error {
	if num, err := strconv.Atoi(dst); err != nil {
		return errors.New("输入错误！")
	} else if !(num >= 0 && num <= 128) {
		return errors.New("范围错误！")
	} else {
		if err := pasteChannel(configs, from, dst); err != nil {
			return err
		}
		return nil
	}
}
func ShowCopiedChannel() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	tmp := make([]interface{}, len(available_choices))
	for i := range available_choices {
		tmp[i] = available_choices[i]
	}
	tmp[0] = "序号"
	t.AppendHeader(tmp)
	for i := range copiedChannel {
		tmp := make([]interface{}, len(copiedChannel[i]))
		for j := range copiedChannel[i] {
			tmp[j] = copiedChannel[i][j]
		}
		t.AppendRow(tmp)
	}
	t.SetStyle(table.StyleColoredBright)
	t.Render()
}

func isChannelEmpty(target []string) bool {
	for i, v := range target {
		if i != 0 {
			if v != "" {
				return false
			}
		}
	}
	return true
}

func printMenu() {
	fmt.Println("请选择操作：")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"操作码", "操作"})
	t.AppendRows([]table.Row{
		{"m", "开始修改信道"},
		{"q", "丢弃所有更改并退出"},
		{"s", "保存更改到文件"},
	})
	t.Render()
}

func printSubMenu() {
	fmt.Println("请选择操作：")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"操作码", "操作"})
	t.AppendRows([]table.Row{
		{"r", "撤销上一步更改(测试功能)"},
		{"p", "打印信道"},
		{"c", "清空信道"},
		{"z", "批量清空信道(测试功能)"},
		{"i", "插入信道"},
		{"e", "更改信道"},
		{"k", "删除所有空信道"},
		{"m", "复制信道到缓冲区"},
		{"v", "粘贴信道"},
		{"d", "展示缓冲区中的信道"},
		{"a", "打印原始数据(测试功能)"},
		{"b", "回到主界面"},
	})
	t.Render()
}
func displayChannels(ChannelData [128][]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	tmp := make([]interface{}, len(available_choices))
	for i := range available_choices {
		tmp[i] = available_choices[i]
	}
	t.AppendHeader(tmp)
	var processedChan = make([][]string, 0)
	emptyCount := 0
	start := -1
	for i, v := range ChannelData {
		if isChannelEmpty(v) {
			emptyCount += 1
			if start == -1 {
				start = i
			}
			if i == len(ChannelData)-1 {
				processedChan = append(processedChan, []string{"跳过信道" + strconv.Itoa(start) + "-" + strconv.Itoa(len(ChannelData)-1)})
			}
		} else {
			if emptyCount != 0 {
				if emptyCount >= SKIP_EMPTY_CHANNEL_COUNT {
					processedChan = append(processedChan, []string{"跳过信道" + strconv.Itoa(start) + "-" + strconv.Itoa(i-1)})
				} else {
					for h := start; h < i; h++ {
						processedChan = append(processedChan, ChannelData[h])
					}
				}
				start = -1
				emptyCount = 0
			}
			processedChan = append(processedChan, ChannelData[i])
		}
	}
	for i := range processedChan {
		tmp := make([]interface{}, len(processedChan[i]))
		for j := range processedChan[i] {
			tmp[j] = processedChan[i][j]
		}
		t.AppendRow(tmp)
	}
	t.SetStyle(table.StyleColoredBright)
	t.Render()
}
func ClearChannel(configs *radio.ClassTheRadioData, chanNo string) error {
	for i, v := range configs.ChannelData {
		if v[0] == chanNo {
			configs.ChannelData[i] = make([]string, 14)
			configs.ChannelData[i][0] = chanNo
			return nil
		}
	}
	return errors.New("未找到指定信道！")
}

func RemoveAllEmpty(configs *radio.ClassTheRadioData) {
	che := new([128][]string)
	c := 0
	for _, v := range configs.ChannelData {
		for k, l := range v {
			if l != "" && k != 0 {
				che[c] = v
				che[c][0] = strconv.Itoa(c)
				c += 1
				break
			}
		}
	}
	for i := c; i < len(che); i++ {
		che[i] = []string{
			strconv.Itoa(i),
			"", "", "", "", "", "", "", "", "", "", "", "", "",
		}
	}
	configs.ChannelData = *che
}
func BatchClear(configs *radio.ClassTheRadioData, chanClearRange []string) error {
	for _, v := range chanClearRange {
		if err := ClearChannel(configs, v); err != nil {
			return err
		}
	}
	return nil
}

func AddChannel(configs *radio.ClassTheRadioData, pos string, data [14]string) error {
	for i, v := range configs.ChannelData {
		if v[0] == pos {
			if !isChannelEmpty(configs.ChannelData[127]) {
				fmt.Println("信道127不为空！继续操作将移除信道127，是否需要复制信道127？输入y确认")
				var need string
				_, _ = fmt.Scanln(&need)
				if need == "y" {
					if err := copyChannel(*configs, "127"); err != nil {
						slog.Warnf("复制失败")
					} else {
						slog.Info("复制成功")
					}
				}
			}
			cache := configs.ChannelData[:]
			tmp := append([][]string{}, cache[i:]...)
			cache = append(cache[0:i], data[:])
			for g := range tmp {
				cc, _ := strconv.Atoi(tmp[g][0])
				cc += 1
				tmp[g][0] = strconv.Itoa(cc)
			}
			cache = append(cache, tmp...)
			for j := range configs.ChannelData {
				if j < 128 {
					configs.ChannelData[j] = cache[j]
				}
			}
			return nil
		}
	}
	return errors.New("指定信道无效")
}
func AskForDetail() (*[14]string, error) {
	inp := new([14]string)
	for i := 1; i < 14; i++ {
		fmt.Printf("请输入%s,输入ABORT取消增加信道", available_choices[i])
		var choice string
		_, _ = fmt.Scanln(&choice)
		if choice == "ABORT" {
			return nil, errors.New("用户取消")
		}
		inp[i] = choice
	}
	return inp, nil
}
func ModifyChannel(configs *radio.ClassTheRadioData, pos string, modPos int, value string) error {
	for i, v := range configs.ChannelData {
		if v[0] == pos {
			configs.ChannelData[i][modPos] = value
			return nil
		}
	}
	return errors.New("指定信道无效")
}
func displayModify() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	for i := 1; i < len(available_choices); i++ {
		t.AppendRow(table.Row{
			i, available_choices[i],
		})
	}
	t.Render()
}

func parseChannel(channel string) ([]string, error) {
	spliter := strings.Split(channel, ",")
	chanss := make([]string, 0)
	if len(spliter) == 0 {
		return nil, errors.New("长度不能为空")
	}
	for index, v := range spliter {
		if v == "" {
			continue
		}
		if strings.Contains(v, "-") {
			tmp := strings.Split(v, "-")
			if len(tmp) > 2 {
				return nil, errors.New("输入错误: 分隔符有错误")
			}
			for _, v1 := range tmp {
				if _, err := strconv.Atoi(v1); err != nil {
					return nil, errors.New("输入错误：输入需为数字")
				}
			}
			rge1, _ := strconv.Atoi(tmp[0])
			rge2, _ := strconv.Atoi(tmp[1])
			if rge1 > rge2 {
				return nil, errors.New("输入错误：开始大于结束")
			}
			start := rge1
			for start <= rge2 {
				chanss = append(chanss, strconv.Itoa(start))
				start += 1
			}
		} else {
			if _, err := strconv.Atoi(v); err != nil {
				return nil, errors.New("输入错误：输入需为数字")
			}
			chanss = append(chanss, spliter[index])
		}
	}
	return chanss, nil
}
func dCopyByGob(src, dst interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buffer).Decode(dst)
}
func backupConfig(config *radio.ClassTheRadioData, listBack *[]radio.ClassTheRadioData) {
	backup := new(radio.ClassTheRadioData)
	_ = dCopyByGob(config, backup)
	*listBack = append(*listBack, *backup)
}

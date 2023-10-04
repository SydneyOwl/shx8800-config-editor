package entrance

var available_choices = []string{
	"信道", "发射允许", "接收频率", "QT/DTQ解码(接收哑音)", "发射频率", "QT/DQT编码（发射亚音）", "发射功率", "带宽", "PTT-ID", "繁忙锁定", "扫描添加", "信令码", "信道名称", "加密",
}

const (
	SKIP_EMPTY_CHANNEL_COUNT = 3
)

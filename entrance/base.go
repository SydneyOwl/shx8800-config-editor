package entrance

import (
	"fmt"
	"github.com/gookit/slog"
	"github.com/spf13/cobra"
	"github.com/sydneyowl/shx8800-config-editor/config"
	"github.com/sydneyowl/shx8800-config-editor/pkg/filetools"
	"github.com/sydneyowl/shx8800-config-editor/pkg/logger"
)

var (
	Verbose      = false
	Vverbose     = false
	ReleaseDep   = false
	PrintVersion = false
	DepPath      = ""
	ConfigPath   = ""
)

func printVer() {
	fmt.Println("SHX8800 dat editor")
	fmt.Printf("Version: %s\n", config.VER)
	fmt.Printf("Commit: %s\n", config.COMMIT)
	fmt.Printf("Build Time: %s\n", config.BUILDTIME)
}

var BaseCmd = &cobra.Command{
	Use:   "SHX8800_EDITOR",
	Short: "SHX8800_EDITOR",
	//Version: config.VER,
	Long: `SHX8800_EDITOR - SHX8800 dat editor`,
	Run: func(cmd *cobra.Command, args []string) {
		if PrintVersion {
			printVer()
			return
		}
		logger.InitLog(Verbose, Vverbose)
		if ReleaseDep {
			err := filetools.ReleaseFile()
			if err != nil {
				slog.Fatalf("无法释放依赖！")
				_, _ = fmt.Scanln()
				return
			}
			slog.Info("释放成功！")
			return
		}
		mainRunner()
	},
}

func init() {
	BaseCmd.PersistentFlags().BoolVar(&Verbose, "verbose", false, "打印啰嗦的日志")
	BaseCmd.PersistentFlags().BoolVar(&Vverbose, "vverbose", false, "打印超级啰嗦的日志")
	BaseCmd.PersistentFlags().BoolVar(&PrintVersion, "version", false, "打印软件版本信息")
	if filetools.Build == "release" {
		BaseCmd.Flags().BoolVar(&ReleaseDep, "release-dep", false, "释放依赖的C#文件")
	}
	BaseCmd.Flags().StringVar(&DepPath, "dep-path", "", "C#依赖所在的路径")
	BaseCmd.Flags().StringVar(&ConfigPath, "dat-path", "", "DAT文件所在的路径")
}

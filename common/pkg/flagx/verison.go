package flagx

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var (
	builtAt   string
	buildUser string
	builtOn   string
	goVersion string
	gitAuthor string
	gitCommit string
	gitTag    string
)

func init() {
	flag.Bool("v", false, "打印版本信息")
	Register("version", PrintVersionInfo)
}
func PrintVersionInfo(val string) {
	fmt.Printf("%-20s %s\n", "builtAt", color.GreenString(builtAt))
	fmt.Printf("%-20s %s\n", "builtOn", color.GreenString(builtOn))
	fmt.Printf("%-20s %s\n", "buildUser", color.GreenString(buildUser))
	fmt.Printf("%-20s %s\n", "goVersion", color.GreenString(goVersion))
	fmt.Printf("%-20s %s\n", "gitAuthor", color.GreenString(gitAuthor))
	fmt.Printf("%-20s %s\n", "gitCommit", color.GreenString(gitCommit))
	fmt.Printf("%-20s %s\n", "gitTag", color.GreenString(gitTag))
	os.Exit(1)
}

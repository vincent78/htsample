package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"htSample/global"
	"os"
	"runtime"
	"strings"
)

var (
	VersionCommand = &cli.Command{
		Action:    version,
		Name:      "version",
		Usage:     "Print version numbers",
		ArgsUsage: " ",
		Category:  "",
		Description: `
The output of this command is supposed to be machine-readable.
`,
	}
)

func version(_ *cli.Context) error {
	fmt.Println(strings.Title(global.ClientIdentifier))
	fmt.Println("Version:", VersionWithMeta)
	if global.GitCommit != "" {
		fmt.Println("Git Commit:", global.GitCommit)
	}
	if global.BuildDate != "" {
		fmt.Println("Build Date:", global.BuildDate)
	}
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Printf("GOPATH=%s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOROOT=%s\n", runtime.GOROOT())
	return nil
}

var Version = func() string {
	return fmt.Sprintf("%d.%d.%d", global.VersionMajor, global.VersionMinor, global.VersionPatch)
}()

// VersionWithMeta holds the textual version string including the metadata.
var VersionWithMeta = func() string {
	v := Version
	if global.VersionMeta != "" {
		v += "-" + global.VersionMeta
	}
	return v
}()

func VersionWithCommit(gitCommit, buildDate string) string {
	vsn := VersionWithMeta
	if len(gitCommit) >= 8 {
		vsn += "-" + gitCommit[:8]
	} else if len(gitCommit) > 0 {
		vsn += "-" + gitCommit
	}
	if (global.VersionMeta != "stable") && (buildDate != "") {
		vsn += "-" + buildDate
	}
	return vsn
}

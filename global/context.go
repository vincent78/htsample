package global

import (
	"fmt"
	"gorm.io/gorm"
)

var DB *gorm.DB

var GitCommit string
var BuildDate string

const (
	ClientIdentifier = "htsample"

	VersionMajor = 0          // Major version component of the current release
	VersionMinor = 0          // Minor version component of the current release
	VersionPatch = 1          // Patch version component of the current release
	VersionMeta  = "snapshot" // Version metadata to append to the version string   stable   snapshot
)

func EnvName(n string) string {
	return fmt.Sprintf("%v_%v", "HTS", n)
}

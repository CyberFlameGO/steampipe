package parse

import (
	"github.com/Masterminds/semver"
	"github.com/turbot/steampipe/steampipeconfig/modconfig"
)

type InstalledMod struct {
	Mod     *modconfig.Mod
	Version *semver.Version
}

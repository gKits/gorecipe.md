package version

import (
	"fmt"
	"runtime/debug"
)

func Version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}

	ver := "0.0.0-local"
	time := "0000-00-00"
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs", "vcs.modified":
			ver = s.Value
		case "vcs.revision":
			if ver == "" {
				ver = s.Value
			}
		case "vcs.time":
			time = s.Value
		}
	}

	return fmt.Sprintf("v%s build with %s at %s", ver, info.GoVersion, time)
}

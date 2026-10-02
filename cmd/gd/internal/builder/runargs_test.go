package builder

import (
	"slices"
	"strings"
	"testing"
)

func TestSplitRunArgs(t *testing.T) {
	for _, test := range []struct {
		args, build, engine string
	}{
		{args: ""},
		{
			args:   "--quit-after 5",
			engine: "--quit-after 5",
		},
		{
			args:   "-gcflags=graphics.gd/classdb/...=-N -tags debug --write-movie out.avi --fixed-fps 60",
			build:  "-gcflags=graphics.gd/classdb/...=-N -tags debug",
			engine: "--write-movie out.avi --fixed-fps 60",
		},
		{
			// go takes its flags with two dashes as well.
			args:   "--tags=debug --race --headless",
			build:  "--tags=debug --race",
			engine: "--headless",
		},
		{
			// the value of a go flag is never mistaken for an engine one.
			args:   "-ldflags -w -d main.tscn",
			build:  "-ldflags -w",
			engine: "-d main.tscn",
		},
		{
			args:   "-v -x",
			build:  "-v -x",
			engine: "--verbose",
		},
		{
			args:   "--verbose",
			build:  "-v",
			engine: "--verbose",
		},
		{
			// arguments for the project are never taken for go flags.
			args:   "-race --quit -- -tags mine -x",
			build:  "-race",
			engine: "--quit -- -tags mine -x",
		},
		{
			args:   "res://main.tscn ++ -a",
			engine: "res://main.tscn ++ -a",
		},
		{
			args:  "-tags",
			build: "-tags",
		},
	} {
		build, engine := splitRunArgs(strings.Fields(test.args))
		if !slices.Equal(build, strings.Fields(test.build)) {
			t.Errorf("splitRunArgs(%q) build = %q, want %q", test.args, build, test.build)
		}
		if !slices.Equal(engine, strings.Fields(test.engine)) {
			t.Errorf("splitRunArgs(%q) engine = %q, want %q", test.args, engine, test.engine)
		}
	}
}

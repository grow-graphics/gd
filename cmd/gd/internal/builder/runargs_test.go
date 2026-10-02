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
			build:  "-gcflags=graphics.gd/classdb/...=-N -tags=debug",
			engine: "--write-movie out.avi --fixed-fps 60",
		},
		{
			// go takes its flags with two dashes as well.
			args:   "--tags=debug --race --headless",
			build:  "-tags=debug -race",
			engine: "--headless",
		},
		{
			// the value of a go flag is never mistaken for an engine one.
			args:   "-ldflags -w -d main.tscn",
			build:  "-ldflags=-w",
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
		build, engine := SplitRunArgs(strings.Fields(test.args))
		if !slices.Equal(build, strings.Fields(test.build)) {
			t.Errorf("SplitRunArgs(%q) build = %q, want %q", test.args, build, test.build)
		}
		if !slices.Equal(engine, strings.Fields(test.engine)) {
			t.Errorf("SplitRunArgs(%q) engine = %q, want %q", test.args, engine, test.engine)
		}
	}
}

func TestUserTags(t *testing.T) {
	for _, test := range []struct {
		args, rest, tags string
	}{
		{args: "", tags: "musl"},
		{args: "-v -trimpath", rest: "-v -trimpath", tags: "musl"},
		{args: "-tags=debug -v", rest: "-v", tags: "musl,debug"},
		{args: "-v --tags a,b", rest: "-v", tags: "musl,a,b"},
		{args: "-tags=a -x -tags=b", rest: "-x", tags: "musl,b"},
		{args: "-test.run Tags -tags", rest: "-test.run Tags", tags: "musl"},
	} {
		rest, tags := userTags("musl", strings.Fields(test.args))
		if !slices.Equal(rest, strings.Fields(test.rest)) || tags != test.tags {
			t.Errorf("userTags(%q) = %q, %q, want %q, %q", test.args, rest, tags, test.rest, test.tags)
		}
	}
	// one argument holding a space-separated list, the way go still takes it.
	if _, tags := userTags("musl", []string{"-tags=a b"}); tags != "musl,a,b" {
		t.Errorf("userTags of a space-separated list = %q", tags)
	}
}

func TestWithLinkerFlag(t *testing.T) {
	for _, test := range []struct {
		args, want []string
	}{
		{nil, []string{"-ldflags=-extar=ar"}},
		{[]string{"-v"}, []string{"-v", "-ldflags=-extar=ar"}},
		{[]string{"-ldflags=-s -w"}, []string{"-ldflags=-s -w -extar=ar"}},
		// the user's flags come after gd's, and are the ones that go takes.
		{[]string{"-ldflags=-s -w", "-ldflags=-X main.v=1"}, []string{"-ldflags=-s -w", "-ldflags=-X main.v=1 -extar=ar"}},
		{[]string{"-ldflags=-s -w", "--ldflags", "-X main.v=1", "-v"}, []string{"-ldflags=-s -w", "--ldflags", "-X main.v=1 -extar=ar", "-v"}},
	} {
		if got := withLinkerFlag(test.args, "-extar=ar"); !slices.Equal(got, test.want) {
			t.Errorf("withLinkerFlag(%q) = %q, want %q", test.args, got, test.want)
		}
	}
}

package builder

import (
	"fmt"
	"os"
	"strings"
)

// goBuildFlags are the flags that `go build` accepts and whether each of
// them takes a value (see `go help build`).
var goBuildFlags = map[string]bool{
	"C": true, "a": false, "n": false, "p": true, "race": false, "msan": false,
	"asan": false, "cover": false, "covermode": true, "coverpkg": true,
	"work": false, "x": false, "asmflags": true, "buildmode": true,
	"buildvcs": false, "compiler": true, "gccgoflags": true, "gcflags": true,
	"installsuffix": true, "json": false, "ldflags": true, "linkshared": false,
	"mod": true, "modcacherw": false, "modfile": true, "overlay": true,
	"pgo": true, "pkgdir": true, "tags": true, "trimpath": false,
	"toolexec": true, "o": true,
	"debug-actiongraph": true, "debug-runtime-trace": true, "debug-trace": true,
}

// splitRunArgs routes the arguments of `gd run` to the command each of them
// is for: the flags of `go build` (along with their values) configure the
// build, everything else is for the engine that runs the project, so that
//
//	gd run -tags debug --quit-after 100 --write-movie out.avi
//
// builds with the debug tag and hands the rest to godot. The set of go build
// flags is small and fixed, the engine's is neither (it grows with each
// release, and a fork may add its own), so the go flags are the ones
// recognised. The two commands share verbosity, which goes to both, and
// where a short flag means something to each (-p) it is the go one. Anything
// from the engine's own "--" (or "++") separator onwards is for the project
// and reaches it untouched.
func splitRunArgs(args []string) (build, engine []string) {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--" || arg == "++" {
			return build, append(engine, args[i:]...)
		}
		if !strings.HasPrefix(arg, "-") {
			engine = append(engine, arg)
			continue
		}
		// go accepts its flags with either one dash or two.
		name, _, inline := strings.Cut(strings.TrimPrefix(strings.TrimPrefix(arg, "-"), "-"), "=")
		if arg == "-v" || arg == "--verbose" {
			build = append(build, "-v")
			engine = append(engine, "--verbose")
			continue
		}
		value, ok := goBuildFlags[name]
		if !ok {
			engine = append(engine, arg)
			continue
		}
		build = append(build, arg)
		if value && !inline && i+1 < len(args) {
			i++
			build = append(build, args[i])
		}
	}
	return build, engine
}

// splitDeviceRunArgs is [splitRunArgs] for the targets that `gd run` installs
// and launches on a device, where there is no command line to hand the
// engine's arguments to.
func splitDeviceRunArgs(args []string) (build []string) {
	build, engine := splitRunArgs(args)
	if len(engine) > 0 {
		fmt.Fprintln(os.Stderr, "gd: ignoring", strings.Join(engine, " "), "— engine arguments cannot be passed to a project launched on a device")
	}
	return build
}

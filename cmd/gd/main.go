// The 'gd' command is designed as a drop-in replacement of the 'go' command when working
// with Godot-based projects. It will automatically download and install the supported
// version of Godot and ensure all go commands behave as expected (go build, go run, etc).
//
// The 'gd' command assumes that the Go module lives at the root of the project, an empty
// godot project will be created under a 'graphics' directory, this is where the user can
// keep the graphical representation of their project and manage their assets. Running the
// command without any command line arguments will launch the Godot editor for managing
// the assets in this directory.
package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"go/build"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	_ "embed"
)

// These are our initial Godot project template files, we create
// these automatically when the user runs the 'gd' command. They
// are minimally setup for including the Go shared library such
// that it will be executed on startup.
var (
	//go:embed graphics/project.godot
	project_godot string

	//go:embed graphics/library.gdextension
	library_gdextension string

	//go:embed graphics/main.tscn
	main_tscn string

	//go:embed graphics/.godot/extension_list.cfg
	extension_list_cfg string
)

func setupFile(force bool, name, embed string, args ...any) error {
	if _, err := os.Stat(name); force || os.IsNotExist(err) {
		if len(args) > 0 {
			embed = fmt.Sprintf(embed, args...)
		}
		if err := os.WriteFile(name, []byte(embed), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := wrap(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func useGodot() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	godot, err := exec.LookPath("godot")
	if err == nil {
		version, err := exec.Command(godot, "--version").CombinedOutput()
		if err == nil {
			if strings.HasPrefix(string(version), "4.2.2") {
				return godot, nil
			}
		}
	}
	// Use existing godot-4.2.2 if available.
	if binary, err := exec.LookPath("godot-4.2.2"); err == nil {
		return binary, nil
	}
	info, err := os.Stat(gopath + "/bin/godot-4.2.2")
	if os.IsNotExist(err) {
		switch runtime.GOOS {
		case "android":
			return "echo", nil
		case "darwin":
			if _, err := exec.LookPath("brew"); err == nil {
				fmt.Println("gd: downloading Godot v4.2.2 stable for macOS (via brew)")
				if err := exec.Command("brew", "install", "godot").Run(); err != nil {
					return "", err
				}
			}
		case "linux":
			fmt.Println("gd: downloading Godot v4.2.2 stable for linux")
			resp, err := http.Get("https://github.com/godotengine/godot-builds/releases/download/4.2.2-stable/Godot_v4.2.2-stable_linux.x86_64.zip")
			if err != nil {
				return "", err
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				return "", err
			}
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return "", err
			}
			archive, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
			if err != nil {
				return "", err
			}
			inZip, err := archive.Open("Godot_v4.2.2-stable_linux.x86_64")
			if err != nil {
				return "", err
			}
			defer inZip.Close()
			file, err := os.Create(gopath + "/bin/godot-4.2.2")
			if err != nil {
				return "", err
			}
			defer file.Close()
			if _, err = io.Copy(file, inZip); err != nil {
				return "", err
			}
		default:
			return "", err
		}
	} else if err != nil {
		return "", err
	} else {
		if info.Mode()&0o111 == 0 {
			if err := os.Chmod(gopath+"/bin/godot-4.2.2", 0o755); err != nil {
				return "", err
			}
		}
	}
	return gopath + "/bin/godot-4.2.2", nil
}

func wrap() error {
	GOOS, GOARCH := runtime.GOOS, runtime.GOARCH
	if os.Getenv("GOOS") != "" {
		GOOS = os.Getenv("GOOS")
	}
	if os.Getenv("GOARCH") != "" {
		GOARCH = os.Getenv("GOARCH")
	}
	if GOARCH != "amd64" && GOARCH != "arm64" {
		return errors.New("gd requires an amd64, or an arm64 system")
	}
	godot, err := useGodot()
	if err != nil {
		return fmt.Errorf("gd requires Godot to be installed as a binary at $GOPATH/bin/godot-4.2.2: %w", err)
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	// look for a go.mod file
	for wd := wd; true; wd = filepath.Dir(wd) {
		if wd == "/" {
			return fmt.Errorf("gd requires your project to have a go.mod file")
		}
		_, err := os.Stat(wd + "/go.mod")
		if err == nil {
			break
		} else if os.IsNotExist(err) {
			continue
		} else {
			return err
		}
	}
	graphics := "./graphics"
	if GOOS == "android" {
		graphics = "/sdcard/gd/" + filepath.Base(wd)
	}
	setup := func() error {
		if err := os.MkdirAll(graphics+"/.godot", 0o755); err != nil {
			return err
		}
		if err := setupFile(false, graphics+"/main.tscn", main_tscn); err != nil {
			return err
		}
		if err := setupFile(false, graphics+"/project.godot", project_godot, filepath.Base(wd)); err != nil {
			return err
		}
		if err := setupFile(true, graphics+"/library.gdextension", library_gdextension); err != nil {
			return err
		}
		if err := setupFile(false, graphics+"/.godot/extension_list.cfg", extension_list_cfg); err != nil {
			return err
		}
		_, err := os.Stat(graphics + "/.godot")
		if os.IsNotExist(err) {
			godot := exec.Command(godot, "--import", "--headless")
			godot.Dir = graphics
			godot.Stderr = os.Stderr
			godot.Stdout = os.Stdout
			godot.Stdin = os.Stdin
			return godot.Run()
		}
		return nil
	}
	libraryName := fmt.Sprintf("%v_%v", GOOS, GOARCH)
	switch GOOS {
	case "windows":
		libraryName += ".dll"
	case "darwin":
		libraryName += ".dylib"
	default:
		libraryName += ".so"
	}
	if len(os.Args) == 1 {
		golang := exec.Command("go", "build", "-buildmode=c-shared", "-o", graphics+"/"+libraryName)
		golang.Env = append(os.Environ(), "CGO_ENABLED=1")
		golang.Stderr = os.Stderr
		golang.Stdout = os.Stdout
		golang.Stdin = os.Stdin
		if err := golang.Run(); err != nil {
			return err
		}
		if err := setup(); err != nil {
			return err
		}
		godot := exec.Command(godot, "-e")
		godot.Dir = graphics
		godot.Stderr = os.Stderr
		godot.Stdout = os.Stdout
		godot.Stdin = os.Stdin
		return godot.Run()
	}
	args := make([]string, len(os.Args)-1)

	commands := [][]string{}
	switch os.Args[1] {
	case "run", "build":
		copy(args, os.Args[1:])
		args[0] = "build"

		args = append(args, "-buildmode=c-shared", "-o", graphics+"/"+libraryName)

	case "test":
		args = []string{"test", "-buildmode=c-shared", "-c", "-o", graphics + "/" + libraryName}
	default:
		copy(args, os.Args[1:])
	}

	commands = append(commands, args)
	if runtime.GOOS == "darwin" {
		// GOARCH possible values = "amd64", "arm64"
		missingArch := "arm64"
		if GOARCH == "arm64" {
			missingArch = "amd64"
		}
		missingArgs := make([]string, len(os.Args)-1)

		missingLibraryName := fmt.Sprintf("%v_%v", GOOS, missingArch)
		missingArgs = append(args, "-buildmode=c-shared", "-o", graphics+"/"+missingLibraryName)

		commands = append(commands, missingArgs)

	}

	for _, commandArgs := range commands {
		buildCommand := exec.Command("go", commandArgs...)
		buildCommand.Env = append(os.Environ(), "CGO_ENABLED=1")
		buildCommand.Stderr = os.Stderr
		buildCommand.Stdout = os.Stdout
		buildCommand.Stdin = os.Stdin

		if err := buildCommand.Run(); err != nil {
			return err
		}
	}

	if runtime.GOOS == "darwin" {
		// check if command is available in the system
		_, err := exec.LookPath("lipo")
		if err != nil {
			return err
		}

		lipoCommand := exec.Command("lipo", "-create", graphics+"/darwin_amd64.dylib", graphics+"/darwing_arm64.dylib", "-output", graphics+"/universal.dylib")

		lipoCommand.Stderr = os.Stderr
		lipoCommand.Stdout = os.Stdout
		lipoCommand.Stdin = os.Stdin
		if err := lipoCommand.Run(); err != nil {
			return err
		}
	}

	if err := setup(); err != nil {
		return err
	}
	switch os.Args[1] {
	case "run":
		godot := exec.Command(godot)
		godot.Dir = graphics
		godot.Stderr = os.Stderr
		godot.Stdout = os.Stdout
		godot.Stdin = os.Stdin
		return godot.Run()
	case "test":
		args := []string{"--headless"}
		for _, arg := range os.Args[2:] {
			switch arg {
			case "-bench", "-benchmem", "-benchtime", "blockprofile",
				"-blockprofilerate", "-count", "-coverprofile", "-cpu",
				"-cpuprofile", "-failfast", "-fullpath", "-fuzz", "-fuzzcachedir",
				"-fuzzminimizetime", "-fuzztime", "-fuzzworker", "-gocoverdir",
				"-list", "-memprofile", "-memprofilerate", "-mutexprofile",
				"-mutexprofilefraction", "-outputdir", "-paniconexit0",
				"-parallel", "-run", "-short", "-shuffle", "-skip", "-testlogfile",
				"-timeout", "-trace", "-v":
				args = append(args, "-test."+strings.TrimPrefix(arg, "-"))
			default:
				args = append(args, arg)
			}
		}
		godot := exec.Command(godot, args...)
		godot.Dir = graphics
		godot.Stderr = os.Stderr
		godot.Stdout = os.Stdout
		godot.Stdin = os.Stdin
		return godot.Run()
	}
	return nil
}

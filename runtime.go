package runtime

import (
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ProjectDirectory is the develop project directory
var ProjectDirectory = ""

func init() {
	if _, err := os.Stat(build.Default.GOPATH); err != nil {
		return
	}
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	moduleName := strings.TrimSpace(string(output))
	ProjectDirectory = "/src/" + moduleName
}

// GetRuntimeDirectory will return the runtime directory for
// file access whether during development or release
//
//	param subDir ex. ".config", "http"
func GetRuntimeDirectory(subDir string) (path string) {
	subDir = "/" + subDir + "/"
	executableDirectory, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	path = filepath.Clean(executableDirectory+subDir) + string(filepath.Separator)
	if _, err := os.Stat(path); err == nil {
		return
	}
	developDirectory := build.Default.GOPATH + ProjectDirectory
	path = filepath.Clean(developDirectory+subDir) + string(filepath.Separator)
	if _, err := os.Stat(path); err == nil {
		return
	}
	path = ""
	return
}

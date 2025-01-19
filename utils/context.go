package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/civet148/golazy/version"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// NL defines a new line.
const (
	NL              = "\n"
	golazyDir       = ".golazy"
	gitDir          = ".git"
	autoCompleteDir = ".auto_complete"
	cacheDir        = "cache"
)

var errModuleCheck = errors.New("the work directory must be found in the go mod or the $GOPATH")

// ProjectContext is a structure for the project,
// which contains WorkDir, Name, Path and Dir
type ProjectContext struct {
	WorkDir string
	// Name is the root name of the project
	// eg: go-zero、greet
	Name string
	// Path identifies which module a project belongs to, which is module value if it's a go mod project,
	// or else it is the root name of the project, eg: github.com/zeromicro/go-zero、greet
	Path string
	// Dir is the path of the project, eg: /Users/keson/goland/go/go-zero、/Users/keson/go/src/greet
	Dir string
}

// Prepare checks the project which module belongs to,and returns the path and module.
// workDir parameter is the directory of the source of generating code,
// where can be found the project path and the project module,
func Prepare(workDir string) (*ProjectContext, error) {
	//ctx, err := background(workDir)
	//if err == nil {
	//	return ctx, nil
	//}

	name := filepath.Base(workDir)
	Run("go mod init "+name, workDir)
	return background(workDir)
}

func background(workDir string) (*ProjectContext, error) {
	isGoMod, err := IsGoMod(workDir)
	if err != nil {
		return nil, err
	}

	if isGoMod {
		return projectFromGoMod(workDir)
	}
	return projectFromGoPath(workDir)
}

// RunFunc defines a function type of Run function
type RunFunc func(string, string, ...*bytes.Buffer) (string, error)

// Run provides the execution of shell scripts in golang,
// which can support macOS, Windows, and Linux operating systems.
// Other operating systems are currently not supported
func Run(arg, dir string, in ...*bytes.Buffer) (string, error) {
	goos := runtime.GOOS
	var cmd *exec.Cmd
	switch goos {
	case version.OsMac, version.OsLinux:
		cmd = exec.Command("sh", "-c", arg)
	case version.OsWindows:
		cmd = exec.Command("cmd.exe", "/c", arg)
	default:
		return "", fmt.Errorf("unexpected os: %v", goos)
	}
	if len(dir) > 0 {
		cmd.Dir = dir
	}
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	if len(in) > 0 {
		cmd.Stdin = in[0]
	}
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	if err != nil {
		if stderr.Len() > 0 {
			return "", errors.New(strings.TrimSuffix(stderr.String(), NL))
		}
		return "", err
	}

	return strings.TrimSuffix(stdout.String(), NL), nil
}

// IsGoMod is used to determine whether workDir is a go module project through command `go list -json -m`
func IsGoMod(workDir string) (bool, error) {
	if len(workDir) == 0 {
		return false, errors.New("the work directory is not found")
	}
	if _, err := os.Stat(workDir); err != nil {
		return false, err
	}

	data, err := Run("go list -m -f '{{.GoMod}}'", workDir)
	if err != nil || len(data) == 0 {
		return false, nil
	}

	return true, nil
}

const goModuleWithoutGoFiles = "command-line-arguments"

var errInvalidGoMod = errors.New("invalid go module")

// Module contains the relative data of go module,
// which is the result of the command go list
type Module struct {
	Path      string
	Main      bool
	Dir       string
	GoMod     string
	GoVersion string
}

func (m *Module) validate() error {
	if m.Path == goModuleWithoutGoFiles || m.Dir == "" {
		return errInvalidGoMod
	}
	return nil
}

// projectFromGoMod is used to find the go module and project file path
// the workDir flag specifies which folder we need to detect based on
// only valid for go mod project
func projectFromGoMod(workDir string) (*ProjectContext, error) {
	if len(workDir) == 0 {
		return nil, errors.New("the work directory is not found")
	}
	if _, err := os.Stat(workDir); err != nil {
		return nil, err
	}

	workDir, err := ReadLink(workDir)
	if err != nil {
		return nil, err
	}

	m, err := getRealModule(workDir, Run)
	if err != nil {
		return nil, err
	}
	if err := m.validate(); err != nil {
		return nil, err
	}

	var ret ProjectContext
	ret.WorkDir = workDir
	ret.Name = filepath.Base(m.Dir)
	dir, err := ReadLink(m.Dir)
	if err != nil {
		return nil, err
	}

	ret.Dir = dir
	ret.Path = m.Path
	return &ret, nil
}

func getRealModule(workDir string, execRun RunFunc) (*Module, error) {
	data, err := execRun("go list -json -m", workDir)
	if err != nil {
		return nil, err
	}

	modules, err := decodePackages(strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	for _, m := range modules {
		realDir, err := ReadLink(m.Dir)
		if err != nil {
			return nil, fmt.Errorf("failed to read go.mod, dir: %s, error: %w", m.Dir, err)
		}

		if strings.HasPrefix(workDir, realDir) {
			return &m, nil
		}
	}

	return nil, errors.New("no matched module")
}

func decodePackages(reader io.Reader) ([]Module, error) {
	br := bufio.NewReader(reader)
	if _, err := br.ReadSlice('{'); err != nil {
		return nil, err
	}

	if err := br.UnreadByte(); err != nil {
		return nil, err
	}

	var modules []Module
	decoder := json.NewDecoder(br)
	for decoder.More() {
		var m Module
		if err := decoder.Decode(&m); err != nil {
			return nil, fmt.Errorf("invalid module: %v", err)
		}

		modules = append(modules, m)
	}

	return modules, nil
}


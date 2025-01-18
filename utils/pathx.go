package utils

import (
	"fmt"
	"github.com/civet148/log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	pkgSep           = "/"
	goModeIdentifier = "go.mod"
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// FileExists returns true if the specified file is exists.
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// MaybeCreateFile creates file if not exists
func MaybeCreateFile(dir, subdir, file string) (fp *os.File, created bool, err error) {

	Must(MkdirIfNotExist(path.Join(dir, subdir)))

	fpath := path.Join(dir, subdir, file)
	if FileExists(fpath) {
		return nil, false, nil
	}
	log.Printf("create file %s successful", fpath)
	fp, err = CreateIfNotExist(fpath)
	created = err == nil
	return
}

// MkdirIfNotExist makes directories if the input path is not exists
func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}

// CreateIfNotExist creates a file if it is not exists.
func CreateIfNotExist(file string) (*os.File, error) {
	_, err := os.Stat(file)
	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("%s already exist", file)
	}

	return os.Create(file)
}

// JoinPackages calls strings.Join and returns
func JoinPackages(pkgs ...string) string {
	return strings.Join(pkgs, pkgSep)
}

func GetParentPackage(dir string) (string, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	projectCtx, err := Prepare(abs)
	if err != nil {
		return "", err
	}

	// fix https://github.com/zeromicro/go-zero/issues/1058
	wd := projectCtx.WorkDir
	d := projectCtx.Dir
	same, err := SameFile(wd, d)
	if err != nil {
		return "", err
	}

	trim := strings.TrimPrefix(projectCtx.WorkDir, projectCtx.Dir)
	if same {
		trim = strings.TrimPrefix(strings.ToLower(projectCtx.WorkDir), strings.ToLower(projectCtx.Dir))
	}

	return filepath.ToSlash(filepath.Join(projectCtx.Path, trim)), nil
}

func ReadLink(name string) (string, error) {
	return name, nil
}

// SameFile compares the between path if the same path,
// it maybe the same path in case case-ignore, such as:
// /Users/go_zero and /Users/Go_zero, as far as we know,
// this case maybe appear on macOS and Windows.
func SameFile(path1, path2 string) (bool, error) {
	stat1, err := os.Stat(path1)
	if err != nil {
		return false, err
	}

	stat2, err := os.Stat(path2)
	if err != nil {
		return false, err
	}

	return os.SameFile(stat1, stat2), nil
}


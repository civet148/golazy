package utils

import (
	"bytes"
	"errors"
	"fmt"
	"go/build"
	goformat "go/format"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type style int

const (
	flagGo   = "GO"
	flagLazy = "LAZY"

	unknown style = iota
	title
	lower
	upper
)

type styleFormat struct {
	before    string
	through   string
	after     string
	goStyle   style
	zeroStyle style
}

// ErrNamingFormat defines an error for unknown format
var ErrNamingFormat = errors.New("unsupported format")

func FormatCode(code string) string {
	ret, err := goformat.Source([]byte(code))
	if err != nil {
		return code
	}

	return string(ret)
}

// FileNamingFormat is used to format the file name. You can define the format style
// through the go and Lazy formatting characters. For example, you can define the snake
// format as go_lazy, and the camel case format as goZero. You can even specify the split
// character, such as go#Lazy, theoretically any combination can be used, but the prerequisite
// must meet the naming conventions of each operating system file name.
// Note: Formatting is based on snake or camel string
func FileNamingFormat(format, content string) (string, error) {
	upperFormat := strings.ToUpper(format)
	indexGo := strings.Index(upperFormat, flagGo)
	indexZero := strings.Index(upperFormat, flagLazy)
	if indexGo < 0 || indexZero < 0 || indexGo > indexZero {
		return "", ErrNamingFormat
	}
	var (
		before, through, after string
		flagGo, flagZero       string
		goStyle, zeroStyle     style
		err                    error
	)
	before = format[:indexGo]
	flagGo = format[indexGo : indexGo+2]
	through = format[indexGo+2 : indexZero]
	flagZero = format[indexZero : indexZero+4]
	after = format[indexZero+4:]

	goStyle, err = getStyle(flagGo)
	if err != nil {
		return "", err
	}

	zeroStyle, err = getStyle(flagZero)
	if err != nil {
		return "", err
	}
	var formatStyle styleFormat
	formatStyle.goStyle = goStyle
	formatStyle.zeroStyle = zeroStyle
	formatStyle.before = before
	formatStyle.through = through
	formatStyle.after = after
	return doFormat(formatStyle, content)
}

func doFormat(f styleFormat, content string) (string, error) {
	splits, err := split(content)
	if err != nil {
		return "", err
	}
	var join []string
	for index, split := range splits {
		if index == 0 {
			join = append(join, transferTo(split, f.goStyle))
			continue
		}
		join = append(join, transferTo(split, f.zeroStyle))
	}
	joined := strings.Join(join, f.through)
	return f.before + joined + f.after, nil
}

func getStyle(flag string) (style, error) {
	compare := strings.ToLower(flag)
	switch flag {
	case strings.ToLower(compare):
		return lower, nil
	case strings.ToUpper(compare):
		return upper, nil
	case strings.Title(compare):
		return title, nil
	default:
		return unknown, fmt.Errorf("unexpected format: %s", flag)
	}
}

func split(content string) ([]string, error) {
	var (
		list   []string
		reader = strings.NewReader(content)
		buffer = bytes.NewBuffer(nil)
	)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				if buffer.Len() > 0 {
					list = append(list, buffer.String())
				}
				return list, nil
			}
			return nil, err
		}
		if r == '_' {
			if buffer.Len() > 0 {
				list = append(list, buffer.String())
			}
			buffer.Reset()
			continue
		}

		if r >= 'A' && r <= 'Z' {
			if buffer.Len() > 0 {
				list = append(list, buffer.String())
			}
			buffer.Reset()
		}
		buffer.WriteRune(r)
	}
}

func transferTo(in string, style style) string {
	switch style {
	case upper:
		return strings.ToUpper(in)
	case lower:
		return strings.ToLower(in)
	case title:
		return strings.Title(in)
	default:
		return in
	}
}

// projectFromGoPath is used to find the main module and project file path
// the workDir flag specifies which folder we need to detect based on
// only valid for go mod project
func projectFromGoPath(workDir string) (*ProjectContext, error) {
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

	buildContext := build.Default
	goPath := buildContext.GOPATH
	goPath, err = ReadLink(goPath)
	if err != nil {
		return nil, err
	}

	goSrc := filepath.Join(goPath, "src")
	if !FileExists(goSrc) {
		return nil, errModuleCheck
	}

	wd, err := filepath.Abs(workDir)
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(wd, goSrc) {
		return nil, errModuleCheck
	}

	projectName := strings.TrimPrefix(wd, goSrc+string(filepath.Separator))
	return &ProjectContext{
		WorkDir: workDir,
		Name:    projectName,
		Path:    projectName,
		Dir:     filepath.Join(goSrc, projectName),
	}, nil
}


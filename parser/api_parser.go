package parser

import (
	"bufio"
	"fmt"
	"github.com/civet148/log"
	"os"
	"regexp"
	"strings"
)

const (
	docCompRegex     = `@doc\s+"([^"]+)"\s+@handler\s+(\S+)\s+(\S+)\s+(\S+)\s+\(([^)]+)\)\s+returns\s+\(([^)]+)\)`
	handlerCompRegex = `@handler\s+(\S+)\s+(\S+)\s+(\S+)\s+\(([^)]+)\)\s+returns\s+\(([^)]+)\)`
)

// ApiServer 结构体表示 @server 块的内容
type ApiServer struct {
	Prefix     string `json:"prefix"`
	Group      string `json:"group"`
	Middleware string `json:"middleware"`
}

// ApiSpec 结构体表示每个 ApiSpec 接口的定义
type ApiSpec struct {
	Doc      string `json:"doc"`
	Handler  string `json:"handler"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Request  string `json:"request"`
	Response string `json:"response"`
}

// ApiService 结构体表示一个完整的 service api 块
type ApiService struct {
	Server *ApiServer `json:"server"`
	APIs   []*ApiSpec `json:"apis"`
}

// 从文件加载并解析 ApiSpec 内容
func ParseApiFile(filename string) ([]*ApiService, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var services []*ApiService
	scanner := bufio.NewScanner(file)
	var currentService *ApiService

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // 忽略空行
		}

		// 检查 @server 块
		if strings.HasPrefix(line, "@server") {
			if currentService != nil {
				return nil, fmt.Errorf("invalid syntax: @server block must be followed by service api")
			}
			server, err := parseServerBlock(scanner)
			if err != nil {
				return nil, err
			}
			currentService = &ApiService{Server: &server}
		}

		// 检查 service api 块
		if strings.HasPrefix(line, "service") {
			if currentService == nil {
				return nil, fmt.Errorf("invalid syntax: service api must be preceded by @server block")
			}
			// 继续读取 ApiSpec 接口
			for scanner.Scan() {
				apiLine := strings.TrimSpace(scanner.Text())
				if apiLine == "}" {
					break // 结束 service api 块
				}
				if strings.HasPrefix(apiLine, "@doc") || strings.HasPrefix(apiLine, "@handler") {
					api, err := parseAPIBlock(apiLine, scanner)
					if err != nil {
						return nil, err
					}
					currentService.APIs = append(currentService.APIs, &api)
				}
			}
			services = append(services, currentService)
			currentService = nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return services, nil
}

// 解析 @server 块
func parseServerBlock(scanner *bufio.Scanner) (ApiServer, error) {
	server := ApiServer{}
	var serverContent strings.Builder

	// 读取 @server 块的所有内容
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		serverContent.WriteString(line + " ")
		if strings.Contains(line, ")") {
			break // 结束 @server 块
		}
	}

	log.Infof("%s", serverContent.String())
	// 提取 prefix和middleware
	//	re := regexp.MustCompile(`prefix:\s*(\S+).*group:\s*(\S+)(?:.*middleware:\s*(\S+))?`)
	re := regexp.MustCompile(`(?s)prefix:\s*(\S+)(?:.*?middleware:\s*(\S+))?`)
	matches := re.FindStringSubmatch(serverContent.String())
	log.Infof("regex matches: %+v", matches)
	if len(matches) < 1 {
		return server, log.Errorf("invalid @server format: %s", serverContent.String())
	}
	prefix := matches[1]
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	server.Prefix = matches[1]
	if len(matches) > 2 {
		server.Middleware = matches[2]
	}
	return server, nil
}

// 解析 ApiSpec 接口
func parseAPIBlock(apiLine string, scanner *bufio.Scanner) (ApiSpec, error) {
	var haveDoc bool
	api := ApiSpec{}
	var apiContent strings.Builder

	haveDoc = strings.Contains(apiLine, "@doc")

	apiContent.WriteString(apiLine + " ")
	// 读取 ApiSpec 接口的所有内容
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		apiContent.WriteString(line + " ")
		if strings.Contains(line, "returns") {
			break // 结束 ApiSpec 接口
		}
	}
	//log.Infof("api content [%s]", apiContent.String())

	// 提取 doc、handler、method、path、request 和 response
	var re *regexp.Regexp
	if haveDoc {
		re = regexp.MustCompile(docCompRegex)
		matches := re.FindStringSubmatch(apiContent.String())
		if len(matches) != 7 {
			return api, fmt.Errorf("invalid ApiSpec format: %s", apiContent.String())
		}
		api.Doc = matches[1]
		api.Handler = matches[2]
		api.Method = strings.ToLower(matches[3])
		api.Path = matches[4]
		api.Request = matches[5]
		api.Response = matches[6]
	} else {
		re = regexp.MustCompile(handlerCompRegex)
		matches := re.FindStringSubmatch(apiContent.String())
		if len(matches) != 6 {
			return api, fmt.Errorf("invalid ApiSpec format: %s", apiContent.String())
		}
		api.Handler = matches[1]
		api.Method = strings.ToLower(matches[2])
		api.Path = matches[3]
		api.Request = matches[4]
		api.Response = matches[5]
	}
	return api, nil
}

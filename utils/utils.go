package utils

import (
	"github.com/civet148/golazy/parser"
	"strings"
)

func GetMiddleware(as *parser.ApiServer) []string {
	var results []string
	middleware := as.Middleware
	if len(middleware) > 0 {
		for _, item := range strings.Split(middleware, ",") {
			if item != "" {
				results = append(results, strings.TrimSpace(item))
			}
		}
	}
	return results
}

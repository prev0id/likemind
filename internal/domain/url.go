package domain

import (
	"regexp"
	"strings"
)

const (
	PathPing      = "/ping"
	PathUserPage  = "/user/{username}"
	PathGroupPage = "/group/{group}"
	PathSignIn    = "/signin"
	PathSignUp    = "/signup"

	PathPrefixRoot   = "/"
	PathPrefixAPI    = "/api"
	PathPrefixStatic = "/static"

	PatternFile       = "/{file}"
	PathPatternSearch = "/search/{type}"

	PathVarUsername = "username"
	PathVarGroup    = "group"
	PathVarType     = "type"

	PathStaticHTMX         = "/static/htmx.js"
	PathStaticErrorHandler = "/static/error_handler.js"
	PathStaticFavicon      = "/static/favicon.png"
	PathStaticStyles       = "/static/styles.css"
)

var varRegex = regexp.MustCompile(`{\w+}`)

func SetPathVariables(path string, values map[string]string) string {
	return varRegex.ReplaceAllStringFunc(path, func(pathVar string) string {
		key := strings.Trim(pathVar, "{}")
		val, ok := values[key]
		if !ok {
			return pathVar
		}
		return val
	})
}

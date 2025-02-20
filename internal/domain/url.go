package domain

const (
	PathPing       = "/ping"
	PathUserPage   = "/user/{nickname}"
	PathGroupPage  = "/group/{group}"
	PathSignIn     = "/signin"
	PathSignUp     = "/signup"
	PathLogOut     = "/logout"
	PathSearch     = "/search/{type}"
	PathSearchUser = "/search/user"

	PathStaticHTMX         = "/static/htmx.js"
	PathStaticAlpine       = "/static/alpine.js"
	PathStaticErrorHandler = "/static/error_handler.js"
	PathStaticFavicon      = "/static/favicon.png"
	PathStaticStyles       = "/static/styles.css"

	PathAPISignIn = PathPrefixAPI + PathSignIn
	PathAPISignUp = PathPrefixAPI + PathSignUp

	PathPrefixRoot   = "/"
	PathPrefixAPI    = "/api"
	PathPrefixStatic = "/static"

	PatternFile = "/{file}"

	PathVarNickname = "nickname"
	PathVarGroup    = "group"
	PathVarType     = "type"

	TypeGroup = "group"
	TypeUser  = "user"
)

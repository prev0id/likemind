package domain

const (
	PathPageRoot = "/"

	PathPageSignIn      = "/signin"
	PathPageSignUp      = "/signup"
	PathPageCurrentUser = "/user"
	PathPageUser        = "/user/{username}"
	PathPageGroup       = "/group/{groupname}"
	PathPageSearchUser  = "/search/user"
	PathPageSearchGroup = "/search/group"

	PathAPIPing        = "/api/ping"
	PathAPISearchUser  = "/api/search/user"
	PathAPISearchGroup = "/api/search/group"

	PathAPIUserSignup    = "/api/user"
	PathAPIUserSignin    = "/api/user/sigin"
	PathAPIUserLogout    = "/api/user/logout"
	PathAPIUserDesc      = "/api/user/desc"
	PathAPIUserEmail     = "/api/user/email"
	PathAPIUserPassword  = "/api/user/password"
	PathAPIUserPicture   = "/api/user/picture"
	PathAPIUserInterests = "/api/user/interests"

	PathAPIGroup          = "/api/group"
	PathAPIGroupDesc      = "/api/group/desc"
	PathAPIGroupInterests = "/api/group/interests"

	PathStatic             = "/static/{filepath}"
	PathStaticHTMX         = "/static/htmx.js"
	PathStaticAlpine       = "/static/alpine.js"
	PathStaticErrorHandler = "/static/error_handler.js"
	PathStaticFavicon      = "/static/favicon.png"
	PathStaticStyles       = "/static/styles.css"

	PathParamUsername  = "username"
	PathParamGroupname = "groupname"
	PathParamFilePath  = "filepath"
)

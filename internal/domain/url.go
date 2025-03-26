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

	PathAPIPing        = prefixAPI + "/ping"
	PathAPISearchUser  = prefixAPI + "/search/user"
	PathAPISearchGroup = prefixAPI + "/search/group"

	PathAPIUserSignup    = prefixAPI + "/user"
	PathAPIUserSignin    = prefixAPI + "/user/sigin"
	PathAPIUserLogout    = prefixAPI + "/user/logout"
	PathAPIUserDesc      = prefixAPI + "/user/desc"
	PathAPIUserEmail     = prefixAPI + "/user/email"
	PathAPIUserPassword  = prefixAPI + "/user/password"
	PathAPIUserPicture   = prefixAPI + "/user/picture"
	PathAPIUserInterests = prefixAPI + "/user/interests"

	PathAPIGroup          = prefixAPI + "/group"
	PathAPIGroupDesc      = prefixAPI + "/group/desc"
	PathAPIGroupInterests = prefixAPI + "/group/interests"

	PathStaticHTMX         = prefixStatic + "/js/htmx.min.js"
	PathStaticAlpine       = prefixStatic + "/js/alpine.js"
	PathStaticErrorHandler = prefixStatic + "/js/error_handler.js"
	PathStaticFavicon      = prefixStatic + "/img/favicon.png"
	PathStaticStyles       = prefixStatic + "/css/styles.css"

	PathParamUsername  = "username"
	PathParamGroupname = "groupname"
	PathParamFilePath  = "filepath"

	prefixAPI    = "/v1/api"
	prefixStatic = "/static"
)

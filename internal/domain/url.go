package domain

const (
	PathPageRoot = "/"

	PathPageSignIn      = "/signin"
	PathPageSignUp      = "/signup"
	PathPageCurrentUser = "/user"
	PathPageUser        = "/user/{username}"
	PathPageGroup       = "/group/{groupname}"
	PathPageSearch      = "/search"

	PathAPIPing        = prefixAPI + "/ping"
	PathAPISearchUser  = prefixAPI + "/search/user"
	PathAPISearchGroup = prefixAPI + "/search/group"

	PathAPISignup = prefixAPI + "/signup"
	PathAPISignin = prefixAPI + "/signin"
	PathAPILogout = prefixAPI + "/logout"

	PathAPIProfile    = prefixAPI + "/profile"
	PathAPIContact    = prefixAPI + "/contact"
	PathAPIContactID  = prefixAPI + "/contact/{contact_id}"
	PathAPIInterestID = prefixAPI + "/interest/{interest_id}"

	PathAPIGroup           = prefixAPI + "/group"
	PathAPIGroupInterestID = prefixAPI + "/group/interest/{interest_id}"

	PathStaticHTMX         = prefixStatic + "/js/htmx.min.js"
	PathStaticAlpine       = prefixStatic + "/js/alpine.js"
	PathStaticErrorHandler = prefixStatic + "/js/error_handler.js"
	PathStaticFavicon      = prefixStatic + "/img/favicon.png"
	PathStaticStyles       = prefixStatic + "/css/styles.css"

	PathParamUsername   = "username"
	PathParamGroupname  = "groupname"
	PathParamFilePath   = "filepath"
	PathParamContactID  = "contact_id"
	PathParamInterestID = "interest_id"

	prefixAPI     = "/v1/api"
	prefixStatic  = "/static"
	prefixProfiel = "/static"
)

package domain

const (
	PathPageRoot = "/"

	PathPageSignIn      = "/signin"
	PathPageSignUp      = "/signup"
	PathPageCurrentUser = "/user"
	PathPageProfile     = "/profile/{username}"
	PathPageGroup       = "/group/{groupname}"
	PathPageSearch      = "/search"

	PathAPIPing        = prefixAPI + "/ping"
	PathAPISearchUser  = prefixAPI + "/search/user"
	PathAPISearchGroup = prefixAPI + "/search/group"

	PathAPISignin = prefixAPI + "/signin"
	PathAPILogout = prefixAPI + "/logout"

	PathAPIProfile    = prefixAPI + "/profile"
	PathAPIContact    = prefixAPI + "/contact"
	PathAPIContactID  = prefixAPI + "/contact/{contact_id}"
	PathAPIInterestID = prefixAPI + "/interest/{interest_id}"

	PathAPIGroup     = prefixAPI + "/group"
	PathAPIGroupID   = prefixAPI + "/group/{group_id}"
	PathAPIPost      = prefixAPI + "/group/{group_id}/post"
	PathAPIPostID    = prefixAPI + "/group/{group_id}/post/{post_id}"
	PathAPIComment   = prefixAPI + "/group/{group_id}/post/{post_id}/comment"
	PathAPICommentID = prefixAPI + "/group/{group_id}/post/{post_id}/comment/{comment_id}"

	PathStaticHTMX         = prefixStatic + "/js/htmx.min.js"
	PathStaticAlpine       = prefixStatic + "/js/alpine.js"
	PathStaticErrorHandler = prefixStatic + "/js/error_handler.js"
	PathStaticFavicon      = prefixStatic + "/img/favicon.png"
	PathStaticStyles       = prefixStatic + "/css/styles.css"

	PathParamUsername   = "username"
	PathParamGroupID    = "group_id"
	PathParamPostID     = "post_id"
	PathParamCommentID  = "comment_id"
	PathParamFilePath   = "filepath"
	PathParamContactID  = "contact_id"
	PathParamInterestID = "interest_id"

	prefixAPI     = "/v1/api"
	prefixStatic  = "/static"
	prefixProfiel = "/static"
)

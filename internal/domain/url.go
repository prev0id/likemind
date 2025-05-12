package domain

const (
	PathPageRoot = "/"

	PathPageSignIn      = "/signin"
	PathPageSignUp      = "/signup"
	PathPageCurrentUser = "/user"
	PathPageOwnProfile  = "/profile"
	PathPageProfile     = "/profile/{username}"
	PathPageOwnGroups   = "/group"
	PathPageGroup       = "/group/{groupname}"
	PathPageSearch      = "/search"

	PathAPIPing        = prefixAPI + "/ping"
	PathAPISearchUser  = prefixAPI + "/search/user"
	PathAPISearchGroup = prefixAPI + "/search/group"

	PathAPISignin = prefixAPI + "/signin"
	PathAPILogout = prefixAPI + "/logout"

	PathAPIProfile           = prefixAPI + "/profile"
	PathAPIProfileContact    = prefixAPI + "/profile/contact"
	PathAPIProfileContactID  = prefixAPI + "/profile/contact/{contact_id}"
	PathAPIProfileInterestID = prefixAPI + "/profile/interest/{interest_id}"
	PathAPIImage             = prefixAPI + "/profile/image"
	PathAPIImageID           = prefixAPI + "/profile/image/{image_id}"
	PathAPIPassword          = prefixAPI + "/profile/password"
	PathAPIEmail             = prefixAPI + "/profile/email"

	PathAPIGroup           = prefixAPI + "/group"
	PathAPIGroupID         = prefixAPI + "/group/{group_id}"
	PathAPGroupIPost       = prefixAPI + "/group/{group_id}/post"
	PathAPGroupIPostID     = prefixAPI + "/group/{group_id}/post/{post_id}"
	PathAPGroupIComment    = prefixAPI + "/group/{group_id}/post/{post_id}/comment"
	PathAPGroupICommentID  = prefixAPI + "/group/{group_id}/post/{post_id}/comment/{comment_id}"
	PathAPIGroupInterestID = prefixAPI + "/group/interest/{interest_id}"

	PathStaticHTMX         = prefixStatic + "/js/htmx.min.js"
	PathStaticErrorHandler = prefixStatic + "/js/error_handler.js"
	PathStaticUploadFile   = prefixStatic + "/js/upload_file.js"
	PathStaticModal        = prefixStatic + "/js/modal.js"
	PathStaticFavicon      = prefixStatic + "/img/favicon.png"
	PathStaticStyles       = prefixStatic + "/css/styles.css"

	PathParamUsername   = "username"
	PathParamGroupID    = "group_id"
	PathParamPostID     = "post_id"
	PathParamCommentID  = "comment_id"
	PathParamFilePath   = "filepath"
	PathParamContactID  = "contact_id"
	PathParamInterestID = "interest_id"
	PathParamImageID    = "image_id"

	prefixAPI     = "/v1/api"
	prefixStatic  = "/static"
	prefixProfiel = "/static"
)

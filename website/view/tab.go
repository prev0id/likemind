package view

import "likemind/internal/domain"

type Tab struct {
	URL        string
	Name       string
	Kaomoji    string
	Authorized bool
}

var (
	WithoutSelectedTab = Tab{}
	GroupTab           = Tab{
		URL:        domain.PathPageOwnGroups,
		Name:       "Group",
		Kaomoji:    `\( ˙▿˙ )/\( ˙▿˙ )/`,
		Authorized: true,
	}
	SearchTab = Tab{
		URL:        domain.PathPageSearch,
		Name:       "Search",
		Kaomoji:    `⊂(￣▽￣)⊃`,
		Authorized: true,
	}
	ProfileTab = Tab{
		URL:        domain.PathPageOwnProfile,
		Name:       "My Profile",
		Kaomoji:    `(„• ᴗ •„)`,
		Authorized: true,
	}
	SignInTab = Tab{
		URL:     domain.PathPageSignIn,
		Name:    "Signin",
		Kaomoji: `(„• ֊ •„)੭`,
	}
	SignUpTab = Tab{
		URL:     domain.PathPageSignUp,
		Name:    "Signup",
		Kaomoji: `(⸝⸝ᵕᴗᵕ⸝⸝)`,
	}
)

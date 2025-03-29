package view

import "likemind/internal/domain"

type Tab struct {
	URL     string
	Name    string
	Kaomoji string
}

var (
	GroupTab = Tab{
		URL:     domain.PathPageGroup,
		Name:    "Group",
		Kaomoji: `\( ˙▿˙ )/\( ˙▿˙ )/`,
	}
	SearchTab = Tab{
		URL:     domain.PathPageSearch,
		Name:    "Search",
		Kaomoji: `⊂(￣▽￣)⊃`,
	}
	ProfileTab = Tab{
		// URL generated
		Name:    "My Profile",
		Kaomoji: `(„• ᴗ •„)`,
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

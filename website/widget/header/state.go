package header

import "likemind/internal/domain"

const (
	GroupTab   = "Groups"
	PeopleTab  = "People"
	ProfileTab = "Profile"
	SignInTab  = "Sign In"
	SignUpTab  = "Sign up"
)

var (
	AuthorizedTabs = State{
		Tabs: []TabState{
			{
				URL:     domain.PathPageGroup,
				Name:    GroupTab,
				Kaomoji: `\( ˙▿˙ )/\( ˙▿˙ )/`,
			},
			{
				URL:     "/search",
				Name:    PeopleTab,
				Kaomoji: `⊂(￣▽￣)⊃`,
			},
			{
				URL:     domain.PathPageUser,
				Name:    ProfileTab,
				Kaomoji: `(„• ᴗ •„)`,
			},
		},
	}

	UnauthorizedTabs = State{
		Tabs: []TabState{
			{
				URL:     domain.PathPageSignIn,
				Name:    SignInTab,
				Kaomoji: `(„• ֊ •„)੭`,
			},
			{
				URL:     domain.PathPageSignUp,
				Name:    SignUpTab,
				Kaomoji: `(⸝⸝ᵕᴗᵕ⸝⸝)`,
			},
		},
	}
)

type State struct {
	Tabs []TabState
}

type TabState struct {
	Name     string
	URL      string
	Kaomoji  string
	Selected bool
}

func (s State) Select(selected string) State {
	tabs := make([]TabState, 0, len(s.Tabs))

	for _, tab := range s.Tabs {
		if tab.Name == selected {
			tab.Selected = true
		}
		tabs = append(tabs, tab)
	}

	return State{Tabs: tabs}
}

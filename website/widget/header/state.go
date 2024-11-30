package header

const (
	GroupTab    = "Groups"
	PeopleTab   = "People"
	ProfileTab  = "Profile"
	SignInTab   = "Sign In"
	RegisterTab = "Register"
)

var (
	AuthorizedTabs = []TabState{
		{
			URL:     "/group",
			Name:    GroupTab,
			Kaomoji: `\( ˙▿˙ )/\( ˙▿˙ )/`,
		},
		{
			URL:     "/search/user",
			Name:    PeopleTab,
			Kaomoji: `⊂(￣▽￣)⊃`,
		},
		{
			URL:     "/profile",
			Name:    ProfileTab,
			Kaomoji: `(„• ᴗ •„)`,
		},
	}

	UnauthorizedTabs = []TabState{
		{
			URL:     "/sign_in",
			Name:    SignInTab,
			Kaomoji: `(„• ֊ •„)੭`,
		},
		{
			URL:     "/register",
			Name:    RegisterTab,
			Kaomoji: `(⸝⸝ᵕᴗᵕ⸝⸝)`,
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

func NewState(tabs []TabState, selectedTab string) State {
	result := make([]TabState, 0, len(tabs))
	for _, tab := range tabs {
		if tab.Name == selectedTab {
			tab.Selected = true
		}
		result = append(result, tab)
	}

	return State{
		Tabs: result,
	}
}

package common

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/rs/zerolog/log"
)

func RenderComponent(ctx context.Context, component templ.Component) io.Reader {
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()

		if err := component.Render(ctx, writer); err != nil {
			log.Err(err).Msg("component.Render")
		}
	}()

	return reader
}

func PrettyTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	switch {
	case diff < time.Minute:
		return "now"
	case diff < time.Hour:
		minutes := int(diff.Minutes())
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	case diff < 24*time.Hour:
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	default:
		days := int(diff.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	}
}

const (
	paramStart = '{'
	paramEnd   = '}'
)

// Example: FillPath("/user/{username}", map[string]string{"username": "john"}) returns "/user/john"
func FillPath(path string, params map[string]string) string {
	result := &strings.Builder{}

	var (
		isParam  bool
		paramIdx int
	)

	for idx, char := range path {
		switch char {
		case paramStart:
			isParam = true
			paramIdx = idx + 1
		case paramEnd:
			isParam = false
			paramName := path[paramIdx:idx]
			if param, ok := params[paramName]; ok {
				result.WriteString(param)
			}
		default:
			if isParam {
				continue
			}
			result.WriteRune(char)
		}
	}
	return result.String()
}

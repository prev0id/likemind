package widget_registry

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/a-h/templ"
)

const (
	nilTest = "nil-test"
)

var fileNameRegex = regexp.MustCompile(`^(\w+)_(\w+)_test\.json$`)

type WidgetGenerator func(data []byte) (templ.Component, error)

type Registry struct {
	widgets map[string]WidgetGenerator
}

func New() *Registry {
	return &Registry{
		widgets: widgets,
	}
}

func (r *Registry) ListTestsForWidget(widget string) ([]string, error) {
	testDir, err := os.ReadDir("./website/widget/" + widget + "/tests")
	if err != nil {
		return nil, fmt.Errorf("os.ReadDir: %w", err)
	}

	tests := make([]string, 0, len(testDir))
	tests = append(tests, nilTest)

	for _, testFile := range testDir {
		testName, withPrefix := strings.CutSuffix(testFile.Name(), "_test.json")
		if !withPrefix {
			continue
		}

		tests = append(tests, testName)
	}

	return tests, nil
}

func (r *Registry) ListWidgets() []string {
	widgets := make([]string, 0, len(r.widgets))

	for name := range r.widgets {
		widgets = append(widgets, name)
	}

	slices.Sort(widgets)

	return widgets
}

func (r *Registry) GetWidget(widgetName, testName string) (templ.Component, error) {
	widget, exists := r.widgets[widgetName]
	if !exists {
		return nil, fmt.Errorf("widget '%s' not exists", widgetName)
	}

	if testName == nilTest {
		return widget(nil)
	}

	file, err := os.ReadFile(fmt.Sprintf("website/widget/" + widgetName + "/tests/" + testName + "_test.json"))
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}

	return widget(file)
}

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFillPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		params   map[string]string
		expected string
	}{
		{
			name:     "Simple parameter replacement",
			path:     "/user/{username}",
			params:   map[string]string{"username": "john"},
			expected: "/user/john",
		},
		{
			name:     "Multiple parameter replacement",
			path:     "/user/{username}/post/{postId}",
			params:   map[string]string{"username": "john", "postId": "123"},
			expected: "/user/john/post/123",
		},
		{
			name:     "No parameters",
			path:     "/users/all",
			params:   map[string]string{},
			expected: "/users/all",
		},
		{
			name:     "Missing parameter value",
			path:     "/user/{username}",
			params:   map[string]string{},
			expected: "/user/",
		},
		{
			name:     "Empty path",
			path:     "",
			params:   map[string]string{"username": "john"},
			expected: "",
		},
		{
			name:     "Parameter at the beginning",
			path:     "{prefix}/user",
			params:   map[string]string{"prefix": "api"},
			expected: "api/user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FillPath(tt.path, tt.params)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

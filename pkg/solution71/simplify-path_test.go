package solution71

import (
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Test Case 1",
			path: "/home/",
			want: "/home",
		},
		{
			name: "Test Case 2",
			path: "/../",
			want: "/",
		},
		{
			name: "Test Case 3",
			path: "/home//foo/",
			want: "/home/foo",
		},
		{
			name: "Test Case 4",
			path: "/a/./b/../../c/",
			want: "/c",
		},
		{
			name: "Test Case 5",
			path: "/a////b",
			want: "/a/b",
		},
		{
			name: "Test Case 6",
			path: "/a/../../b",
			want: "/b",
		},
		{
			name: "Test Case 7",
			path: "/a/./b/././c",
			want: "/a/b/c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

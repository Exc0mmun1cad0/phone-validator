package phonenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Test 1",
			input: "8(982)583-86-16",
			want:  "+7-982-583-8616",
		},
		{
			name:  "Test 2",
			input: "+7 982 583 86-16",
			want:  "+7-982-583-8616",
		},
		{
			name:  "Test 3",
			input: "8-986-583-86-16",
			want:  "+7-986-583-8616",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.want, Normalize(test.input))
		})
	}
}

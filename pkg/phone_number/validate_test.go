package phonenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPhoneNum(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "OK test 1",
			input: "+7 982 583 8616",
			want:  true,
		},
		{
			name:  "OK test 2",
			input: "+79865838616",
			want:  true,
		},
        {
			name:  "OK test 3",
			input: "+7 (986) 583-8616",
			want:  true,
		},
		{
			name:  "OK test 4",
			input: "8 (912) 583-8616",
			want:  true,
		},
		{
			name:  "OK test 5",
			input: "89345838616",
			want:  true,
		},
		{
			name:  "Wrong opcode",
			input: "8 (933) 583-8616",
			want:  false,
		},
		{
			name:  "Wrong format with +7",
			input: "8 9345838616",
			want:  false,
		},
		{
			name:  "Wrong format with 8",
			input: "+7(934)5838616",
			want:  false,
		},        
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := IsValidPhoneNum(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

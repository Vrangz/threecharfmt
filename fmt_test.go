package threecharfmt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat_DefaultSeparator_NoReplacements(t *testing.T) {
	tests := []struct {
		s            string
		separator    rune
		replacements []Replacement
		want         string
		assertion    assert.ErrorAssertionFunc
	}{
		{
			s:            "",
			separator:    DefaultSeparator,
			replacements: []Replacement{},
			want:         "",
			assertion:    assert.NoError,
		},
		{
			s:            "A",
			separator:    DefaultSeparator,
			replacements: []Replacement{},
			want:         "A",
			assertion:    assert.NoError,
		},
		{
			s:            "AB",
			separator:    DefaultSeparator,
			replacements: []Replacement{},
			want:         "AB",
			assertion:    assert.NoError,
		},
		{
			s:            "ABC",
			separator:    DefaultSeparator,
			replacements: []Replacement{},
			want:         "ABC",
			assertion:    assert.NoError,
		},
		{
			s:            "ABCD",
			separator:    DefaultSeparator,
			replacements: []Replacement{},
			want:         "AB CD",
			assertion:    assert.NoError,
		},
		{
			s:            "ABCDE12345",
			separator:    DefaultSeparator,
			replacements: []Replacement{},
			want:         "ABC DE1 23 45",
			assertion:    assert.NoError,
		},
		{
			s:            "ABCDE123456",
			separator:    DefaultSeparator,
			replacements: []Replacement{},
			want:         "ABC DE1 234 56",
			assertion:    assert.NoError,
		},
		{
			s:            "ABCDE1234567",
			separator:    DefaultSeparator,
			replacements: []Replacement{},
			want:         "ABC DE1 234 567",
			assertion:    assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := Format(tt.s, tt.separator, tt.replacements...)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFormat_CustomSeparator_NoReplacements(t *testing.T) {
	tests := []struct {
		s            string
		separator    rune
		replacements []Replacement
		want         string
		assertion    assert.ErrorAssertionFunc
	}{
		{
			s:            "",
			separator:    rune('-'),
			replacements: []Replacement{},
			want:         "",
			assertion:    assert.NoError,
		},
		{
			s:            "A",
			separator:    rune('-'),
			replacements: []Replacement{},
			want:         "A",
			assertion:    assert.NoError,
		},
		{
			s:            "AB",
			separator:    rune('-'),
			replacements: []Replacement{},
			want:         "AB",
			assertion:    assert.NoError,
		},
		{
			s:            "ABC",
			separator:    rune('-'),
			replacements: []Replacement{},
			want:         "ABC",
			assertion:    assert.NoError,
		},
		{
			s:            "ABCD",
			separator:    rune('-'),
			replacements: []Replacement{},
			want:         "AB-CD",
			assertion:    assert.NoError,
		},
		{
			s:            "ABCDE12345",
			separator:    rune('-'),
			replacements: []Replacement{},
			want:         "ABC-DE1-23-45",
			assertion:    assert.NoError,
		},
		{
			s:            "ABCDE123456",
			separator:    rune('-'),
			replacements: []Replacement{},
			want:         "ABC-DE1-234-56",
			assertion:    assert.NoError,
		},
		{
			s:            "ABCDE1234567",
			separator:    rune('-'),
			replacements: []Replacement{},
			want:         "ABC-DE1-234-567",
			assertion:    assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := Format(tt.s, tt.separator, tt.replacements...)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFormat_CustomSeparator_EmptyStringReplacements(t *testing.T) {
	tests := []struct {
		s            string
		separator    rune
		replacements []Replacement
		want         string
		assertion    assert.ErrorAssertionFunc
	}{
		{
			s:         "",
			separator: rune('-'),
			replacements: []Replacement{
				{Old: " ", New: ""},
			},
			want:      "",
			assertion: assert.NoError,
		},
		{
			s:         "   A",
			separator: rune('-'),
			replacements: []Replacement{
				{Old: " ", New: ""},
			},
			want:      "A",
			assertion: assert.NoError,
		},
		{
			s:         "A   B",
			separator: rune('-'),
			replacements: []Replacement{
				{Old: " ", New: ""},
			},
			want:      "AB",
			assertion: assert.NoError,
		},
		{
			s:         "AB  C",
			separator: rune('-'),
			replacements: []Replacement{
				{Old: " ", New: ""},
			},
			want:      "ABC",
			assertion: assert.NoError,
		},
		{
			s:         "ABC  D",
			separator: rune('-'),
			replacements: []Replacement{
				{Old: " ", New: ""},
			},
			want:      "AB-CD",
			assertion: assert.NoError,
		},
		{
			s:         "ABC DE123 45",
			separator: rune('-'),
			replacements: []Replacement{
				{Old: " ", New: ""},
			},
			want:      "ABC-DE1-23-45",
			assertion: assert.NoError,
		},
		{
			s:         "AB CDE123 456",
			separator: rune('-'),
			replacements: []Replacement{
				{Old: " ", New: ""},
			},
			want:      "ABC-DE1-234-56",
			assertion: assert.NoError,
		},
		{
			s:         "ABCD E12345 67",
			separator: rune('-'),
			replacements: []Replacement{
				{Old: " ", New: ""},
			},
			want:      "ABC-DE1-234-567",
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := Format(tt.s, tt.separator, tt.replacements...)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

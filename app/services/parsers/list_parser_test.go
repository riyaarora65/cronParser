package parsers

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewListParser(t *testing.T) {
	testCases := []struct {
		field        string
		expectedList []int
		expectedErr  bool
	}{
		{"1,2,3,4,5", []int{1, 2, 3, 4, 5}, false},     // valid list of integers
		{"10,20,30", []int{10, 20, 30}, false},        // valid list of integers
		{"", nil, true},                              // empty field should return error
		{"1,2,3,4,invalid", nil, true},                // invalid field with non-integer value
		{"1-5", nil, true},                            // range format should return error
		{"*/5", nil, true},                            // step format should return error
	}

	for _, tc := range testCases {
		t.Run("Parsing "+tc.field, func(t *testing.T) {
			parser, err := NewListParser(tc.field)

			if tc.expectedErr {
				assert.Error(t, err, "Expected error for field: %s", tc.field)
				assert.Nil(t, parser, "Parser should be nil for field: %s", tc.field)
			} else {
				assert.NoError(t, err, "Unexpected error for field: %s", tc.field)
				assert.NotNil(t, parser, "Parser should not be nil for field: %s", tc.field)
				assert.Equal(t, tc.expectedList, parser.values, "Parsed values mismatch for field: %s", tc.field)
			}
		})
	}
}

func TestListParser_ExpandField(t *testing.T) {
	testCases := []struct {
		values       []int
		expectedList []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{10, 20, 30}, []int{10, 20, 30}},
		{[]int{}, []int{}}, // empty list case
	}

	for _, tc := range testCases {
		t.Run("Expanding values "+strconv.Itoa(len(tc.values)), func(t *testing.T) {
			parser := &ListParser{values: tc.values}
			expanded, err := parser.ExpandField()

			assert.NoError(t, err, "Unexpected error expanding values")
			assert.Equal(t, tc.expectedList, expanded, "Expanded values mismatch")
		})
	}
}

package parsers

import (
	"mygolearning/deliverooProject/app/utils"
	"reflect"
	"testing"
)

func TestStepParser_ExpandField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		min, max int
		expected []int
		wantErr  *string
	}{
		// Valid Cases
		{"Step 2 between 0-10", "*/2", 0, 10, []int{0, 2, 4, 6, 8, 10}, nil},
		{"Step 3 between 1-12", "*/3", 1, 12, []int{1, 4, 7, 10}, nil},
		{"Step 1 between 5-5", "*/1", 5, 5, []int{5}, nil},
		{"Step 5 between 10-20", "*/5", 10, 20, []int{10, 15, 20}, nil},

		// Edge Cases
		{"Step 1 between 0-1", "*/1", 0, 1, []int{0, 1}, nil},
		{"Step 2 between 0-1", "*/2", 0, 1, []int{0}, utils.StrPtr("invalid step value")}, // No value in range

		// Negative Cases
		{"Negative step", "*/-1", 0, 10, nil, utils.StrPtr("invalid step value")},        // Negative step
		{"Zero step", "*/0", 0, 10, nil, utils.StrPtr("invalid step value")},             // Zero step
		{"Step greater than max", "*/15", 0, 10, nil, utils.StrPtr("invalid step value")}, // Step greater than max

		// Invalid Case
		{"Invalid step format", "*/", 0, 10, nil, utils.StrPtr("invalid step format")},     // Invalid step format
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := NewStepParser(tt.field, tt.min, tt.max)
			if err != nil && err.Error() != *tt.wantErr {
				t.Errorf("NewStepParser() error = %v, wantErr %v", err, *tt.wantErr)
				return
			}
			if err == nil {
				got, err := parser.ExpandField()
				if err != nil && err.Error() != *tt.wantErr {
					t.Errorf("StepParser.ExpandField() error = %v, wantErr %v", err, *tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("StepParser.ExpandField() = %v, want %v", got, tt.expected)
				}
			}
		})
	}
}

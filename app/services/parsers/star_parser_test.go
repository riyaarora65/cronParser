package parsers

import (
	"mygolearning/deliverooProject/app/utils"
	"testing"
)

func TestStarParser_ExpandField(t *testing.T) {
	tests := []struct {
		name    string
		min     int
		max     int
		want    []int
		wantErr bool
	}{
		{"Valid range 0-10", 0, 10, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
		{"Valid range 5-15", 5, 15, []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, false},
		{"Single value range 7-7", 7, 7, []int{7}, false},
		{"Negative range -5-5", -5, 5, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, false},
		{"Zero range 0-0", 0, 0, []int{0}, false},
		{"Invalid range 10-5", 10, 5, nil, true},
		{"Negative minimum -10 to 5", -10, 5, []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewStarParser(tt.min, tt.max)
			got, err := parser.ExpandField()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpandField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !utils.IsEqual(got, tt.want) {
				t.Errorf("ExpandField() = %v, want %v", got, tt.want)
			}
		})
	}
}



package parsers

import (
	"mygolearning/deliverooProject/app/utils"
	"testing"
)

func TestRangeParser_ExpandField(t *testing.T) {
	tests := []struct {
		name    string
		field   string
		want    []int
		wantErr *string
	}{
		{"Valid range 1-5", "1-5", []int{1, 2, 3, 4, 5}, utils.StrPtr("")},
		{"Valid range 0-3", "0-3", []int{0, 1, 2, 3}, utils.StrPtr("")},
		{"Single value range 7-7", "7-7", []int{7}, utils.StrPtr("")},
		{"Zero range 0-0", "0-0", []int{0}, utils.StrPtr("invalid Range: start value cannot be greater than end")},
		{"Invalid range 5-1", "5-1", nil, utils.StrPtr("invalid Range: start value cannot be greater than end value")},
		{"Invalid range format no dash", "5,1", nil, utils.StrPtr("invalid range format")},
		{"Invalid range format no second value", "5-", nil, utils.StrPtr("invalid range format")},
		{"Invalid range format non-numeric", "a-b", nil, utils.StrPtr("invalid field value")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := NewRangeParser(tt.field)
			if err != nil && err.Error() != *tt.wantErr {
				t.Errorf("NewRangeParser() error = %v, wantErr %v", err, *tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			got, err := parser.ExpandField()
			if err != nil && err.Error() != *tt.wantErr {
				t.Errorf("ExpandField() error = %v, wantErr %v", err, *tt.wantErr)
				return
			}
			if !utils.IsEqual(got, tt.want) {
				t.Errorf("ExpandField() = %v, want %v", got, tt.want)
			}
		})
	}
}

package factory

import (
	"testing"

	cronParsers "mygolearning/deliverooProject/app/services/parsers"
	"github.com/stretchr/testify/assert"
)

func TestCronFactory_CreateParser(t *testing.T) {
	factory := NewCronFactory()

	tests := []struct {
		name      string
		fields    []string
		wantTypes []interface{}
		wantErr   bool
	}{
		{
			name:      "Valid star fields",
			fields:    []string{"*", "*", "*", "*", "*"},
			wantTypes: []interface{}{&cronParsers.StarParser{}, &cronParsers.StarParser{}, &cronParsers.StarParser{}, &cronParsers.StarParser{}, &cronParsers.StarParser{}},
			wantErr:   false,
		},
		{
			name:      "Valid step fields",
			fields:    []string{"*/2", "*/3", "*/5", "*/1", "*/4"},
			wantTypes: []interface{}{&cronParsers.StepParser{}, &cronParsers.StepParser{}, &cronParsers.StepParser{}, &cronParsers.StepParser{}, &cronParsers.StepParser{}},
			wantErr:   false,
		},
		{
			name:      "Valid range fields",
			fields:    []string{"0-10", "0-23", "1-31", "1-12", "1-7"},
			wantTypes: []interface{}{&cronParsers.RangeParser{}, &cronParsers.RangeParser{}, &cronParsers.RangeParser{}, &cronParsers.RangeParser{}, &cronParsers.RangeParser{}},
			wantErr:   false,
		},
		{
			name:      "Valid list fields",
			fields:    []string{"0,10", "0,23", "1,31", "1,12", "1,7"},
			wantTypes: []interface{}{&cronParsers.ListParser{}, &cronParsers.ListParser{}, &cronParsers.ListParser{}, &cronParsers.ListParser{}, &cronParsers.ListParser{}},
			wantErr:   false,
		},
		{
			name:      "Invalid field format",
			fields:    []string{"invalid", "*", "*", "*", "*"},
			wantTypes: nil,
			wantErr:   true,
		},
		{
			name:      "Invalid range field",
			fields:    []string{"10-5", "*", "*", "*", "*"},
			wantTypes: nil,
			wantErr:   true,
		},
		{
			name:      "Invalid step field",
			fields:    []string{"*/invalid", "*", "*", "*", "*"},
			wantTypes: nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsers, err := factory.CreateParser(tt.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, len(tt.wantTypes), len(parsers))
				for i, parser := range parsers {
					assert.IsType(t, tt.wantTypes[i], parser)
				}
			}
		})
	}
}

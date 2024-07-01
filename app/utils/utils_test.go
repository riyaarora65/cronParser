package utils

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	customErr "mygolearning/deliverooProject/app/errors"
// )

// func TestParseInt(t *testing.T) {
// 	tests := []struct {
// 		name             string
// 		value            string
// 		minValue, maxValue int
// 		expectedNum      int
// 		expectedError    error
// 	}{
// 		{
// 			name:          "Valid integer within range",
// 			value:         "5",
// 			minValue:      0,
// 			maxValue:      10,
// 			expectedNum:   5,
// 			expectedError: nil,
// 		},
// 		{
// 			name:          "Valid minimum value",
// 			value:         "0",
// 			minValue:      0,
// 			maxValue:      10,
// 			expectedNum:   0,
// 			expectedError: nil,
// 		},
// 		{
// 			name:          "Valid maximum value",
// 			value:         "10",
// 			minValue:      0,
// 			maxValue:      10,
// 			expectedNum:   10,
// 			expectedError: nil,
// 		},
// 		{
// 			name:          "Value below minimum",
// 			value:         "-1",
// 			minValue:      0,
// 			maxValue:      10,
// 			expectedNum:   0,
// 			expectedError: customErr.NewCronError(customErr.ErrInvalidField, "-1"),
// 		},
// 		{
// 			name:          "Value above maximum",
// 			value:         "15",
// 			minValue:      0,
// 			maxValue:      10,
// 			expectedNum:   0,
// 			expectedError: customErr.NewCronError(customErr.ErrInvalidField, "15"),
// 		},
// 		{
// 			name:          "Invalid integer format",
// 			value:         "abc",
// 			minValue:      0,
// 			maxValue:      10,
// 			expectedNum:   0,
// 			expectedError: customErr.NewCronError(customErr.ErrInvalidField, "abc"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			num, err := ParseInt(tt.value, tt.minValue, tt.maxValue)
// 			if tt.expectedError != nil {
// 				assert.Error(t, err)
// 				assert.Equal(t, tt.expectedError, err)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tt.expectedNum, num)
// 			}
// 		})
// 	}
// }

// func TestParseRange(t *testing.T) {
// 	tests := []struct {
// 		name             string
// 		field            string
// 		expectedStart    int
// 		expectedEnd      int
// 		expectedError    error
// 	}{
// 		{
// 			name:          "Valid range",
// 			field:         "1-5",
// 			expectedStart: 1,
// 			expectedEnd:   5,
// 			expectedError: nil,
// 		},
// 		{
// 			name:          "Negative range",
// 			field:         "-1--5",
// 			expectedStart: 0,
// 			expectedEnd:   0,
// 			expectedError: customErr.NewCronError(customErr.ErrInvalidRange, "-1--5"),
// 		},
// 		{
// 			name:          "Invalid range format",
// 			field:         "1",
// 			expectedStart: 0,
// 			expectedEnd:   0,
// 			expectedError: customErr.NewCronError(customErr.ErrInvalidRange, "1"),
// 		},
// 		{
// 			name:          "End before start",
// 			field:         "5-1",
// 			expectedStart: 0,
// 			expectedEnd:   0,
// 			expectedError: customErr.NewCronError(customErr.ErrInvalidRange, "5-1"),
// 		},
// 		{
// 			name:          "Non-numeric start",
// 			field:         "a-5",
// 			expectedStart: 0,
// 			expectedEnd:   0,
// 			expectedError: customErr.NewCronError(customErr.ErrInvalidRange, "a"),
// 		},
// 		{
// 			name:          "Non-numeric end",
// 			field:         "1-b",
// 			expectedStart: 0,
// 			expectedEnd:   0,
// 			expectedError: customErr.NewCronError(customErr.ErrInvalidRange, "b"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			start, end, err := ParseRange(tt.field)
// 			if tt.expectedError != nil {
// 				assert.Error(t, err)
// 				assert.Equal(t, tt.expectedError, err)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tt.expectedStart, start)
// 				assert.Equal(t, tt.expectedEnd, end)
// 			}
// 		})
// 	}
// }

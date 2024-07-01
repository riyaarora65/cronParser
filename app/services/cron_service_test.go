package services

import (
	"bytes"
	"io"
	"mygolearning/deliverooProject/app/utils"
	"os"
	"testing"
)

func TestCronService_ParseAndPrint(t *testing.T) {
	tests := []struct {
		cronString     string
		expectedOutput string
	}{
		{
		cronString: "*/15 1-3 1,15 1-12 1-5 /usr/bin/find",
		expectedOutput: `minute        0 15 30 45
hour          1 2 3
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
`,
	},
	{
		cronString: "*/15 1-3 1,15 1-12 1-5 /bin/echo",
		expectedOutput: `minute        0 15 30 45
hour          1 2 3
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /bin/echo
`,
	},
	// Invalid step value for minute
	{
		cronString: "*/61 1-3 1,15 1-12 1-5 /bin/echo",
		expectedOutput: "value out of range\n",
	},
	// Invalid range for hour
	{
		cronString: "*/15 24-25 1,15 1-12 1-5 /bin/echo",
		expectedOutput: "value out of range\n",
	},
	// Edge case: minimum and maximum values for minute
	{
		cronString: "0 1-3 1,15 1-12 1-5 /bin/echo",
		expectedOutput: `minute        0
hour          1 2 3
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /bin/echo
`,
	},
	// Edge case: step value for minute
	{
		cronString: "*/5 1-3 1,15 1-12 1-5 /bin/echo",
		expectedOutput: `minute        0 5 10 15 20 25 30 35 40 45 50 55
hour          1 2 3
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /bin/echo
`,
	},
	// Invalid range for day of month
	{
		cronString: "*/15 1-3 0,32 1-12 1-5 /bin/echo",
		expectedOutput: "value out of range\n",
	},
	// Invalid range for month
	{
		cronString: "*/15 1-3 1,15 0-13 1-5 /bin/echo",
		expectedOutput: "value out of range\n",
	},
	// Invalid range for day of week
	{
		cronString: "*/15 1-3 1,15 1-12 0-7 /bin/echo",
		expectedOutput: "value out of range\n",
	},
	// Combination of valid and invalid fields
	{
		cronString: "*/60 1-25 1,32 1-12 1-7 /bin/echo",
		expectedOutput: "value out of range\n",
	},
	// Missing fields
	{
		cronString: "*/60 1-25 1,32 /bin/echo",
		expectedOutput: "invalid field format\n",
	},
	//Invalid characters present
	{
		cronString: "&/60 1-25 1#32 1-12 /bin/echo mixed case",
		expectedOutput: "invalid field value\n",
	},
	//Extra fields if present will be taken as command
	{
		cronString: "15 10 1 * * /usr/bin/find extra fields present",
		expectedOutput: `minute         15
        hour           10
        day of month   1
        month          1 2 3 4 5 6 7 8 9 10 11 12
        day of week    1 2 3 4 5 6 7
        command        /usr/bin/find extra fields present`,
	},
	//putting extra fields in starting
	{
		cronString: "extra 10 1 * * /usr/bin/find 15 fields present",
		expectedOutput: `invalid field value`,
	},
}

	for _, test := range tests {
		t.Run(test.cronString, func(t *testing.T) {

			// Create a pipe to capture stdout
			reader, writer, _ := os.Pipe()
			defer reader.Close()
			defer writer.Close()

			// Redirect stdout to the writer
			oldStdout := os.Stdout
			os.Stdout = writer

			NewCronService().ParseAndPrint(test.cronString)

			// Reset stdout
			writer.Close()
			os.Stdout = oldStdout

			// Read the output
			var buf bytes.Buffer
			io.Copy(&buf, reader)
			actualOutput := buf.String()

			if utils.RemoveSpacesAndLines(actualOutput) != utils.RemoveSpacesAndLines(test.expectedOutput) {
				t.Errorf("Test failed for cronString: %s\nExpected:\n%s\nGot:\n%s\n", test.cronString, test.expectedOutput, actualOutput)
			}
		})
	}
}

package handler

import (
	"testing"
)

func TestAppExistCheck(t *testing.T) {
	cases := []struct {
		input         string
		errorExpected bool
		message       string
	}{
		{
			input:         "Mail ; echo Hello!",
			errorExpected: true,
			message:       "Should not be able to use command that is not only Application name.",
		},
		{
			input:         "Mail",
			errorExpected: false,
			message:       "Should be able to use command that is only Application name.",
		},
	}

	for caseNum, c := range cases {
		actualStatus := AppExistCheck(c.input)
		actual := actualStatus != nil
		if actual != c.errorExpected {
			t.Errorf("No.%d: Expect %v, But got %v.\n Message: %v", caseNum, c.errorExpected, actual, c.message)
		}
	}
}

package customer

import (
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T) {

	type TestCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []TestCase{
		{
			test:        "Empty Name Validation",
			name:        "",
			expectedErr: ErrInvalidPerson,
		},
		{
			test:        "Valid Name",
			name:        "Peter",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewCustomer(tc.name)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got error %v", tc.expectedErr, err)
			}
		})
	}
}

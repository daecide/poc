package math

import "testing"

func Test_Square(t *testing.T) {
	testCases := []struct{
		Param int
		ExpectedResult int
	}{
		{1, 1},
		{2, 4},
		{3, 9},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			if resp := Square(tc.Param); tc.ExpectedResult != resp {
				t.Fatalf("Expected result: %v, got: %v", tc.ExpectedResult, resp)
			}
		})
	}
}
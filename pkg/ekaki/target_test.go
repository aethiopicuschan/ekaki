package ekaki_test

import (
	"testing"

	"github.com/aethiopicuschan/ekaki/pkg/ekaki"
	"github.com/stretchr/testify/require"
)

func TestTargetFromExpr(t *testing.T) {
	testcases := []struct {
		name          string
		expr          string
		expectedValue ekaki.Target
		expectedErr   error
	}{
		{
			name:          "jpg",
			expr:          "jpg",
			expectedValue: ekaki.TargetJPG,
			expectedErr:   nil,
		},
		{
			name:          "jpeg",
			expr:          "jpeg",
			expectedValue: ekaki.TargetJPG,
			expectedErr:   nil,
		},
		{
			name:          "png",
			expr:          "png",
			expectedValue: ekaki.TargetPNG,
			expectedErr:   nil,
		},
		{
			name:          "unsupported",
			expr:          "unsupported",
			expectedValue: "",
			expectedErr:   ekaki.ErrUnsupported,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()
			v, err := ekaki.TargetFromExpr(testcase.expr)
			require.Equal(t, testcase.expectedValue, v)
			require.Equal(t, testcase.expectedErr, err)
		})
	}

}

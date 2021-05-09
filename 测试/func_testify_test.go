package funcTest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_add2(t *testing.T) {
	t.Run("test_assert_equal", func(t *testing.T) {
		assert.Equal(t, 3, add(1, 2))

	})

	t.Run("test_require_equal", func(t *testing.T) {
		require.Equal(t, 3, add(1, 2))
	})

}

package ff_common_lib

import "testing"

func TestSumTesting(t *testing.T) {
	count := Sum(1, 1)

	if count != 2 {
		t.Errorf("Retcode, got: %d", count)
	}
}

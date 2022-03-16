package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

type testRetriever struct {
}

func (c *testRetriever) GetValidators(blockID ids.ID) (Set, error) {
	return NewDefaultSet(constants.CostonID), nil
}

func TestGetValidatorsFromDefaultSet(t *testing.T) {
	testBlockID := ids.ID{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
	}
	cr := NewCachingRetriever(&testRetriever{})
	set, err := cr.GetValidators(testBlockID)
	require.NoError(t, err)
	assert.Equal(t, uint64(1000000), set.Weight())
}

func TestUpdaterFromDefaultSet(t *testing.T) {
	testBlockID := ids.ID{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
	}
	u := NewUpdater(NewDefaultSet(constants.CostonID), &testRetriever{})
	err := u.UpdateValidators(testBlockID)
	assert.NoError(t, err)
}

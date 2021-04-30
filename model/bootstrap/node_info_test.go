package bootstrap_test

import (
	"testing"

	"github.com/onflow/flow-go/model/bootstrap"
	"github.com/onflow/flow-go/model/flow/order"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	nodes := unittest.NodeInfosFixture(20)
	nodes = bootstrap.Sort(nodes, order.Canonical)
	require.True(t, bootstrap.ToIdentityList(nodes).Sorted(order.Canonical))
}

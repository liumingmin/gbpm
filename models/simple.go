package models

type TestNode struct {
	InstanceNode
}

func (n *TestNode) TableName() string {
	return "test_node"
}

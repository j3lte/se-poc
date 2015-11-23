package lib

func GetNewNetworkMap()(*NetworkMap) {
  return &NetworkMap{}
}

func GetNewEdge()(*Edge) {
  return &Edge{}
}

func GetNewNode()(*Node) {
  return &Node{}
}

type NetworkMap struct {
  Nodes []Node   `json:"nodes" bson:"nodes"`
  Edges []Edge   `json:"edges" bson:"edges"`
}

type Node struct {
  Id       int     `json:"id" bson:"id"`
  Label    string  `json:"label" bson:"label"`
  Group    string  `json:"group" bson:"group"`
}

type Edge struct {
  From  int     `json:"from" bson:"from"`
  To    int     `json:"to" bson:"to"`
}

// Merges two network maps
func (n *NetworkMap)Absorb(f *NetworkMap) {
  n.Nodes = append(n.Nodes, f.Nodes...)
  n.Edges = append(n.Edges, f.Edges...)
}
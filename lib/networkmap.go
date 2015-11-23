package lib

func GetNewNetworkMap() *NetworkMap {
  return &NetworkMap{}
}

func GetNewEdge() *Edge {
  return &Edge{}
}

func GetNewNode() *Node {
  return &Node{}
}

type NetworkMap struct {
  Nodes []Node `json:"nodes" bson:"nodes"`
  Edges []Edge `json:"edges" bson:"edges"`
}

// Adds a node internally
func (n *NetworkMap) AddNode(node *Node) {
  n.Nodes = append(n.Nodes, *node)
}

// Adds an edge internall
func (n *NetworkMap) AddEdge(edge *Edge) {
  n.Edges = append(n.Edges, *edge)
}

func (n *NetworkMap) CreateEdge(from int, to int) {
  edge := GetNewEdge()
  edge.From = from
  edge.To = to

  n.AddEdge(edge)
}

func (n *NetworkMap) CreateNode(id int, label string, group string) {
  node := GetNewNode()
  node.Id = id
  node.Label = label
  node.Group = group

  n.AddNode(node)
}

type Node struct {
  Id    int    `json:"id" bson:"id"`
  Label string `json:"label" bson:"label"`
  Group string `json:"group" bson:"group"`
}

type Edge struct {
  From int `json:"from" bson:"from"`
  To   int `json:"to" bson:"to"`
}

// Merges two network maps
func (n *NetworkMap) Absorb(f *NetworkMap) {
  n.Nodes = append(n.Nodes, f.Nodes...)
  n.Edges = append(n.Edges, f.Edges...)
}

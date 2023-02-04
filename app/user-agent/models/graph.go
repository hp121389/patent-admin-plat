package models

type Node struct {
	NodeId            string  `grom:"primaryKey;size:128" json:"id" `
	NodeName          string  `grom:"size:128" json:"name"`
	NodeSymbolizeSize float32 `grom:"" json:"symbolSize"`
	NodeValue         int     `grom:"" json:"value"`
	NodeCategory      int     `grom:"" json:"category"`
}

type PreNode struct {
	NodeId            int
	NodeName          string
	NodeSymbolizeSize float32
	NodeValue         int
	NodeCategory      int
}

// Link n--n的关系
type Link struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Value  int    `json:"value"`
}
type PreLink struct {
	Source int
	Target int
	Value  int
}

type Graph struct {
	Nodes []Node `json:"nodes"`
	Links []Link `json:"links"`
}

func (e *Node) TableName() string {
	return "Node"
}

type InventorPatent struct {
	InventorId int
	PatentId   int
}
type SimplifiedNode struct { //has some basic information of nodes
	Id                 int
	Name               string
	TheNumberOfPatents int
	InTheGraph         bool
}
type Inventor struct {
	Id                 int
	Name               string
	TheNumberOfPatents int
	PatentsId          []int
}
type KeyWord struct {
	Id                 int
	Name               string
	TheNumberOfPatents int
	PatentsId          []int
}

package models

type Node struct {
	NodeId            string  `grom:"primaryKey;size:128" json:"id" `
	NodeName          string  `grom:"size:128" json:"name"` //用户名
	NodeSymbolizeSize float32 `grom:"" json:"symbolSize"`
	NodeX             float32 `grom:"" json:"x"`
	NodeY             float32 `grom:"" json:"y"`
	NodeValue         int     `grom:"" json:"value"` //重复次数
	NodeCategory      int     `grom:"" json:"category"`
}

func (e *Node) TableName() string {
	return "Node"
}

// n--n的关系
type Link struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Value  int    `json:"value"`
}

package dto

import "go-admin/app/user-agent/models"

type GraphResult struct {
	Nnodes interface{} `json:"Nnodes"`
	Links  interface{} `json:"Links"`
}

func (g *GraphResult) GetNodesAndLinks(nodes *[]models.Node, links *[]models.Link) {
	g.Nnodes = *nodes
	g.Links = *links
	return
}

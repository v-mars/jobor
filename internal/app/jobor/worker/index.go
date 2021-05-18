package worker

import "jobor/internal/models"

var tbName = models.JoborWorker

type update struct {
	Status *string `json:"status"`
}


type ShowRoutingKey struct {
	RoutingKey string `json:"routing_key"`
}

func (ShowRoutingKey) TableName() string {
	return tbName
}

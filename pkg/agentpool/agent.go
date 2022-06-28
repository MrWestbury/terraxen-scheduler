package agentpool

import "time"

type Agent struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"last_seen"`
}

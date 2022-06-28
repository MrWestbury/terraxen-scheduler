package agentpool

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

var (
	ErrAgentAlreadyRegistered = errors.New("agent with that id already resigtered in pool")
)

type Agentpool struct {
	Id     string
	Name   string
	Agents map[string]*Agent
}

func NewAgentpool(name string) (*Agentpool, error) {
	poolid, err := uuid.NewUUID()
	if err != nil {
		log.Printf("failed to create pool %s: %v", name, err)
		return nil, err
	}

	pool := &Agentpool{
		Id:     poolid.String(),
		Name:   name,
		Agents: make(map[string]*Agent),
	}

	return pool, nil
}

func (ap *Agentpool) RegisterAgent(agent *Agent) error {
	_, exists := ap.Agents[agent.Id]
	if exists {
		return ErrAgentAlreadyRegistered
	}

	ap.Agents[agent.Id] = agent
	return nil
}

func (ap *Agentpool) UnregisterAgent(agentId string) {
	delete(ap.Agents, agentId)
}

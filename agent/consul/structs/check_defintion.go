package structs

import (
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/types"
)

// CheckDefinition is used to JSON decode the Check definitions
type CheckDefinition struct {
	ID        types.CheckID
	Name      string
	Notes     string
	ServiceID string
	Token     string
	Status    string

	// Copied fields from CheckType without the fields
	// already present in CheckDefinition:
	//
	//   ID (CheckID), Name, Status, Notes
	//
	Script                         string
	HTTP                           string
	Header                         map[string][]string
	Method                         string
	TCP                            string
	Interval                       time.Duration
	DockerContainerID              string
	Shell                          string
	TLSSkipVerify                  bool
	Timeout                        time.Duration
	TTL                            time.Duration
	DeregisterCriticalServiceAfter time.Duration
}

func (c *CheckDefinition) HealthCheck(node string) *HealthCheck {
	health := &HealthCheck{
		Node:      node,
		CheckID:   c.ID,
		Name:      c.Name,
		Status:    api.HealthCritical,
		Notes:     c.Notes,
		ServiceID: c.ServiceID,
	}
	if c.Status != "" {
		health.Status = c.Status
	}
	if health.CheckID == "" && health.Name != "" {
		health.CheckID = types.CheckID(health.Name)
	}
	return health
}

func (c *CheckDefinition) CheckType() *CheckType {
	return &CheckType{
		CheckID:           c.ID,
		Name:              c.Name,
		Script:            c.Script,
		HTTP:              c.HTTP,
		TCP:               c.TCP,
		Interval:          c.Interval,
		DockerContainerID: c.DockerContainerID,
		Shell:             c.Shell,
		TLSSkipVerify:     c.TLSSkipVerify,
		Timeout:           c.Timeout,
		TTL:               c.TTL,
		DeregisterCriticalServiceAfter: c.DeregisterCriticalServiceAfter,
		Status: c.Status,
		Notes:  c.Notes,
	}
}

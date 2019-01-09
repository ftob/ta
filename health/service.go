package health

import (
	"context"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"time"
)

const (
	StatusPass = "pass"
	StatusFail = "fail"
	StatusWarn = "warn"
)

type Service interface {
	Health() Status
}

type service struct {
	ctx context.Context
}

type Notes []string

type Avg struct {
	ComponentId   string  `json:"componentId"`
	ComponentType string  `json:"componentType"`
	MetricValue   float64 `json:"metric_value"`
	MetricUnit    string  `json:"metric_unit"`
	Status        string  `json:"status"`
	Time          string  `json:"time"`
}

type Memory struct {
	ComponentId   string  `json:"componentId"`
	ComponentType string  `json:"componentType"`
	MetricValue   float64 `json:"metric_value"`
	MetricUnit    string  `json:"metric_unit"`
	Status        string  `json:"status"`
	Time          string  `json:"time"`
}

type Uptime struct {
	ComponentType string  `json:"componentType"`
	MetricValue   float64 `json:"metric_value"`
	MetricUnit    string  `json:"metric_unit"`
	Status        string  `json:"status"`
	Time          string  `json:"time"`
	ComponentId   string  `json:"componentId"`
}

type Details struct {
	CPUUtilization    Avg    `json:"cpu:utilization"`
	MemoryUtilization Memory `json:"memory:utilization"`
	Uptime            Uptime `json:"uptime"`
}

// https://tools.ietf.org/id/draft-inadarei-api-health-check-01.html#rfc.section.3
type Status struct {
	Status      string            `json:"status"`
	Version     string            `json:"version"`
	ReleaseID   string            `json:"releaseID"`
	Notes       Notes             `json:"notes"`
	Output      string            `json:"output"`
	Details     Details           `json:"details"`
	Links       map[string]string `json:"links"`
	ServiceID   string            `json:"serviceID"`
	Description string            `json:"description"`
}

func NewService(ctx context.Context) Service {
	return &service{ctx}
}

func (s *service) Health() Status {
	ram := s.newMemory()
	up := s.newUptime()
	avg := s.newCpu()

	status := StatusPass

	if ram.Status == StatusWarn || avg.Status == StatusWarn || up.Status == StatusWarn {
		status = StatusWarn
	}

	if ram.Status == StatusFail || avg.Status == StatusFail || up.Status == StatusFail {
		status = StatusFail
	}

	return Status{
		Status:    status,
		Version:   s.ctx.Value("Version").(string),
		ReleaseID: s.ctx.Value("Version").(string),
		Details: Details{avg, ram, up},
		ServiceID: s.ctx.Value("ServiceID").(string),
		Description: "",
	}

}

func (s *service) newUptime() Uptime {
	t := time.Since(s.ctx.Value("startTime").(time.Time))

	return Uptime{
		Time:          time.Now().String(),
		MetricValue:   t.Seconds(),
		MetricUnit:    "sec",
		ComponentType: s.ctx.Value("ComponentType").(string),
		ComponentId:   s.ctx.Value("ComponentId").(string),
		Status: StatusPass,
	}
}

func (s *service) newCpu() Avg {
	l, _ := load.Avg()
	i, _ := cpu.Info()
	avg := l.Load15

	status := StatusPass

	if avg > float64(len(i)) {
		status = StatusWarn
	}

	return Avg{
		Time:          time.Now().String(),
		MetricUnit:    "avg",
		MetricValue:   avg,
		Status:        status,
		ComponentType: s.ctx.Value("ComponentType").(string),
		ComponentId:   s.ctx.Value("ComponentId").(string),
	}
}

func (s *service) newMemory() Memory {
	v, _ := mem.VirtualMemory()

	status := StatusPass

	if v.UsedPercent > 80. {
		status = StatusWarn
	} else if v.UsedPercent > 95 {
		status = StatusFail
	}
	
	return Memory{
		Status:        status,
		ComponentId:   s.ctx.Value("ComponentId").(string),
		ComponentType: s.ctx.Value("ComponentType").(string),
		MetricValue:   v.UsedPercent,
		MetricUnit:    "percent",
		Time:          time.Now().String(),
	}
}

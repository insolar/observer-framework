// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/observer-framework/blob/master/LICENSE.md.

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/profefe/profefe/agent"

	"github.com/insolar/observer-framework/configuration"
)

// Profefe is a service to pushing the pprof data to pprof collector
type Profefe struct {
	startAgent  bool
	address     string
	serviceName string
	labels      []string

	hasStarted bool
	pffAgent   *agent.Agent
}

func NewProfefe(cfg configuration.Profefe, serviceName string) *Profefe {
	labels := make([]string, 0)
	for _, v := range strings.Split(cfg.Labels, ",") {
		labels = append(labels, strings.Trim(v, " "))
	}
	return &Profefe{
		startAgent:  cfg.StartAgent,
		address:     cfg.Address,
		serviceName: serviceName,
		labels:      labels,
	}
}

func (p *Profefe) Start() error {
	if !p.hasStarted {
		pffAgent, err := agent.Start(p.address,
			p.serviceName,
			agent.WithCPUProfile(10*time.Second),
			agent.WithHeapProfile(),
			agent.WithBlockProfile(),
			agent.WithMutexProfile(),
			agent.WithGoroutineProfile(),
			agent.WithThreadcreateProfile(),
			agent.WithLogger(agentLogger),
			agent.WithLabels(p.labels...),
		)

		if err != nil {
			return errors.Wrap(err, "cannot start profefe agent")
		}
		p.pffAgent = pffAgent
	}
	return nil
}

func (p *Profefe) Stop() error {
	if p.hasStarted {
		err := p.pffAgent.Stop()
		p.pffAgent = nil
		return errors.Wrap(err, "cannot stop profefe agent")
	}
	return nil
}

func agentLogger(format string, v ...interface{}) {
	log.Println(fmt.Sprintf(format, v...))
}

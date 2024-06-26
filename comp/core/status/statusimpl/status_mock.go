// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

//go:build test

package statusimpl

import (
	"go.uber.org/fx"

	"github.com/DataDog/datadog-agent/comp/core/status"
	"github.com/DataDog/datadog-agent/pkg/util/fxutil"
)

type statusMock struct {
}

func (s *statusMock) GetStatus(string, bool, ...string) ([]byte, error) {
	return []byte{}, nil
}

func (s *statusMock) GetStatusBySections([]string, string, bool) ([]byte, error) {
	return []byte{}, nil
}

func (s *statusMock) GetSections() []string {
	return []string{}
}

// newMock returns a status Mock
func newMock() status.Mock {
	return &statusMock{}
}

// MockModule defines the fx options for the mock component.
func MockModule() fxutil.Module {
	return fxutil.Component(
		fx.Provide(newMock),
	)
}

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package multicluster

import (
	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/multicluster/internal/controllers"
	"github.com/hashicorp/consul/internal/multicluster/internal/types"
	"github.com/hashicorp/consul/internal/resource"
)

var (
	// API Group Information
	APIGroup       = types.GroupName
	VersionV2Beta1 = types.VersionV2Beta1
	CurrentVersion = types.CurrentVersion
)

// RegisterTypes adds all resource types within the "multicluster" API group
// to the given type registry
func RegisterTypes(r resource.Registry) {
	types.Register(r)
}

type ControllerDependencies = controllers.Dependencies

// Registercontrollers registers controllers for the multicluster types with the
// given controller manager.
func RegisterControllers(mgr *controller.Manager, deps ControllerDependencies) {
	controllers.Register(mgr, deps)
}

/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package object

import (
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/vim25"
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/vim25/types"
)

type DistributedVirtualSwitch struct {
	Common
}

func NewDistributedVirtualSwitch(c *vim25.Client, ref types.ManagedObjectReference) *DistributedVirtualSwitch {
	return &DistributedVirtualSwitch{
		Common: NewCommon(c, ref),
	}
}

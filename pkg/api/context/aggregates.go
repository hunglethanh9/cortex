/*
Copyright 2019 Cortex Labs, Inc.

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

package context

import (
	"github.com/cortexlabs/cortex/pkg/api/userconfig"
	"github.com/cortexlabs/cortex/pkg/utils/cast"
)

type Aggregates map[string]*Aggregate

type Aggregate struct {
	*userconfig.Aggregate
	*ComputedResourceFields
	Type interface{} `json:"type"`
	Key  string      `json:"key"`
}

func (aggregate *Aggregate) GetType() interface{} {
	return aggregate.Type
}

// Returns map[string]string because after autogen, arg values are constant or aggregate names
func (aggregate *Aggregate) Args() map[string]string {
	args, _ := cast.InterfaceToStrStrMap(aggregate.Inputs.Args)
	return args
}

func (aggregates Aggregates) OneByID(id string) *Aggregate {
	for _, aggregate := range aggregates {
		if aggregate.ID == id {
			return aggregate
		}
	}
	return nil
}

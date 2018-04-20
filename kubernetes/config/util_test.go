/*
Copyright 2018 The Kubernetes Authors.

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

package config

import (
	"reflect"
	"testing"

	"k8s.io/go/kubernetes/config/api"
)

func TestSetUserWithName(t *testing.T) {
	tcs := []struct {
		Origin   []api.NamedAuthInfo
		Name     string
		User     *api.AuthInfo
		Expected []api.NamedAuthInfo
	}{
		{
			Origin: []api.NamedAuthInfo{
				{"A", api.AuthInfo{}},
				{"B", api.AuthInfo{}},
				{"C", api.AuthInfo{}},
			},
			Name: "B",
			User: &api.AuthInfo{Token: "test-token"},
			Expected: []api.NamedAuthInfo{
				{"A", api.AuthInfo{}},
				{"B", api.AuthInfo{Token: "test-token"}},
				{"C", api.AuthInfo{}},
			},
		},
	}

	for _, tc := range tcs {
		if err := setUserWithName(tc.Origin, tc.Name, tc.User); err != nil {
			t.Errorf("unexpected error setting user with name: %v", err)
		}

		if !reflect.DeepEqual(tc.Origin, tc.Expected) {
			t.Errorf("setUserWithName mismatch: want %v, got %v", tc.Expected, tc.Origin)
		}
	}
}

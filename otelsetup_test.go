// Copyright 2023-2025 The sacloud/go-otelsetup Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package otelsetup

import (
	"testing"
)

// Test_newResource semconvとotel SDKの不整合を検知するために正常系のみテストする
//
// Note: 不整合があると実行時にcannot merge resource due to conflicting Schema URLというエラーが出る
func Test_newResource(t *testing.T) {
	_, err := newResource(Options{
		ServiceName:    "test",
		ServiceVersion: "v0.0.1",
	})
	if err != nil {
		t.Fatal(err)
	}
}

/*
 * Copyright 2020-2021 by Nedim Sabic Sabic
 * https://www.fibratus.io
 * All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ql

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFunction(t *testing.T) {
	var tests = []struct {
		expr string
		err  error
	}{
		{expr: "cidr_contains(net.dip)", err: errors.New("CIDR_CONTAINS function requires 2 argument(s) but 1 argument(s) given")},
		{expr: "cidr_contains(net.dip, 12)", err: errors.New("argument #2 (cidr) in function CIDR_CONTAINS should be one of: string")},
		{expr: "cidr_contains(net.dip, '172.17.12.4/24')"},
		{expr: "md('172.17.12.4')", err: errors.New("md function is undefined. Did you mean one of CIDR_CONTAINS|MD5?")},
	}

	for i, tt := range tests {
		p := NewParser(tt.expr)
		_, err := p.ParseExpr()
		if err == nil && tt.err != nil {
			t.Errorf("%d. exp=%s expected error=%v", i, tt.expr, tt.err)
		} else if err != nil {
			assert.EqualError(t, err, tt.err.Error())
		} else if err != nil && tt.err == nil {
			t.Errorf("%d. exp=%s got error=%v", i, tt.expr, err)
		}
	}
}

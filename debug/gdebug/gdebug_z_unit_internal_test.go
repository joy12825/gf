// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/joy12825/gf.

package gdebug

import (
	"testing"
)

func Test_getPackageFromCallerFunction(t *testing.T) {
	dataMap := map[string]string{
		"github.com/joy12825/gf/v2/test/a":       "github.com/joy12825/gf/v2/test/a",
		"github.com/joy12825/gf/v2/test/a.C":     "github.com/joy12825/gf/v2/test/a",
		"github.com/joy12825/gf/v2/test/aa.C":    "github.com/joy12825/gf/v2/test/aa",
		"github.com/joy12825/gf/v2/test/gtest.C": "github.com/joy12825/gf/v2/test/gtest",
	}
	for functionName, packageName := range dataMap {
		if result := getPackageFromCallerFunction(functionName); result != packageName {
			t.Logf(`%s != %s`, result, packageName)
			t.Fail()
		}
	}
}

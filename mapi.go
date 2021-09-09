/*
   Copyright NetFoundry, Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   https://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package cf

import (
	"fmt"
)

func MapIToMapS(in map[interface{}]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range in {
		result[fmt.Sprintf("%v", k)] = CleanUpMapValue(v)
	}
	return result
}

func CleanUpInterfaceArray(in []interface{}) []interface{} {
	result := make([]interface{}, len(in))
	for i, v := range in {
		result[i] = CleanUpMapValue(v)
	}
	return result
}

func CleanUpMapValue(v interface{}) interface{} {
	switch v := v.(type) {
	case []interface{}:
		return CleanUpInterfaceArray(v)

	case map[interface{}]interface{}:
		return MapIToMapS(v)

	default:
		return v
	}
}

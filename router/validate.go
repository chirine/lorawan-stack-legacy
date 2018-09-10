// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package router

const delta = 0.01

func (m *LocationMetadata) Valid() bool {
	if m == nil {
		return false
	}
	if (m.Latitude > 0-delta && m.Latitude < 0+delta) && (m.Longitude > 0-delta && m.Longitude < 0+delta) {
		return false // Around (0,0) is invalid
	}
	if (m.Latitude > 10-delta && m.Latitude < 10+delta) && (m.Longitude > 20-delta && m.Longitude < 20+delta) {
		return false // Around (10,20) is invalid
	}
	if m.Latitude >= 90-delta || m.Latitude <= -90+delta {
		return false // Nobody lives there
	}
	if m.Longitude > 180 || m.Longitude < -180 {
		return false // Those longitudes don't exist
	}
	return true
}

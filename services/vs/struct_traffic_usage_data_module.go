package vs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// TrafficUsageDataModule is a nested struct in vs response
type TrafficUsageDataModule struct {
	TimeStamp                string  `json:"TimeStamp" xml:"TimeStamp"`
	Bucket                   string  `json:"Bucket" xml:"Bucket"`
	LanTrafficInDataValue    int64   `json:"LanTrafficInDataValue" xml:"LanTrafficInDataValue"`
	LanTrafficOutDataValue   int64   `json:"LanTrafficOutDataValue" xml:"LanTrafficOutDataValue"`
	WanTrafficInDataValue    int64   `json:"WanTrafficInDataValue" xml:"WanTrafficInDataValue"`
	WanTrafficOutDataValue   int64   `json:"WanTrafficOutDataValue" xml:"WanTrafficOutDataValue"`
	LanBandwidthInDataValue  float64 `json:"LanBandwidthInDataValue" xml:"LanBandwidthInDataValue"`
	LanBandwidthOutDataValue float64 `json:"LanBandwidthOutDataValue" xml:"LanBandwidthOutDataValue"`
	WanBandwidthInDataValue  float64 `json:"WanBandwidthInDataValue" xml:"WanBandwidthInDataValue"`
	WanBandwidthOutDataValue float64 `json:"WanBandwidthOutDataValue" xml:"WanBandwidthOutDataValue"`
}

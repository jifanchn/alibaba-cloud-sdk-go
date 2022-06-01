package amptest

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// HuichengTestGraySix invokes the amptest.HuichengTestGraySix API synchronously
func (client *Client) HuichengTestGraySix(request *HuichengTestGraySixRequest) (response *HuichengTestGraySixResponse, err error) {
	response = CreateHuichengTestGraySixResponse()
	err = client.DoAction(request, response)
	return
}

// HuichengTestGraySixWithChan invokes the amptest.HuichengTestGraySix API asynchronously
func (client *Client) HuichengTestGraySixWithChan(request *HuichengTestGraySixRequest) (<-chan *HuichengTestGraySixResponse, <-chan error) {
	responseChan := make(chan *HuichengTestGraySixResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.HuichengTestGraySix(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// HuichengTestGraySixWithCallback invokes the amptest.HuichengTestGraySix API asynchronously
func (client *Client) HuichengTestGraySixWithCallback(request *HuichengTestGraySixRequest, callback func(response *HuichengTestGraySixResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *HuichengTestGraySixResponse
		var err error
		defer close(result)
		response, err = client.HuichengTestGraySix(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// HuichengTestGraySixRequest is the request struct for api HuichengTestGraySix
type HuichengTestGraySixRequest struct {
	*requests.RpcRequest
	Home HuichengTestGraySixHome `position:"Query" name:"Home"  type:"Struct"`
}

// HuichengTestGraySixHome is a repeated param struct in HuichengTestGraySixRequest
type HuichengTestGraySixHome struct {
	Address      HuichengTestGraySixHomeAddress             `name:"Address" type:"Struct"`
	T            HuichengTestGraySixHomeT                   `name:"T" type:"Struct"`
	PhoneNumbers *[]string                                  `name:"PhoneNumbers" type:"Repeated"`
	DMap         map[string]HuichengTestGraySixHomeDMapItem `name:"DMap" type:"Map"`
	NameToAge    map[string]string                          `name:"NameToAge" type:"Map"`
	Locations    *[]HuichengTestGraySixHomeLocationsItem    `name:"Locations" type:"Repeated"`
}

// HuichengTestGraySixHomeAddress is a repeated param struct in HuichengTestGraySixRequest
type HuichengTestGraySixHomeAddress struct {
	T        HuichengTestGraySixHomeAddressT        `name:"T" type:"Struct"`
	Location HuichengTestGraySixHomeAddressLocation `name:"Location" type:"Struct"`
	Detail   string                                 `name:"Detail"`
}

// HuichengTestGraySixHomeT is a repeated param struct in HuichengTestGraySixRequest
type HuichengTestGraySixHomeT struct {
	Class string `name:"Class"`
}

// HuichengTestGraySixHomeDMapItem is a repeated param struct in HuichengTestGraySixRequest
type HuichengTestGraySixHomeDMapItem struct {
	Location HuichengTestGraySixHomeDMapItemLocation `name:"Location" type:"Struct"`
	Detail   string                                  `name:"Detail"`
}

// HuichengTestGraySixHomeLocationsItem is a repeated param struct in HuichengTestGraySixRequest
type HuichengTestGraySixHomeLocationsItem struct {
	Late string `name:"Late"`
	Lon  string `name:"Lon"`
}

// HuichengTestGraySixHomeAddressT is a repeated param struct in HuichengTestGraySixRequest
type HuichengTestGraySixHomeAddressT struct {
	Class string `name:"Class"`
}

// HuichengTestGraySixHomeAddressLocation is a repeated param struct in HuichengTestGraySixRequest
type HuichengTestGraySixHomeAddressLocation struct {
	Late string `name:"Late"`
	Lon  string `name:"Lon"`
}

// HuichengTestGraySixHomeDMapItemLocation is a repeated param struct in HuichengTestGraySixRequest
type HuichengTestGraySixHomeDMapItemLocation struct {
	Late string `name:"Late"`
	Lon  string `name:"Lon"`
}

// HuichengTestGraySixResponse is the response struct for api HuichengTestGraySix
type HuichengTestGraySixResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Size      int    `json:"Size" xml:"Size"`
	Value     string `json:"Value" xml:"Value"`
}

// CreateHuichengTestGraySixRequest creates a request to invoke HuichengTestGraySix API
func CreateHuichengTestGraySixRequest() (request *HuichengTestGraySixRequest) {
	request = &HuichengTestGraySixRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("AmpTest", "2020-12-30", "HuichengTestGraySix", "AmpTest", "openAPI")
	request.Method = requests.POST
	return
}

// CreateHuichengTestGraySixResponse creates a response to parse from HuichengTestGraySix response
func CreateHuichengTestGraySixResponse() (response *HuichengTestGraySixResponse) {
	response = &HuichengTestGraySixResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

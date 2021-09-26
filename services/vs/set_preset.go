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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SetPreset invokes the vs.SetPreset API synchronously
func (client *Client) SetPreset(request *SetPresetRequest) (response *SetPresetResponse, err error) {
	response = CreateSetPresetResponse()
	err = client.DoAction(request, response)
	return
}

// SetPresetWithChan invokes the vs.SetPreset API asynchronously
func (client *Client) SetPresetWithChan(request *SetPresetRequest) (<-chan *SetPresetResponse, <-chan error) {
	responseChan := make(chan *SetPresetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetPreset(request)
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

// SetPresetWithCallback invokes the vs.SetPreset API asynchronously
func (client *Client) SetPresetWithCallback(request *SetPresetRequest, callback func(response *SetPresetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetPresetResponse
		var err error
		defer close(result)
		response, err = client.SetPreset(request)
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

// SetPresetRequest is the request struct for api SetPreset
type SetPresetRequest struct {
	*requests.RpcRequest
	SubProtocol string           `position:"Query" name:"SubProtocol"`
	Id          string           `position:"Query" name:"Id"`
	PresetId    string           `position:"Query" name:"PresetId"`
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}

// SetPresetResponse is the response struct for api SetPreset
type SetPresetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateSetPresetRequest creates a request to invoke SetPreset API
func CreateSetPresetRequest() (request *SetPresetRequest) {
	request = &SetPresetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "SetPreset", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetPresetResponse creates a response to parse from SetPreset response
func CreateSetPresetResponse() (response *SetPresetResponse) {
	response = &SetPresetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

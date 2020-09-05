package vcs

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

// DescribeDevices invokes the vcs.DescribeDevices API synchronously
// api document: https://help.aliyun.com/api/vcs/describedevices.html
func (client *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
	response = CreateDescribeDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDevicesWithChan invokes the vcs.DescribeDevices API asynchronously
// api document: https://help.aliyun.com/api/vcs/describedevices.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDevicesWithChan(request *DescribeDevicesRequest) (<-chan *DescribeDevicesResponse, <-chan error) {
	responseChan := make(chan *DescribeDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDevices(request)
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

// DescribeDevicesWithCallback invokes the vcs.DescribeDevices API asynchronously
// api document: https://help.aliyun.com/api/vcs/describedevices.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDevicesWithCallback(request *DescribeDevicesRequest, callback func(response *DescribeDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDevicesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDevices(request)
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

// DescribeDevicesRequest is the request struct for api DescribeDevices
type DescribeDevicesRequest struct {
	*requests.RpcRequest
	PageNum    requests.Integer `position:"Body" name:"PageNum"`
	CorpIdList string           `position:"Body" name:"CorpIdList"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
}

// DescribeDevicesResponse is the response struct for api DescribeDevices
type DescribeDevicesResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeDevicesRequest creates a request to invoke DescribeDevices API
func CreateDescribeDevicesRequest() (request *DescribeDevicesRequest) {
	request = &DescribeDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "DescribeDevices", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDevicesResponse creates a response to parse from DescribeDevices response
func CreateDescribeDevicesResponse() (response *DescribeDevicesResponse) {
	response = &DescribeDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

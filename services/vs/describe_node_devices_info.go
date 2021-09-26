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

// DescribeNodeDevicesInfo invokes the vs.DescribeNodeDevicesInfo API synchronously
func (client *Client) DescribeNodeDevicesInfo(request *DescribeNodeDevicesInfoRequest) (response *DescribeNodeDevicesInfoResponse, err error) {
	response = CreateDescribeNodeDevicesInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNodeDevicesInfoWithChan invokes the vs.DescribeNodeDevicesInfo API asynchronously
func (client *Client) DescribeNodeDevicesInfoWithChan(request *DescribeNodeDevicesInfoRequest) (<-chan *DescribeNodeDevicesInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeNodeDevicesInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNodeDevicesInfo(request)
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

// DescribeNodeDevicesInfoWithCallback invokes the vs.DescribeNodeDevicesInfo API asynchronously
func (client *Client) DescribeNodeDevicesInfoWithCallback(request *DescribeNodeDevicesInfoRequest, callback func(response *DescribeNodeDevicesInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNodeDevicesInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeNodeDevicesInfo(request)
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

// DescribeNodeDevicesInfoRequest is the request struct for api DescribeNodeDevicesInfo
type DescribeNodeDevicesInfoRequest struct {
	*requests.RpcRequest
	ShowLog  string           `position:"Query" name:"ShowLog"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	NodeName string           `position:"Query" name:"NodeName"`
}

// DescribeNodeDevicesInfoResponse is the response struct for api DescribeNodeDevicesInfo
type DescribeNodeDevicesInfoResponse struct {
	*responses.BaseResponse
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	NodeDevices []NodeDevice `json:"NodeDevices" xml:"NodeDevices"`
}

// CreateDescribeNodeDevicesInfoRequest creates a request to invoke DescribeNodeDevicesInfo API
func CreateDescribeNodeDevicesInfoRequest() (request *DescribeNodeDevicesInfoRequest) {
	request = &DescribeNodeDevicesInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeNodeDevicesInfo", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeNodeDevicesInfoResponse creates a response to parse from DescribeNodeDevicesInfo response
func CreateDescribeNodeDevicesInfoResponse() (response *DescribeNodeDevicesInfoResponse) {
	response = &DescribeNodeDevicesInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

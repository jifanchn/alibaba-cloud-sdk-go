package rds

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

// DescribeActionEventPolicy invokes the rds.DescribeActionEventPolicy API synchronously
func (client *Client) DescribeActionEventPolicy(request *DescribeActionEventPolicyRequest) (response *DescribeActionEventPolicyResponse, err error) {
	response = CreateDescribeActionEventPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeActionEventPolicyWithChan invokes the rds.DescribeActionEventPolicy API asynchronously
func (client *Client) DescribeActionEventPolicyWithChan(request *DescribeActionEventPolicyRequest) (<-chan *DescribeActionEventPolicyResponse, <-chan error) {
	responseChan := make(chan *DescribeActionEventPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeActionEventPolicy(request)
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

// DescribeActionEventPolicyWithCallback invokes the rds.DescribeActionEventPolicy API asynchronously
func (client *Client) DescribeActionEventPolicyWithCallback(request *DescribeActionEventPolicyRequest, callback func(response *DescribeActionEventPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeActionEventPolicyResponse
		var err error
		defer close(result)
		response, err = client.DescribeActionEventPolicy(request)
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

// DescribeActionEventPolicyRequest is the request struct for api DescribeActionEventPolicy
type DescribeActionEventPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeActionEventPolicyResponse is the response struct for api DescribeActionEventPolicy
type DescribeActionEventPolicyResponse struct {
	*responses.BaseResponse
	EnableEventLog string `json:"EnableEventLog" xml:"EnableEventLog"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	RegionId       string `json:"RegionId" xml:"RegionId"`
}

// CreateDescribeActionEventPolicyRequest creates a request to invoke DescribeActionEventPolicy API
func CreateDescribeActionEventPolicyRequest() (request *DescribeActionEventPolicyRequest) {
	request = &DescribeActionEventPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeActionEventPolicy", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeActionEventPolicyResponse creates a response to parse from DescribeActionEventPolicy response
func CreateDescribeActionEventPolicyResponse() (response *DescribeActionEventPolicyResponse) {
	response = &DescribeActionEventPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

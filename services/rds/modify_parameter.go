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

// ModifyParameter invokes the rds.ModifyParameter API synchronously
func (client *Client) ModifyParameter(request *ModifyParameterRequest) (response *ModifyParameterResponse, err error) {
	response = CreateModifyParameterResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyParameterWithChan invokes the rds.ModifyParameter API asynchronously
func (client *Client) ModifyParameterWithChan(request *ModifyParameterRequest) (<-chan *ModifyParameterResponse, <-chan error) {
	responseChan := make(chan *ModifyParameterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyParameter(request)
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

// ModifyParameterWithCallback invokes the rds.ModifyParameter API asynchronously
func (client *Client) ModifyParameterWithCallback(request *ModifyParameterRequest, callback func(response *ModifyParameterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyParameterResponse
		var err error
		defer close(result)
		response, err = client.ModifyParameter(request)
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

// ModifyParameterRequest is the request struct for api ModifyParameter
type ModifyParameterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	Parameters           string           `position:"Query" name:"Parameters"`
}

// ModifyParameterResponse is the response struct for api ModifyParameter
type ModifyParameterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyParameterRequest creates a request to invoke ModifyParameter API
func CreateModifyParameterRequest() (request *ModifyParameterRequest) {
	request = &ModifyParameterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2013-05-28", "ModifyParameter", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyParameterResponse creates a response to parse from ModifyParameter response
func CreateModifyParameterResponse() (response *ModifyParameterResponse) {
	response = &ModifyParameterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

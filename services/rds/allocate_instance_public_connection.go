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

// AllocateInstancePublicConnection invokes the rds.AllocateInstancePublicConnection API synchronously
func (client *Client) AllocateInstancePublicConnection(request *AllocateInstancePublicConnectionRequest) (response *AllocateInstancePublicConnectionResponse, err error) {
	response = CreateAllocateInstancePublicConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// AllocateInstancePublicConnectionWithChan invokes the rds.AllocateInstancePublicConnection API asynchronously
func (client *Client) AllocateInstancePublicConnectionWithChan(request *AllocateInstancePublicConnectionRequest) (<-chan *AllocateInstancePublicConnectionResponse, <-chan error) {
	responseChan := make(chan *AllocateInstancePublicConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocateInstancePublicConnection(request)
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

// AllocateInstancePublicConnectionWithCallback invokes the rds.AllocateInstancePublicConnection API asynchronously
func (client *Client) AllocateInstancePublicConnectionWithCallback(request *AllocateInstancePublicConnectionRequest, callback func(response *AllocateInstancePublicConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocateInstancePublicConnectionResponse
		var err error
		defer close(result)
		response, err = client.AllocateInstancePublicConnection(request)
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

// AllocateInstancePublicConnectionRequest is the request struct for api AllocateInstancePublicConnection
type AllocateInstancePublicConnectionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string           `position:"Query" name:"ConnectionStringPrefix"`
	GeneralGroupName       string           `position:"Query" name:"GeneralGroupName"`
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	Port                   string           `position:"Query" name:"Port"`
}

// AllocateInstancePublicConnectionResponse is the response struct for api AllocateInstancePublicConnection
type AllocateInstancePublicConnectionResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	ConnectionString string `json:"ConnectionString" xml:"ConnectionString"`
	DbInstanceName   string `json:"DbInstanceName" xml:"DbInstanceName"`
}

// CreateAllocateInstancePublicConnectionRequest creates a request to invoke AllocateInstancePublicConnection API
func CreateAllocateInstancePublicConnectionRequest() (request *AllocateInstancePublicConnectionRequest) {
	request = &AllocateInstancePublicConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "AllocateInstancePublicConnection", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAllocateInstancePublicConnectionResponse creates a response to parse from AllocateInstancePublicConnection response
func CreateAllocateInstancePublicConnectionResponse() (response *AllocateInstancePublicConnectionResponse) {
	response = &AllocateInstancePublicConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

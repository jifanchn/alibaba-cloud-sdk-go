package retailcloud

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

// ListServices invokes the retailcloud.ListServices API synchronously
func (client *Client) ListServices(request *ListServicesRequest) (response *ListServicesResponse, err error) {
	response = CreateListServicesResponse()
	err = client.DoAction(request, response)
	return
}

// ListServicesWithChan invokes the retailcloud.ListServices API asynchronously
func (client *Client) ListServicesWithChan(request *ListServicesRequest) (<-chan *ListServicesResponse, <-chan error) {
	responseChan := make(chan *ListServicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListServices(request)
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

// ListServicesWithCallback invokes the retailcloud.ListServices API asynchronously
func (client *Client) ListServicesWithCallback(request *ListServicesRequest, callback func(response *ListServicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListServicesResponse
		var err error
		defer close(result)
		response, err = client.ListServices(request)
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

// ListServicesRequest is the request struct for api ListServices
type ListServicesRequest struct {
	*requests.RpcRequest
	ServiceType string           `position:"Query" name:"ServiceType"`
	AppId       requests.Integer `position:"Query" name:"AppId"`
	Name        string           `position:"Query" name:"Name"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	EnvId       requests.Integer `position:"Query" name:"EnvId"`
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
}

// ListServicesResponse is the response struct for api ListServices
type ListServicesResponse struct {
	*responses.BaseResponse
	Code       int               `json:"Code" xml:"Code"`
	ErrorMsg   string            `json:"ErrorMsg" xml:"ErrorMsg"`
	PageNumber int               `json:"PageNumber" xml:"PageNumber"`
	PageSize   int               `json:"PageSize" xml:"PageSize"`
	RequestId  string            `json:"RequestId" xml:"RequestId"`
	TotalCount int64             `json:"TotalCount" xml:"TotalCount"`
	ClusterIP  string            `json:"ClusterIP" xml:"ClusterIP"`
	Data       []ServiceInstance `json:"Data" xml:"Data"`
}

// CreateListServicesRequest creates a request to invoke ListServices API
func CreateListServicesRequest() (request *ListServicesRequest) {
	request = &ListServicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ListServices", "retailcloud", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListServicesResponse creates a response to parse from ListServices response
func CreateListServicesResponse() (response *ListServicesResponse) {
	response = &ListServicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

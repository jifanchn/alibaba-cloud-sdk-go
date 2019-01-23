package cloudapi

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

// DescribePurchasedApiGroups invokes the cloudapi.DescribePurchasedApiGroups API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describepurchasedapigroups.html
func (client *Client) DescribePurchasedApiGroups(request *DescribePurchasedApiGroupsRequest) (response *DescribePurchasedApiGroupsResponse, err error) {
	response = CreateDescribePurchasedApiGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePurchasedApiGroupsWithChan invokes the cloudapi.DescribePurchasedApiGroups API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describepurchasedapigroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePurchasedApiGroupsWithChan(request *DescribePurchasedApiGroupsRequest) (<-chan *DescribePurchasedApiGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribePurchasedApiGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePurchasedApiGroups(request)
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

// DescribePurchasedApiGroupsWithCallback invokes the cloudapi.DescribePurchasedApiGroups API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describepurchasedapigroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePurchasedApiGroupsWithCallback(request *DescribePurchasedApiGroupsRequest, callback func(response *DescribePurchasedApiGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePurchasedApiGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribePurchasedApiGroups(request)
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

// DescribePurchasedApiGroupsRequest is the request struct for api DescribePurchasedApiGroups
type DescribePurchasedApiGroupsRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribePurchasedApiGroupsResponse is the response struct for api DescribePurchasedApiGroups
type DescribePurchasedApiGroupsResponse struct {
	*responses.BaseResponse
	RequestId                   string                      `json:"RequestId" xml:"RequestId"`
	TotalCount                  int                         `json:"TotalCount" xml:"TotalCount"`
	PageSize                    int                         `json:"PageSize" xml:"PageSize"`
	PageNumber                  int                         `json:"PageNumber" xml:"PageNumber"`
	PurchasedApiGroupAttributes PurchasedApiGroupAttributes `json:"PurchasedApiGroupAttributes" xml:"PurchasedApiGroupAttributes"`
}

// CreateDescribePurchasedApiGroupsRequest creates a request to invoke DescribePurchasedApiGroups API
func CreateDescribePurchasedApiGroupsRequest() (request *DescribePurchasedApiGroupsRequest) {
	request = &DescribePurchasedApiGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribePurchasedApiGroups", "", "")
	return
}

// CreateDescribePurchasedApiGroupsResponse creates a response to parse from DescribePurchasedApiGroups response
func CreateDescribePurchasedApiGroupsResponse() (response *DescribePurchasedApiGroupsResponse) {
	response = &DescribePurchasedApiGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

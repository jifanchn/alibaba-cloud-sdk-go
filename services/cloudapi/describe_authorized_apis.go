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

// DescribeAuthorizedApis invokes the cloudapi.DescribeAuthorizedApis API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeauthorizedapis.html
func (client *Client) DescribeAuthorizedApis(request *DescribeAuthorizedApisRequest) (response *DescribeAuthorizedApisResponse, err error) {
	response = CreateDescribeAuthorizedApisResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAuthorizedApisWithChan invokes the cloudapi.DescribeAuthorizedApis API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeauthorizedapis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAuthorizedApisWithChan(request *DescribeAuthorizedApisRequest) (<-chan *DescribeAuthorizedApisResponse, <-chan error) {
	responseChan := make(chan *DescribeAuthorizedApisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAuthorizedApis(request)
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

// DescribeAuthorizedApisWithCallback invokes the cloudapi.DescribeAuthorizedApis API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeauthorizedapis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAuthorizedApisWithCallback(request *DescribeAuthorizedApisRequest, callback func(response *DescribeAuthorizedApisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAuthorizedApisResponse
		var err error
		defer close(result)
		response, err = client.DescribeAuthorizedApis(request)
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

// DescribeAuthorizedApisRequest is the request struct for api DescribeAuthorizedApis
type DescribeAuthorizedApisRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	AppId         requests.Integer `position:"Query" name:"AppId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeAuthorizedApisResponse is the response struct for api DescribeAuthorizedApis
type DescribeAuthorizedApisResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	TotalCount     int            `json:"TotalCount" xml:"TotalCount"`
	PageSize       int            `json:"PageSize" xml:"PageSize"`
	PageNumber     int            `json:"PageNumber" xml:"PageNumber"`
	AuthorizedApis AuthorizedApis `json:"AuthorizedApis" xml:"AuthorizedApis"`
}

// CreateDescribeAuthorizedApisRequest creates a request to invoke DescribeAuthorizedApis API
func CreateDescribeAuthorizedApisRequest() (request *DescribeAuthorizedApisRequest) {
	request = &DescribeAuthorizedApisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeAuthorizedApis", "", "")
	return
}

// CreateDescribeAuthorizedApisResponse creates a response to parse from DescribeAuthorizedApis response
func CreateDescribeAuthorizedApisResponse() (response *DescribeAuthorizedApisResponse) {
	response = &DescribeAuthorizedApisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

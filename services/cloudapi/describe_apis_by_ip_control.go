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

// DescribeApisByIpControl invokes the cloudapi.DescribeApisByIpControl API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapisbyipcontrol.html
func (client *Client) DescribeApisByIpControl(request *DescribeApisByIpControlRequest) (response *DescribeApisByIpControlResponse, err error) {
	response = CreateDescribeApisByIpControlResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApisByIpControlWithChan invokes the cloudapi.DescribeApisByIpControl API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapisbyipcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApisByIpControlWithChan(request *DescribeApisByIpControlRequest) (<-chan *DescribeApisByIpControlResponse, <-chan error) {
	responseChan := make(chan *DescribeApisByIpControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApisByIpControl(request)
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

// DescribeApisByIpControlWithCallback invokes the cloudapi.DescribeApisByIpControl API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapisbyipcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApisByIpControlWithCallback(request *DescribeApisByIpControlRequest, callback func(response *DescribeApisByIpControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApisByIpControlResponse
		var err error
		defer close(result)
		response, err = client.DescribeApisByIpControl(request)
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

// DescribeApisByIpControlRequest is the request struct for api DescribeApisByIpControl
type DescribeApisByIpControlRequest struct {
	*requests.RpcRequest
	IpControlId   string           `position:"Query" name:"IpControlId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeApisByIpControlResponse is the response struct for api DescribeApisByIpControl
type DescribeApisByIpControlResponse struct {
	*responses.BaseResponse
	RequestId  string                            `json:"RequestId" xml:"RequestId"`
	TotalCount int                               `json:"TotalCount" xml:"TotalCount"`
	PageSize   int                               `json:"PageSize" xml:"PageSize"`
	PageNumber int                               `json:"PageNumber" xml:"PageNumber"`
	ApiInfos   ApiInfosInDescribeApisByIpControl `json:"ApiInfos" xml:"ApiInfos"`
}

// CreateDescribeApisByIpControlRequest creates a request to invoke DescribeApisByIpControl API
func CreateDescribeApisByIpControlRequest() (request *DescribeApisByIpControlRequest) {
	request = &DescribeApisByIpControlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApisByIpControl", "", "")
	return
}

// CreateDescribeApisByIpControlResponse creates a response to parse from DescribeApisByIpControl response
func CreateDescribeApisByIpControlResponse() (response *DescribeApisByIpControlResponse) {
	response = &DescribeApisByIpControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package ens

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

// DescribeDataDistResult invokes the ens.DescribeDataDistResult API synchronously
// api document: https://help.aliyun.com/api/ens/describedatadistresult.html
func (client *Client) DescribeDataDistResult(request *DescribeDataDistResultRequest) (response *DescribeDataDistResultResponse, err error) {
	response = CreateDescribeDataDistResultResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataDistResultWithChan invokes the ens.DescribeDataDistResult API asynchronously
// api document: https://help.aliyun.com/api/ens/describedatadistresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDataDistResultWithChan(request *DescribeDataDistResultRequest) (<-chan *DescribeDataDistResultResponse, <-chan error) {
	responseChan := make(chan *DescribeDataDistResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataDistResult(request)
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

// DescribeDataDistResultWithCallback invokes the ens.DescribeDataDistResult API asynchronously
// api document: https://help.aliyun.com/api/ens/describedatadistresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDataDistResultWithCallback(request *DescribeDataDistResultRequest, callback func(response *DescribeDataDistResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataDistResultResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataDistResult(request)
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

// DescribeDataDistResultRequest is the request struct for api DescribeDataDistResult
type DescribeDataDistResultRequest struct {
	*requests.RpcRequest
	MaxDate      string           `position:"Query" name:"MaxDate"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	MinDate      string           `position:"Query" name:"MinDate"`
	DataVersions string           `position:"Query" name:"DataVersions"`
	InstanceIds  string           `position:"Query" name:"InstanceIds"`
	AppId        string           `position:"Query" name:"AppId"`
	DataNames    string           `position:"Query" name:"DataNames"`
}

// DescribeDataDistResultResponse is the response struct for api DescribeDataDistResult
type DescribeDataDistResultResponse struct {
	*responses.BaseResponse
	RequestId   string                              `json:"RequestId" xml:"RequestId"`
	TotalCount  int                                 `json:"TotalCount" xml:"TotalCount"`
	PageNumber  int                                 `json:"PageNumber" xml:"PageNumber"`
	PageSize    int                                 `json:"PageSize" xml:"PageSize"`
	DistResults DistResultsInDescribeDataDistResult `json:"DistResults" xml:"DistResults"`
}

// CreateDescribeDataDistResultRequest creates a request to invoke DescribeDataDistResult API
func CreateDescribeDataDistResultRequest() (request *DescribeDataDistResultRequest) {
	request = &DescribeDataDistResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeDataDistResult", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataDistResultResponse creates a response to parse from DescribeDataDistResult response
func CreateDescribeDataDistResultResponse() (response *DescribeDataDistResultResponse) {
	response = &DescribeDataDistResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
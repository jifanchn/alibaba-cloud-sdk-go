package ddoscoo

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

// ListAsyncTask invokes the ddoscoo.ListAsyncTask API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/listasynctask.html
func (client *Client) ListAsyncTask(request *ListAsyncTaskRequest) (response *ListAsyncTaskResponse, err error) {
	response = CreateListAsyncTaskResponse()
	err = client.DoAction(request, response)
	return
}

// ListAsyncTaskWithChan invokes the ddoscoo.ListAsyncTask API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/listasynctask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAsyncTaskWithChan(request *ListAsyncTaskRequest) (<-chan *ListAsyncTaskResponse, <-chan error) {
	responseChan := make(chan *ListAsyncTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAsyncTask(request)
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

// ListAsyncTaskWithCallback invokes the ddoscoo.ListAsyncTask API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/listasynctask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAsyncTaskWithCallback(request *ListAsyncTaskRequest, callback func(response *ListAsyncTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAsyncTaskResponse
		var err error
		defer close(result)
		response, err = client.ListAsyncTask(request)
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

// ListAsyncTaskRequest is the request struct for api ListAsyncTask
type ListAsyncTaskRequest struct {
	*requests.RpcRequest
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageNo          requests.Integer `position:"Query" name:"PageNo"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Lang            string           `position:"Query" name:"Lang"`
}

// ListAsyncTaskResponse is the response struct for api ListAsyncTask
type ListAsyncTaskResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	Total      int         `json:"Total" xml:"Total"`
	AsyncTasks []AsyncTask `json:"AsyncTasks" xml:"AsyncTasks"`
}

// CreateListAsyncTaskRequest creates a request to invoke ListAsyncTask API
func CreateListAsyncTaskRequest() (request *ListAsyncTaskRequest) {
	request = &ListAsyncTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "ListAsyncTask", "ddoscoo", "openAPI")
	return
}

// CreateListAsyncTaskResponse creates a response to parse from ListAsyncTask response
func CreateListAsyncTaskResponse() (response *ListAsyncTaskResponse) {
	response = &ListAsyncTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
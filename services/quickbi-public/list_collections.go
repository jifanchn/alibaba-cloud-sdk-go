package quickbi_public

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

// ListCollections invokes the quickbi_public.ListCollections API synchronously
func (client *Client) ListCollections(request *ListCollectionsRequest) (response *ListCollectionsResponse, err error) {
	response = CreateListCollectionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCollectionsWithChan invokes the quickbi_public.ListCollections API asynchronously
func (client *Client) ListCollectionsWithChan(request *ListCollectionsRequest) (<-chan *ListCollectionsResponse, <-chan error) {
	responseChan := make(chan *ListCollectionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCollections(request)
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

// ListCollectionsWithCallback invokes the quickbi_public.ListCollections API asynchronously
func (client *Client) ListCollectionsWithCallback(request *ListCollectionsRequest, callback func(response *ListCollectionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCollectionsResponse
		var err error
		defer close(result)
		response, err = client.ListCollections(request)
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

// ListCollectionsRequest is the request struct for api ListCollections
type ListCollectionsRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	UserId      string `position:"Query" name:"UserId"`
}

// ListCollectionsResponse is the response struct for api ListCollections
type ListCollectionsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    []Data `json:"Result" xml:"Result"`
}

// CreateListCollectionsRequest creates a request to invoke ListCollections API
func CreateListCollectionsRequest() (request *ListCollectionsRequest) {
	request = &ListCollectionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2020-08-09", "ListCollections", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListCollectionsResponse creates a response to parse from ListCollections response
func CreateListCollectionsResponse() (response *ListCollectionsResponse) {
	response = &ListCollectionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

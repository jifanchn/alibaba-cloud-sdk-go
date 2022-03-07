package arms

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

// GetAgentDownloadUrl invokes the arms.GetAgentDownloadUrl API synchronously
func (client *Client) GetAgentDownloadUrl(request *GetAgentDownloadUrlRequest) (response *GetAgentDownloadUrlResponse, err error) {
	response = CreateGetAgentDownloadUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetAgentDownloadUrlWithChan invokes the arms.GetAgentDownloadUrl API asynchronously
func (client *Client) GetAgentDownloadUrlWithChan(request *GetAgentDownloadUrlRequest) (<-chan *GetAgentDownloadUrlResponse, <-chan error) {
	responseChan := make(chan *GetAgentDownloadUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAgentDownloadUrl(request)
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

// GetAgentDownloadUrlWithCallback invokes the arms.GetAgentDownloadUrl API asynchronously
func (client *Client) GetAgentDownloadUrlWithCallback(request *GetAgentDownloadUrlRequest, callback func(response *GetAgentDownloadUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAgentDownloadUrlResponse
		var err error
		defer close(result)
		response, err = client.GetAgentDownloadUrl(request)
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

// GetAgentDownloadUrlRequest is the request struct for api GetAgentDownloadUrl
type GetAgentDownloadUrlRequest struct {
	*requests.RpcRequest
}

// GetAgentDownloadUrlResponse is the response struct for api GetAgentDownloadUrl
type GetAgentDownloadUrlResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	ArmsAgentDownloadUrl string `json:"ArmsAgentDownloadUrl" xml:"ArmsAgentDownloadUrl"`
}

// CreateGetAgentDownloadUrlRequest creates a request to invoke GetAgentDownloadUrl API
func CreateGetAgentDownloadUrlRequest() (request *GetAgentDownloadUrlRequest) {
	request = &GetAgentDownloadUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetAgentDownloadUrl", "arms", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetAgentDownloadUrlResponse creates a response to parse from GetAgentDownloadUrl response
func CreateGetAgentDownloadUrlResponse() (response *GetAgentDownloadUrlResponse) {
	response = &GetAgentDownloadUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

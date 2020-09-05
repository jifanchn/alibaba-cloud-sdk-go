package vcs

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

// SearchBody invokes the vcs.SearchBody API synchronously
// api document: https://help.aliyun.com/api/vcs/searchbody.html
func (client *Client) SearchBody(request *SearchBodyRequest) (response *SearchBodyResponse, err error) {
	response = CreateSearchBodyResponse()
	err = client.DoAction(request, response)
	return
}

// SearchBodyWithChan invokes the vcs.SearchBody API asynchronously
// api document: https://help.aliyun.com/api/vcs/searchbody.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchBodyWithChan(request *SearchBodyRequest) (<-chan *SearchBodyResponse, <-chan error) {
	responseChan := make(chan *SearchBodyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchBody(request)
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

// SearchBodyWithCallback invokes the vcs.SearchBody API asynchronously
// api document: https://help.aliyun.com/api/vcs/searchbody.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchBodyWithCallback(request *SearchBodyRequest, callback func(response *SearchBodyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchBodyResponse
		var err error
		defer close(result)
		response, err = client.SearchBody(request)
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

// SearchBodyRequest is the request struct for api SearchBody
type SearchBodyRequest struct {
	*requests.RpcRequest
	CorpId         string           `position:"Body" name:"CorpId"`
	GbId           string           `position:"Body" name:"GbId"`
	StartTimeStamp requests.Integer `position:"Body" name:"StartTimeStamp"`
	EndTimeStamp   requests.Integer `position:"Body" name:"EndTimeStamp"`
	PageNo         requests.Integer `position:"Body" name:"PageNo"`
	PageSize       requests.Integer `position:"Body" name:"PageSize"`
	OptionList     string           `position:"Body" name:"OptionList"`
}

// SearchBodyResponse is the response struct for api SearchBody
type SearchBodyResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSearchBodyRequest creates a request to invoke SearchBody API
func CreateSearchBodyRequest() (request *SearchBodyRequest) {
	request = &SearchBodyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "SearchBody", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSearchBodyResponse creates a response to parse from SearchBody response
func CreateSearchBodyResponse() (response *SearchBodyResponse) {
	response = &SearchBodyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

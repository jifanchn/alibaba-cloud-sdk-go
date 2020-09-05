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

// GetProfileDetail invokes the vcs.GetProfileDetail API synchronously
// api document: https://help.aliyun.com/api/vcs/getprofiledetail.html
func (client *Client) GetProfileDetail(request *GetProfileDetailRequest) (response *GetProfileDetailResponse, err error) {
	response = CreateGetProfileDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetProfileDetailWithChan invokes the vcs.GetProfileDetail API asynchronously
// api document: https://help.aliyun.com/api/vcs/getprofiledetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProfileDetailWithChan(request *GetProfileDetailRequest) (<-chan *GetProfileDetailResponse, <-chan error) {
	responseChan := make(chan *GetProfileDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProfileDetail(request)
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

// GetProfileDetailWithCallback invokes the vcs.GetProfileDetail API asynchronously
// api document: https://help.aliyun.com/api/vcs/getprofiledetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProfileDetailWithCallback(request *GetProfileDetailRequest, callback func(response *GetProfileDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProfileDetailResponse
		var err error
		defer close(result)
		response, err = client.GetProfileDetail(request)
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

// GetProfileDetailRequest is the request struct for api GetProfileDetail
type GetProfileDetailRequest struct {
	*requests.RpcRequest
	IsvSubId  string           `position:"Body" name:"IsvSubId"`
	CorpId    string           `position:"Body" name:"CorpId"`
	ProfileId requests.Integer `position:"Body" name:"ProfileId"`
}

// GetProfileDetailResponse is the response struct for api GetProfileDetail
type GetProfileDetailResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetProfileDetailRequest creates a request to invoke GetProfileDetail API
func CreateGetProfileDetailRequest() (request *GetProfileDetailRequest) {
	request = &GetProfileDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "GetProfileDetail", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetProfileDetailResponse creates a response to parse from GetProfileDetail response
func CreateGetProfileDetailResponse() (response *GetProfileDetailResponse) {
	response = &GetProfileDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

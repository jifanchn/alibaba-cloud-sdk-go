package sas

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

// ModifyLoginSwitchConfig invokes the sas.ModifyLoginSwitchConfig API synchronously
// api document: https://help.aliyun.com/api/sas/modifyloginswitchconfig.html
func (client *Client) ModifyLoginSwitchConfig(request *ModifyLoginSwitchConfigRequest) (response *ModifyLoginSwitchConfigResponse, err error) {
	response = CreateModifyLoginSwitchConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLoginSwitchConfigWithChan invokes the sas.ModifyLoginSwitchConfig API asynchronously
// api document: https://help.aliyun.com/api/sas/modifyloginswitchconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLoginSwitchConfigWithChan(request *ModifyLoginSwitchConfigRequest) (<-chan *ModifyLoginSwitchConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyLoginSwitchConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLoginSwitchConfig(request)
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

// ModifyLoginSwitchConfigWithCallback invokes the sas.ModifyLoginSwitchConfig API asynchronously
// api document: https://help.aliyun.com/api/sas/modifyloginswitchconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLoginSwitchConfigWithCallback(request *ModifyLoginSwitchConfigRequest, callback func(response *ModifyLoginSwitchConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLoginSwitchConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyLoginSwitchConfig(request)
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

// ModifyLoginSwitchConfigRequest is the request struct for api ModifyLoginSwitchConfig
type ModifyLoginSwitchConfigRequest struct {
	*requests.RpcRequest
	Item     string           `position:"Query" name:"Item"`
	SourceIp string           `position:"Query" name:"SourceIp"`
	Status   requests.Integer `position:"Query" name:"Status"`
}

// ModifyLoginSwitchConfigResponse is the response struct for api ModifyLoginSwitchConfig
type ModifyLoginSwitchConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyLoginSwitchConfigRequest creates a request to invoke ModifyLoginSwitchConfig API
func CreateModifyLoginSwitchConfigRequest() (request *ModifyLoginSwitchConfigRequest) {
	request = &ModifyLoginSwitchConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyLoginSwitchConfig", "sas", "openAPI")
	return
}

// CreateModifyLoginSwitchConfigResponse creates a response to parse from ModifyLoginSwitchConfig response
func CreateModifyLoginSwitchConfigResponse() (response *ModifyLoginSwitchConfigResponse) {
	response = &ModifyLoginSwitchConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
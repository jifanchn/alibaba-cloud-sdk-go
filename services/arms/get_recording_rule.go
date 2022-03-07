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

// GetRecordingRule invokes the arms.GetRecordingRule API synchronously
func (client *Client) GetRecordingRule(request *GetRecordingRuleRequest) (response *GetRecordingRuleResponse, err error) {
	response = CreateGetRecordingRuleResponse()
	err = client.DoAction(request, response)
	return
}

// GetRecordingRuleWithChan invokes the arms.GetRecordingRule API asynchronously
func (client *Client) GetRecordingRuleWithChan(request *GetRecordingRuleRequest) (<-chan *GetRecordingRuleResponse, <-chan error) {
	responseChan := make(chan *GetRecordingRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRecordingRule(request)
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

// GetRecordingRuleWithCallback invokes the arms.GetRecordingRule API asynchronously
func (client *Client) GetRecordingRuleWithCallback(request *GetRecordingRuleRequest, callback func(response *GetRecordingRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRecordingRuleResponse
		var err error
		defer close(result)
		response, err = client.GetRecordingRule(request)
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

// GetRecordingRuleRequest is the request struct for api GetRecordingRule
type GetRecordingRuleRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// GetRecordingRuleResponse is the response struct for api GetRecordingRule
type GetRecordingRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetRecordingRuleRequest creates a request to invoke GetRecordingRule API
func CreateGetRecordingRuleRequest() (request *GetRecordingRuleRequest) {
	request = &GetRecordingRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetRecordingRule", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetRecordingRuleResponse creates a response to parse from GetRecordingRule response
func CreateGetRecordingRuleResponse() (response *GetRecordingRuleResponse) {
	response = &GetRecordingRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

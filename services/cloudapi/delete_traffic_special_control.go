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

// DeleteTrafficSpecialControl invokes the cloudapi.DeleteTrafficSpecialControl API synchronously
// api document: https://help.aliyun.com/api/cloudapi/deletetrafficspecialcontrol.html
func (client *Client) DeleteTrafficSpecialControl(request *DeleteTrafficSpecialControlRequest) (response *DeleteTrafficSpecialControlResponse, err error) {
	response = CreateDeleteTrafficSpecialControlResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTrafficSpecialControlWithChan invokes the cloudapi.DeleteTrafficSpecialControl API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/deletetrafficspecialcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTrafficSpecialControlWithChan(request *DeleteTrafficSpecialControlRequest) (<-chan *DeleteTrafficSpecialControlResponse, <-chan error) {
	responseChan := make(chan *DeleteTrafficSpecialControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTrafficSpecialControl(request)
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

// DeleteTrafficSpecialControlWithCallback invokes the cloudapi.DeleteTrafficSpecialControl API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/deletetrafficspecialcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTrafficSpecialControlWithCallback(request *DeleteTrafficSpecialControlRequest, callback func(response *DeleteTrafficSpecialControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTrafficSpecialControlResponse
		var err error
		defer close(result)
		response, err = client.DeleteTrafficSpecialControl(request)
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

// DeleteTrafficSpecialControlRequest is the request struct for api DeleteTrafficSpecialControl
type DeleteTrafficSpecialControlRequest struct {
	*requests.RpcRequest
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	SpecialKey       string `position:"Query" name:"SpecialKey"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	SpecialType      string `position:"Query" name:"SpecialType"`
}

// DeleteTrafficSpecialControlResponse is the response struct for api DeleteTrafficSpecialControl
type DeleteTrafficSpecialControlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTrafficSpecialControlRequest creates a request to invoke DeleteTrafficSpecialControl API
func CreateDeleteTrafficSpecialControlRequest() (request *DeleteTrafficSpecialControlRequest) {
	request = &DeleteTrafficSpecialControlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteTrafficSpecialControl", "", "")
	return
}

// CreateDeleteTrafficSpecialControlResponse creates a response to parse from DeleteTrafficSpecialControl response
func CreateDeleteTrafficSpecialControlResponse() (response *DeleteTrafficSpecialControlResponse) {
	response = &DeleteTrafficSpecialControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

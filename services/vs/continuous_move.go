package vs

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

// ContinuousMove invokes the vs.ContinuousMove API synchronously
func (client *Client) ContinuousMove(request *ContinuousMoveRequest) (response *ContinuousMoveResponse, err error) {
	response = CreateContinuousMoveResponse()
	err = client.DoAction(request, response)
	return
}

// ContinuousMoveWithChan invokes the vs.ContinuousMove API asynchronously
func (client *Client) ContinuousMoveWithChan(request *ContinuousMoveRequest) (<-chan *ContinuousMoveResponse, <-chan error) {
	responseChan := make(chan *ContinuousMoveResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ContinuousMove(request)
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

// ContinuousMoveWithCallback invokes the vs.ContinuousMove API asynchronously
func (client *Client) ContinuousMoveWithCallback(request *ContinuousMoveRequest, callback func(response *ContinuousMoveResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ContinuousMoveResponse
		var err error
		defer close(result)
		response, err = client.ContinuousMove(request)
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

// ContinuousMoveRequest is the request struct for api ContinuousMove
type ContinuousMoveRequest struct {
	*requests.RpcRequest
	Tilt        string           `position:"Query" name:"Tilt"`
	SubProtocol string           `position:"Query" name:"SubProtocol"`
	Id          string           `position:"Query" name:"Id"`
	Pan         string           `position:"Query" name:"Pan"`
	ShowLog     string           `position:"Query" name:"ShowLog"`
	Zoom        string           `position:"Query" name:"Zoom"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}

// ContinuousMoveResponse is the response struct for api ContinuousMove
type ContinuousMoveResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateContinuousMoveRequest creates a request to invoke ContinuousMove API
func CreateContinuousMoveRequest() (request *ContinuousMoveRequest) {
	request = &ContinuousMoveRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "ContinuousMove", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateContinuousMoveResponse creates a response to parse from ContinuousMove response
func CreateContinuousMoveResponse() (response *ContinuousMoveResponse) {
	response = &ContinuousMoveResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package retailcloud

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

// CreateSlbAP invokes the retailcloud.CreateSlbAP API synchronously
func (client *Client) CreateSlbAP(request *CreateSlbAPRequest) (response *CreateSlbAPResponse, err error) {
	response = CreateCreateSlbAPResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSlbAPWithChan invokes the retailcloud.CreateSlbAP API asynchronously
func (client *Client) CreateSlbAPWithChan(request *CreateSlbAPRequest) (<-chan *CreateSlbAPResponse, <-chan error) {
	responseChan := make(chan *CreateSlbAPResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSlbAP(request)
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

// CreateSlbAPWithCallback invokes the retailcloud.CreateSlbAP API asynchronously
func (client *Client) CreateSlbAPWithCallback(request *CreateSlbAPRequest, callback func(response *CreateSlbAPResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSlbAPResponse
		var err error
		defer close(result)
		response, err = client.CreateSlbAP(request)
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

// CreateSlbAPRequest is the request struct for api CreateSlbAP
type CreateSlbAPRequest struct {
	*requests.RpcRequest
	SslCertId          string           `position:"Query" name:"SslCertId"`
	ListenerPort       requests.Integer `position:"Query" name:"ListenerPort"`
	Protocol           string           `position:"Query" name:"Protocol"`
	EstablishedTimeout requests.Integer `position:"Query" name:"EstablishedTimeout"`
	SlbId              string           `position:"Query" name:"SlbId"`
	RealServerPort     requests.Integer `position:"Query" name:"RealServerPort"`
	StickySession      requests.Integer `position:"Query" name:"StickySession"`
	CookieTimeout      requests.Integer `position:"Query" name:"CookieTimeout"`
	Name               string           `position:"Query" name:"Name"`
	EnvId              requests.Integer `position:"Query" name:"EnvId"`
}

// CreateSlbAPResponse is the response struct for api CreateSlbAP
type CreateSlbAPResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateCreateSlbAPRequest creates a request to invoke CreateSlbAP API
func CreateCreateSlbAPRequest() (request *CreateSlbAPRequest) {
	request = &CreateSlbAPRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "CreateSlbAP", "retailcloud", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSlbAPResponse creates a response to parse from CreateSlbAP response
func CreateCreateSlbAPResponse() (response *CreateSlbAPResponse) {
	response = &CreateSlbAPResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

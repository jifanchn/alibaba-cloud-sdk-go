package ocr

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

// RecognizeVietnamIdentityCard invokes the ocr.RecognizeVietnamIdentityCard API synchronously
func (client *Client) RecognizeVietnamIdentityCard(request *RecognizeVietnamIdentityCardRequest) (response *RecognizeVietnamIdentityCardResponse, err error) {
	response = CreateRecognizeVietnamIdentityCardResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeVietnamIdentityCardWithChan invokes the ocr.RecognizeVietnamIdentityCard API asynchronously
func (client *Client) RecognizeVietnamIdentityCardWithChan(request *RecognizeVietnamIdentityCardRequest) (<-chan *RecognizeVietnamIdentityCardResponse, <-chan error) {
	responseChan := make(chan *RecognizeVietnamIdentityCardResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeVietnamIdentityCard(request)
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

// RecognizeVietnamIdentityCardWithCallback invokes the ocr.RecognizeVietnamIdentityCard API asynchronously
func (client *Client) RecognizeVietnamIdentityCardWithCallback(request *RecognizeVietnamIdentityCardRequest, callback func(response *RecognizeVietnamIdentityCardResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeVietnamIdentityCardResponse
		var err error
		defer close(result)
		response, err = client.RecognizeVietnamIdentityCard(request)
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

// RecognizeVietnamIdentityCardRequest is the request struct for api RecognizeVietnamIdentityCard
type RecognizeVietnamIdentityCardRequest struct {
	*requests.RpcRequest
	ImageUrl string `position:"Body" name:"ImageUrl"`
}

// RecognizeVietnamIdentityCardResponse is the response struct for api RecognizeVietnamIdentityCard
type RecognizeVietnamIdentityCardResponse struct {
	*responses.BaseResponse
	RequestId string                             `json:"RequestId" xml:"RequestId"`
	Code      string                             `json:"Code" xml:"Code"`
	Message   string                             `json:"Message" xml:"Message"`
	Data      DataInRecognizeVietnamIdentityCard `json:"Data" xml:"Data"`
}

// CreateRecognizeVietnamIdentityCardRequest creates a request to invoke RecognizeVietnamIdentityCard API
func CreateRecognizeVietnamIdentityCardRequest() (request *RecognizeVietnamIdentityCardRequest) {
	request = &RecognizeVietnamIdentityCardRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ocr", "2019-12-30", "RecognizeVietnamIdentityCard", "ocr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecognizeVietnamIdentityCardResponse creates a response to parse from RecognizeVietnamIdentityCard response
func CreateRecognizeVietnamIdentityCardResponse() (response *RecognizeVietnamIdentityCardResponse) {
	response = &RecognizeVietnamIdentityCardResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

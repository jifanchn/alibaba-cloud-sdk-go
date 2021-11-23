package rtc

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

// StopChannelUserPublish invokes the rtc.StopChannelUserPublish API synchronously
func (client *Client) StopChannelUserPublish(request *StopChannelUserPublishRequest) (response *StopChannelUserPublishResponse, err error) {
	response = CreateStopChannelUserPublishResponse()
	err = client.DoAction(request, response)
	return
}

// StopChannelUserPublishWithChan invokes the rtc.StopChannelUserPublish API asynchronously
func (client *Client) StopChannelUserPublishWithChan(request *StopChannelUserPublishRequest) (<-chan *StopChannelUserPublishResponse, <-chan error) {
	responseChan := make(chan *StopChannelUserPublishResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopChannelUserPublish(request)
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

// StopChannelUserPublishWithCallback invokes the rtc.StopChannelUserPublish API asynchronously
func (client *Client) StopChannelUserPublishWithCallback(request *StopChannelUserPublishRequest, callback func(response *StopChannelUserPublishResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopChannelUserPublishResponse
		var err error
		defer close(result)
		response, err = client.StopChannelUserPublish(request)
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

// StopChannelUserPublishRequest is the request struct for api StopChannelUserPublish
type StopChannelUserPublishRequest struct {
	*requests.RpcRequest
	UserId    string           `position:"Query" name:"UserId"`
	ShowLog   string           `position:"Query" name:"ShowLog"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	AppId     string           `position:"Query" name:"AppId"`
	ChannelId string           `position:"Query" name:"ChannelId"`
}

// StopChannelUserPublishResponse is the response struct for api StopChannelUserPublish
type StopChannelUserPublishResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopChannelUserPublishRequest creates a request to invoke StopChannelUserPublish API
func CreateStopChannelUserPublishRequest() (request *StopChannelUserPublishRequest) {
	request = &StopChannelUserPublishRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "StopChannelUserPublish", "", "")
	request.Method = requests.POST
	return
}

// CreateStopChannelUserPublishResponse creates a response to parse from StopChannelUserPublish response
func CreateStopChannelUserPublishResponse() (response *StopChannelUserPublishResponse) {
	response = &StopChannelUserPublishResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

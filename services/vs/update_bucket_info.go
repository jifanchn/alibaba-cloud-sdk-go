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

// UpdateBucketInfo invokes the vs.UpdateBucketInfo API synchronously
func (client *Client) UpdateBucketInfo(request *UpdateBucketInfoRequest) (response *UpdateBucketInfoResponse, err error) {
	response = CreateUpdateBucketInfoResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateBucketInfoWithChan invokes the vs.UpdateBucketInfo API asynchronously
func (client *Client) UpdateBucketInfoWithChan(request *UpdateBucketInfoRequest) (<-chan *UpdateBucketInfoResponse, <-chan error) {
	responseChan := make(chan *UpdateBucketInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateBucketInfo(request)
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

// UpdateBucketInfoWithCallback invokes the vs.UpdateBucketInfo API asynchronously
func (client *Client) UpdateBucketInfoWithCallback(request *UpdateBucketInfoRequest, callback func(response *UpdateBucketInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateBucketInfoResponse
		var err error
		defer close(result)
		response, err = client.UpdateBucketInfo(request)
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

// UpdateBucketInfoRequest is the request struct for api UpdateBucketInfo
type UpdateBucketInfoRequest struct {
	*requests.RpcRequest
	BucketName string           `position:"Query" name:"BucketName"`
	ShowLog    string           `position:"Query" name:"ShowLog"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	Comment    string           `position:"Query" name:"Comment"`
}

// UpdateBucketInfoResponse is the response struct for api UpdateBucketInfo
type UpdateBucketInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateBucketInfoRequest creates a request to invoke UpdateBucketInfo API
func CreateUpdateBucketInfoRequest() (request *UpdateBucketInfoRequest) {
	request = &UpdateBucketInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "UpdateBucketInfo", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateBucketInfoResponse creates a response to parse from UpdateBucketInfo response
func CreateUpdateBucketInfoResponse() (response *UpdateBucketInfoResponse) {
	response = &UpdateBucketInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

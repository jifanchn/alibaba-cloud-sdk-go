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

// GetBucketInfo invokes the vs.GetBucketInfo API synchronously
func (client *Client) GetBucketInfo(request *GetBucketInfoRequest) (response *GetBucketInfoResponse, err error) {
	response = CreateGetBucketInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetBucketInfoWithChan invokes the vs.GetBucketInfo API asynchronously
func (client *Client) GetBucketInfoWithChan(request *GetBucketInfoRequest) (<-chan *GetBucketInfoResponse, <-chan error) {
	responseChan := make(chan *GetBucketInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBucketInfo(request)
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

// GetBucketInfoWithCallback invokes the vs.GetBucketInfo API asynchronously
func (client *Client) GetBucketInfoWithCallback(request *GetBucketInfoRequest, callback func(response *GetBucketInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBucketInfoResponse
		var err error
		defer close(result)
		response, err = client.GetBucketInfo(request)
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

// GetBucketInfoRequest is the request struct for api GetBucketInfo
type GetBucketInfoRequest struct {
	*requests.RpcRequest
	BucketName string           `position:"Query" name:"BucketName"`
	ShowLog    string           `position:"Query" name:"ShowLog"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// GetBucketInfoResponse is the response struct for api GetBucketInfo
type GetBucketInfoResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	BucketInfo BucketInfo `json:"BucketInfo" xml:"BucketInfo"`
}

// CreateGetBucketInfoRequest creates a request to invoke GetBucketInfo API
func CreateGetBucketInfoRequest() (request *GetBucketInfoRequest) {
	request = &GetBucketInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "GetBucketInfo", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetBucketInfoResponse creates a response to parse from GetBucketInfo response
func CreateGetBucketInfoResponse() (response *GetBucketInfoResponse) {
	response = &GetBucketInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

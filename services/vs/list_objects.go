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

// ListObjects invokes the vs.ListObjects API synchronously
func (client *Client) ListObjects(request *ListObjectsRequest) (response *ListObjectsResponse, err error) {
	response = CreateListObjectsResponse()
	err = client.DoAction(request, response)
	return
}

// ListObjectsWithChan invokes the vs.ListObjects API asynchronously
func (client *Client) ListObjectsWithChan(request *ListObjectsRequest) (<-chan *ListObjectsResponse, <-chan error) {
	responseChan := make(chan *ListObjectsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListObjects(request)
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

// ListObjectsWithCallback invokes the vs.ListObjects API asynchronously
func (client *Client) ListObjectsWithCallback(request *ListObjectsRequest, callback func(response *ListObjectsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListObjectsResponse
		var err error
		defer close(result)
		response, err = client.ListObjects(request)
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

// ListObjectsRequest is the request struct for api ListObjects
type ListObjectsRequest struct {
	*requests.RpcRequest
	MaxKeys           requests.Integer `position:"Query" name:"MaxKeys"`
	ContinuationToken string           `position:"Query" name:"ContinuationToken"`
	Prefix            string           `position:"Query" name:"Prefix"`
	Delimiter         string           `position:"Query" name:"Delimiter"`
	BucketName        string           `position:"Query" name:"BucketName"`
	ShowLog           string           `position:"Query" name:"ShowLog"`
	EncodingType      string           `position:"Query" name:"EncodingType"`
	StartAfter        string           `position:"Query" name:"StartAfter"`
	OwnerId           requests.Integer `position:"Query" name:"OwnerId"`
	Marker            string           `position:"Query" name:"Marker"`
}

// ListObjectsResponse is the response struct for api ListObjects
type ListObjectsResponse struct {
	*responses.BaseResponse
	RequestId             string    `json:"RequestId" xml:"RequestId"`
	BucketName            string    `json:"BucketName" xml:"BucketName"`
	Prefix                string    `json:"Prefix" xml:"Prefix"`
	MaxKeys               int       `json:"MaxKeys" xml:"MaxKeys"`
	KeyCount              int       `json:"KeyCount" xml:"KeyCount"`
	Delimiter             string    `json:"Delimiter" xml:"Delimiter"`
	EncodingType          string    `json:"EncodingType" xml:"EncodingType"`
	IsTruncated           bool      `json:"IsTruncated" xml:"IsTruncated"`
	ContinuationToken     string    `json:"ContinuationToken" xml:"ContinuationToken"`
	NextContinuationToken string    `json:"NextContinuationToken" xml:"NextContinuationToken"`
	Marker                string    `json:"Marker" xml:"Marker"`
	NextMarker            string    `json:"NextMarker" xml:"NextMarker"`
	CommonPrefixes        []string  `json:"CommonPrefixes" xml:"CommonPrefixes"`
	Contents              []Content `json:"Contents" xml:"Contents"`
}

// CreateListObjectsRequest creates a request to invoke ListObjects API
func CreateListObjectsRequest() (request *ListObjectsRequest) {
	request = &ListObjectsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "ListObjects", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListObjectsResponse creates a response to parse from ListObjects response
func CreateListObjectsResponse() (response *ListObjectsResponse) {
	response = &ListObjectsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

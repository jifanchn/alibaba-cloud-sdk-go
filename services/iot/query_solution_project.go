package iot

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

// QuerySolutionProject invokes the iot.QuerySolutionProject API synchronously
func (client *Client) QuerySolutionProject(request *QuerySolutionProjectRequest) (response *QuerySolutionProjectResponse, err error) {
	response = CreateQuerySolutionProjectResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySolutionProjectWithChan invokes the iot.QuerySolutionProject API asynchronously
func (client *Client) QuerySolutionProjectWithChan(request *QuerySolutionProjectRequest) (<-chan *QuerySolutionProjectResponse, <-chan error) {
	responseChan := make(chan *QuerySolutionProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySolutionProject(request)
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

// QuerySolutionProjectWithCallback invokes the iot.QuerySolutionProject API asynchronously
func (client *Client) QuerySolutionProjectWithCallback(request *QuerySolutionProjectRequest, callback func(response *QuerySolutionProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySolutionProjectResponse
		var err error
		defer close(result)
		response, err = client.QuerySolutionProject(request)
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

// QuerySolutionProjectRequest is the request struct for api QuerySolutionProject
type QuerySolutionProjectRequest struct {
	*requests.RpcRequest
	Type        string           `position:"Body" name:"Type"`
	PageId      requests.Integer `position:"Body" name:"PageId"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	ApiProduct  string           `position:"Body" name:"ApiProduct"`
	Name        string           `position:"Body" name:"Name"`
	ApiRevision string           `position:"Body" name:"ApiRevision"`
}

// QuerySolutionProjectResponse is the response struct for api QuerySolutionProject
type QuerySolutionProjectResponse struct {
	*responses.BaseResponse
	RequestId    string                     `json:"RequestId" xml:"RequestId"`
	Success      bool                       `json:"Success" xml:"Success"`
	Code         string                     `json:"Code" xml:"Code"`
	ErrorMessage string                     `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQuerySolutionProject `json:"Data" xml:"Data"`
}

// CreateQuerySolutionProjectRequest creates a request to invoke QuerySolutionProject API
func CreateQuerySolutionProjectRequest() (request *QuerySolutionProjectRequest) {
	request = &QuerySolutionProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QuerySolutionProject", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySolutionProjectResponse creates a response to parse from QuerySolutionProject response
func CreateQuerySolutionProjectResponse() (response *QuerySolutionProjectResponse) {
	response = &QuerySolutionProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

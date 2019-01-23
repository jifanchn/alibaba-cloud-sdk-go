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

// DescribeApiHistory invokes the cloudapi.DescribeApiHistory API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapihistory.html
func (client *Client) DescribeApiHistory(request *DescribeApiHistoryRequest) (response *DescribeApiHistoryResponse, err error) {
	response = CreateDescribeApiHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApiHistoryWithChan invokes the cloudapi.DescribeApiHistory API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapihistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiHistoryWithChan(request *DescribeApiHistoryRequest) (<-chan *DescribeApiHistoryResponse, <-chan error) {
	responseChan := make(chan *DescribeApiHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApiHistory(request)
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

// DescribeApiHistoryWithCallback invokes the cloudapi.DescribeApiHistory API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapihistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiHistoryWithCallback(request *DescribeApiHistoryRequest, callback func(response *DescribeApiHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiHistoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeApiHistory(request)
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

// DescribeApiHistoryRequest is the request struct for api DescribeApiHistory
type DescribeApiHistoryRequest struct {
	*requests.RpcRequest
	StageName      string `position:"Query" name:"StageName"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	GroupId        string `position:"Query" name:"GroupId"`
	ApiId          string `position:"Query" name:"ApiId"`
	HistoryVersion string `position:"Query" name:"HistoryVersion"`
}

// DescribeApiHistoryResponse is the response struct for api DescribeApiHistory
type DescribeApiHistoryResponse struct {
	*responses.BaseResponse
	RequestId              string                                     `json:"RequestId" xml:"RequestId"`
	RegionId               string                                     `json:"RegionId" xml:"RegionId"`
	GroupId                string                                     `json:"GroupId" xml:"GroupId"`
	GroupName              string                                     `json:"GroupName" xml:"GroupName"`
	StageName              string                                     `json:"StageName" xml:"StageName"`
	ApiId                  string                                     `json:"ApiId" xml:"ApiId"`
	ApiName                string                                     `json:"ApiName" xml:"ApiName"`
	Description            string                                     `json:"Description" xml:"Description"`
	HistoryVersion         string                                     `json:"HistoryVersion" xml:"HistoryVersion"`
	Status                 string                                     `json:"Status" xml:"Status"`
	Visibility             string                                     `json:"Visibility" xml:"Visibility"`
	AuthType               string                                     `json:"AuthType" xml:"AuthType"`
	ResultType             string                                     `json:"ResultType" xml:"ResultType"`
	ResultSample           string                                     `json:"ResultSample" xml:"ResultSample"`
	FailResultSample       string                                     `json:"FailResultSample" xml:"FailResultSample"`
	DeployedTime           string                                     `json:"DeployedTime" xml:"DeployedTime"`
	AllowSignatureMethod   string                                     `json:"AllowSignatureMethod" xml:"AllowSignatureMethod"`
	ResultBodyModel        string                                     `json:"ResultBodyModel" xml:"ResultBodyModel"`
	RequestConfig          RequestConfig                              `json:"RequestConfig" xml:"RequestConfig"`
	ServiceConfig          ServiceConfig                              `json:"ServiceConfig" xml:"ServiceConfig"`
	OpenIdConnectConfig    OpenIdConnectConfig                        `json:"OpenIdConnectConfig" xml:"OpenIdConnectConfig"`
	ErrorCodeSamples       ErrorCodeSamplesInDescribeApiHistory       `json:"ErrorCodeSamples" xml:"ErrorCodeSamples"`
	ResultDescriptions     ResultDescriptionsInDescribeApiHistory     `json:"ResultDescriptions" xml:"ResultDescriptions"`
	SystemParameters       SystemParametersInDescribeApiHistory       `json:"SystemParameters" xml:"SystemParameters"`
	CustomSystemParameters CustomSystemParametersInDescribeApiHistory `json:"CustomSystemParameters" xml:"CustomSystemParameters"`
	ConstantParameters     ConstantParametersInDescribeApiHistory     `json:"ConstantParameters" xml:"ConstantParameters"`
	RequestParameters      RequestParametersInDescribeApiHistory      `json:"RequestParameters" xml:"RequestParameters"`
	ServiceParameters      ServiceParametersInDescribeApiHistory      `json:"ServiceParameters" xml:"ServiceParameters"`
	ServiceParametersMap   ServiceParametersMapInDescribeApiHistory   `json:"ServiceParametersMap" xml:"ServiceParametersMap"`
}

// CreateDescribeApiHistoryRequest creates a request to invoke DescribeApiHistory API
func CreateDescribeApiHistoryRequest() (request *DescribeApiHistoryRequest) {
	request = &DescribeApiHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiHistory", "", "")
	return
}

// CreateDescribeApiHistoryResponse creates a response to parse from DescribeApiHistory response
func CreateDescribeApiHistoryResponse() (response *DescribeApiHistoryResponse) {
	response = &DescribeApiHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

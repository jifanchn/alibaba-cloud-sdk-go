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

// UpdateDeployConfig invokes the retailcloud.UpdateDeployConfig API synchronously
func (client *Client) UpdateDeployConfig(request *UpdateDeployConfigRequest) (response *UpdateDeployConfigResponse, err error) {
	response = CreateUpdateDeployConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDeployConfigWithChan invokes the retailcloud.UpdateDeployConfig API asynchronously
func (client *Client) UpdateDeployConfigWithChan(request *UpdateDeployConfigRequest) (<-chan *UpdateDeployConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateDeployConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDeployConfig(request)
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

// UpdateDeployConfigWithCallback invokes the retailcloud.UpdateDeployConfig API asynchronously
func (client *Client) UpdateDeployConfigWithCallback(request *UpdateDeployConfigRequest, callback func(response *UpdateDeployConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDeployConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateDeployConfig(request)
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

// UpdateDeployConfigRequest is the request struct for api UpdateDeployConfig
type UpdateDeployConfigRequest struct {
	*requests.RpcRequest
	CodePath      string           `position:"Query" name:"CodePath"`
	ConfigMapList *[]string        `position:"Query" name:"ConfigMapList"  type:"Repeated"`
	ConfigMap     string           `position:"Query" name:"ConfigMap"`
	StatefulSet   string           `position:"Query" name:"StatefulSet"`
	AppId         requests.Integer `position:"Query" name:"AppId"`
	SecretList    *[]string        `position:"Query" name:"SecretList"  type:"Repeated"`
	Id            requests.Integer `position:"Query" name:"Id"`
	CronJob       string           `position:"Query" name:"CronJob"`
	Deployment    string           `position:"Query" name:"Deployment"`
}

// UpdateDeployConfigResponse is the response struct for api UpdateDeployConfig
type UpdateDeployConfigResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateUpdateDeployConfigRequest creates a request to invoke UpdateDeployConfig API
func CreateUpdateDeployConfigRequest() (request *UpdateDeployConfigRequest) {
	request = &UpdateDeployConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "UpdateDeployConfig", "retailcloud", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateDeployConfigResponse creates a response to parse from UpdateDeployConfig response
func CreateUpdateDeployConfigResponse() (response *UpdateDeployConfigResponse) {
	response = &UpdateDeployConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

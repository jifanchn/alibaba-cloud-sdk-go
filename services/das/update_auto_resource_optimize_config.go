package das

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

// UpdateAutoResourceOptimizeConfig invokes the das.UpdateAutoResourceOptimizeConfig API synchronously
func (client *Client) UpdateAutoResourceOptimizeConfig(request *UpdateAutoResourceOptimizeConfigRequest) (response *UpdateAutoResourceOptimizeConfigResponse, err error) {
	response = CreateUpdateAutoResourceOptimizeConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAutoResourceOptimizeConfigWithChan invokes the das.UpdateAutoResourceOptimizeConfig API asynchronously
func (client *Client) UpdateAutoResourceOptimizeConfigWithChan(request *UpdateAutoResourceOptimizeConfigRequest) (<-chan *UpdateAutoResourceOptimizeConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateAutoResourceOptimizeConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAutoResourceOptimizeConfig(request)
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

// UpdateAutoResourceOptimizeConfigWithCallback invokes the das.UpdateAutoResourceOptimizeConfig API asynchronously
func (client *Client) UpdateAutoResourceOptimizeConfigWithCallback(request *UpdateAutoResourceOptimizeConfigRequest, callback func(response *UpdateAutoResourceOptimizeConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAutoResourceOptimizeConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateAutoResourceOptimizeConfig(request)
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

// UpdateAutoResourceOptimizeConfigRequest is the request struct for api UpdateAutoResourceOptimizeConfig
type UpdateAutoResourceOptimizeConfigRequest struct {
	*requests.RpcRequest
	Context                  string           `position:"Query" name:"__context"`
	Signature                string           `position:"Query" name:"Signature"`
	AutoDefragment           requests.Integer `position:"Query" name:"AutoDefragment"`
	AutoDuplicateIndexDelete requests.Integer `position:"Query" name:"AutoDuplicateIndexDelete"`
	UserId                   string           `position:"Query" name:"UserId"`
	Uid                      string           `position:"Query" name:"Uid"`
	InstanceId               string           `position:"Query" name:"InstanceId"`
	AccessKey                string           `position:"Query" name:"AccessKey"`
	TableSpaceSize           requests.Float   `position:"Query" name:"TableSpaceSize"`
	TableFragmentationRatio  requests.Float   `position:"Query" name:"TableFragmentationRatio"`
}

// UpdateAutoResourceOptimizeConfigResponse is the response struct for api UpdateAutoResourceOptimizeConfig
type UpdateAutoResourceOptimizeConfigResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Synchro   string `json:"Synchro" xml:"Synchro"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateUpdateAutoResourceOptimizeConfigRequest creates a request to invoke UpdateAutoResourceOptimizeConfig API
func CreateUpdateAutoResourceOptimizeConfigRequest() (request *UpdateAutoResourceOptimizeConfigRequest) {
	request = &UpdateAutoResourceOptimizeConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "UpdateAutoResourceOptimizeConfig", "das", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateAutoResourceOptimizeConfigResponse creates a response to parse from UpdateAutoResourceOptimizeConfig response
func CreateUpdateAutoResourceOptimizeConfigResponse() (response *UpdateAutoResourceOptimizeConfigResponse) {
	response = &UpdateAutoResourceOptimizeConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

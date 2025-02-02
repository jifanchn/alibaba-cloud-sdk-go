package arms

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

// InstallManagedPrometheus invokes the arms.InstallManagedPrometheus API synchronously
func (client *Client) InstallManagedPrometheus(request *InstallManagedPrometheusRequest) (response *InstallManagedPrometheusResponse, err error) {
	response = CreateInstallManagedPrometheusResponse()
	err = client.DoAction(request, response)
	return
}

// InstallManagedPrometheusWithChan invokes the arms.InstallManagedPrometheus API asynchronously
func (client *Client) InstallManagedPrometheusWithChan(request *InstallManagedPrometheusRequest) (<-chan *InstallManagedPrometheusResponse, <-chan error) {
	responseChan := make(chan *InstallManagedPrometheusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InstallManagedPrometheus(request)
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

// InstallManagedPrometheusWithCallback invokes the arms.InstallManagedPrometheus API asynchronously
func (client *Client) InstallManagedPrometheusWithCallback(request *InstallManagedPrometheusRequest, callback func(response *InstallManagedPrometheusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InstallManagedPrometheusResponse
		var err error
		defer close(result)
		response, err = client.InstallManagedPrometheus(request)
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

// InstallManagedPrometheusRequest is the request struct for api InstallManagedPrometheus
type InstallManagedPrometheusRequest struct {
	*requests.RpcRequest
	SecurityGroupId string `position:"Query" name:"SecurityGroupId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	KubeConfig      string `position:"Query" name:"KubeConfig"`
	ClusterType     string `position:"Query" name:"ClusterType"`
	VSwitchId       string `position:"Query" name:"VSwitchId"`
	VpcId           string `position:"Query" name:"VpcId"`
}

// InstallManagedPrometheusResponse is the response struct for api InstallManagedPrometheus
type InstallManagedPrometheusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Code      int    `json:"Code" xml:"Code"`
}

// CreateInstallManagedPrometheusRequest creates a request to invoke InstallManagedPrometheus API
func CreateInstallManagedPrometheusRequest() (request *InstallManagedPrometheusRequest) {
	request = &InstallManagedPrometheusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "InstallManagedPrometheus", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInstallManagedPrometheusResponse creates a response to parse from InstallManagedPrometheus response
func CreateInstallManagedPrometheusResponse() (response *InstallManagedPrometheusResponse) {
	response = &InstallManagedPrometheusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

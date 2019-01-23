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

// DeleteDomainCertificate invokes the cloudapi.DeleteDomainCertificate API synchronously
// api document: https://help.aliyun.com/api/cloudapi/deletedomaincertificate.html
func (client *Client) DeleteDomainCertificate(request *DeleteDomainCertificateRequest) (response *DeleteDomainCertificateResponse, err error) {
	response = CreateDeleteDomainCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDomainCertificateWithChan invokes the cloudapi.DeleteDomainCertificate API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/deletedomaincertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDomainCertificateWithChan(request *DeleteDomainCertificateRequest) (<-chan *DeleteDomainCertificateResponse, <-chan error) {
	responseChan := make(chan *DeleteDomainCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDomainCertificate(request)
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

// DeleteDomainCertificateWithCallback invokes the cloudapi.DeleteDomainCertificate API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/deletedomaincertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDomainCertificateWithCallback(request *DeleteDomainCertificateRequest, callback func(response *DeleteDomainCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDomainCertificateResponse
		var err error
		defer close(result)
		response, err = client.DeleteDomainCertificate(request)
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

// DeleteDomainCertificateRequest is the request struct for api DeleteDomainCertificate
type DeleteDomainCertificateRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GroupId       string `position:"Query" name:"GroupId"`
	CertificateId string `position:"Query" name:"CertificateId"`
	DomainName    string `position:"Query" name:"DomainName"`
}

// DeleteDomainCertificateResponse is the response struct for api DeleteDomainCertificate
type DeleteDomainCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDomainCertificateRequest creates a request to invoke DeleteDomainCertificate API
func CreateDeleteDomainCertificateRequest() (request *DeleteDomainCertificateRequest) {
	request = &DeleteDomainCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteDomainCertificate", "", "")
	return
}

// CreateDeleteDomainCertificateResponse creates a response to parse from DeleteDomainCertificate response
func CreateDeleteDomainCertificateResponse() (response *DeleteDomainCertificateResponse) {
	response = &DeleteDomainCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

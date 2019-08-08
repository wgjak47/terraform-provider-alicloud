package alidns

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

// DescribeBatchResultDetail invokes the alidns.DescribeBatchResultDetail API synchronously
// api document: https://help.aliyun.com/api/alidns/describebatchresultdetail.html
func (client *Client) DescribeBatchResultDetail(request *DescribeBatchResultDetailRequest) (response *DescribeBatchResultDetailResponse, err error) {
	response = CreateDescribeBatchResultDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBatchResultDetailWithChan invokes the alidns.DescribeBatchResultDetail API asynchronously
// api document: https://help.aliyun.com/api/alidns/describebatchresultdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBatchResultDetailWithChan(request *DescribeBatchResultDetailRequest) (<-chan *DescribeBatchResultDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeBatchResultDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBatchResultDetail(request)
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

// DescribeBatchResultDetailWithCallback invokes the alidns.DescribeBatchResultDetail API asynchronously
// api document: https://help.aliyun.com/api/alidns/describebatchresultdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBatchResultDetailWithCallback(request *DescribeBatchResultDetailRequest, callback func(response *DescribeBatchResultDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBatchResultDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeBatchResultDetail(request)
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

// DescribeBatchResultDetailRequest is the request struct for api DescribeBatchResultDetail
type DescribeBatchResultDetailRequest struct {
	*requests.RpcRequest
	BatchType    string           `position:"Query" name:"BatchType"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	TaskId       requests.Integer `position:"Query" name:"TaskId"`
}

// DescribeBatchResultDetailResponse is the response struct for api DescribeBatchResultDetail
type DescribeBatchResultDetailResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	TotalCount         int64              `json:"TotalCount" xml:"TotalCount"`
	PageNumber         int64              `json:"PageNumber" xml:"PageNumber"`
	PageSize           int64              `json:"PageSize" xml:"PageSize"`
	BatchResultDetails BatchResultDetails `json:"BatchResultDetails" xml:"BatchResultDetails"`
}

// CreateDescribeBatchResultDetailRequest creates a request to invoke DescribeBatchResultDetail API
func CreateDescribeBatchResultDetailRequest() (request *DescribeBatchResultDetailRequest) {
	request = &DescribeBatchResultDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeBatchResultDetail", "Alidns", "openAPI")
	return
}

// CreateDescribeBatchResultDetailResponse creates a response to parse from DescribeBatchResultDetail response
func CreateDescribeBatchResultDetailResponse() (response *DescribeBatchResultDetailResponse) {
	response = &DescribeBatchResultDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

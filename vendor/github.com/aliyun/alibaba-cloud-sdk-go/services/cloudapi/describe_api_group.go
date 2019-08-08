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

// DescribeApiGroup invokes the cloudapi.DescribeApiGroup API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapigroup.html
func (client *Client) DescribeApiGroup(request *DescribeApiGroupRequest) (response *DescribeApiGroupResponse, err error) {
	response = CreateDescribeApiGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApiGroupWithChan invokes the cloudapi.DescribeApiGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapigroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiGroupWithChan(request *DescribeApiGroupRequest) (<-chan *DescribeApiGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeApiGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApiGroup(request)
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

// DescribeApiGroupWithCallback invokes the cloudapi.DescribeApiGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapigroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiGroupWithCallback(request *DescribeApiGroupRequest, callback func(response *DescribeApiGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeApiGroup(request)
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

// DescribeApiGroupRequest is the request struct for api DescribeApiGroup
type DescribeApiGroupRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GroupId       string `position:"Query" name:"GroupId"`
}

// DescribeApiGroupResponse is the response struct for api DescribeApiGroup
type DescribeApiGroupResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	GroupId       string        `json:"GroupId" xml:"GroupId"`
	GroupName     string        `json:"GroupName" xml:"GroupName"`
	SubDomain     string        `json:"SubDomain" xml:"SubDomain"`
	Description   string        `json:"Description" xml:"Description"`
	CreatedTime   string        `json:"CreatedTime" xml:"CreatedTime"`
	ModifiedTime  string        `json:"ModifiedTime" xml:"ModifiedTime"`
	RegionId      string        `json:"RegionId" xml:"RegionId"`
	Status        string        `json:"Status" xml:"Status"`
	BillingStatus string        `json:"BillingStatus" xml:"BillingStatus"`
	IllegalStatus string        `json:"IllegalStatus" xml:"IllegalStatus"`
	TrafficLimit  int           `json:"TrafficLimit" xml:"TrafficLimit"`
	VpcDomain     string        `json:"VpcDomain" xml:"VpcDomain"`
	InstanceType  string        `json:"InstanceType" xml:"InstanceType"`
	InstanceId    string        `json:"InstanceId" xml:"InstanceId"`
	HttpsPolicy   string        `json:"HttpsPolicy" xml:"HttpsPolicy"`
	CustomDomains CustomDomains `json:"CustomDomains" xml:"CustomDomains"`
	StageItems    StageItems    `json:"StageItems" xml:"StageItems"`
}

// CreateDescribeApiGroupRequest creates a request to invoke DescribeApiGroup API
func CreateDescribeApiGroupRequest() (request *DescribeApiGroupRequest) {
	request = &DescribeApiGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiGroup", "apigateway", "openAPI")
	return
}

// CreateDescribeApiGroupResponse creates a response to parse from DescribeApiGroup response
func CreateDescribeApiGroupResponse() (response *DescribeApiGroupResponse) {
	response = &DescribeApiGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

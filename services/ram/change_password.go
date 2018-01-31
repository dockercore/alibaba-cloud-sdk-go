package ram

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

func (client *Client) ChangePassword(request *ChangePasswordRequest) (response *ChangePasswordResponse, err error) {
	response = CreateChangePasswordResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ChangePasswordWithChan(request *ChangePasswordRequest) (<-chan *ChangePasswordResponse, <-chan error) {
	responseChan := make(chan *ChangePasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangePassword(request)
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

func (client *Client) ChangePasswordWithCallback(request *ChangePasswordRequest, callback func(response *ChangePasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangePasswordResponse
		var err error
		defer close(result)
		response, err = client.ChangePassword(request)
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

type ChangePasswordRequest struct {
	*requests.RpcRequest
	OldPassword string `position:"Query" name:"OldPassword"`
	NewPassword string `position:"Query" name:"NewPassword"`
}

type ChangePasswordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateChangePasswordRequest() (request *ChangePasswordRequest) {
	request = &ChangePasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "ChangePassword", "", "")
	return
}

func CreateChangePasswordResponse() (response *ChangePasswordResponse) {
	response = &ChangePasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
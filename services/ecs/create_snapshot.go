package ecs

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

func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
	response = CreateCreateSnapshotResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateSnapshotWithChan(request *CreateSnapshotRequest) (<-chan *CreateSnapshotResponse, <-chan error) {
	responseChan := make(chan *CreateSnapshotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSnapshot(request)
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

func (client *Client) CreateSnapshotWithCallback(request *CreateSnapshotRequest, callback func(response *CreateSnapshotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSnapshotResponse
		var err error
		defer close(result)
		response, err = client.CreateSnapshot(request)
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

type CreateSnapshotRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Query" name:"ClientToken"`
	SnapshotName         string           `position:"Query" name:"SnapshotName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tag5Key              string           `position:"Query" name:"Tag.5.Key"`
	Tag5Value            string           `position:"Query" name:"Tag.5.Value"`
	Tag3Key              string           `position:"Query" name:"Tag.3.Key"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	Tag1Key              string           `position:"Query" name:"Tag.1.Key"`
	Tag2Key              string           `position:"Query" name:"Tag.2.Key"`
	Tag1Value            string           `position:"Query" name:"Tag.1.Value"`
	DiskId               string           `position:"Query" name:"DiskId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Tag4Value            string           `position:"Query" name:"Tag.4.Value"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Tag3Value            string           `position:"Query" name:"Tag.3.Value"`
	Tag2Value            string           `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string           `position:"Query" name:"Tag.4.Key"`
}

type CreateSnapshotResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	SnapshotId string `json:"SnapshotId" xml:"SnapshotId"`
}

func CreateCreateSnapshotRequest() (request *CreateSnapshotRequest) {
	request = &CreateSnapshotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateSnapshot", "ecs", "openAPI")
	return
}

func CreateCreateSnapshotResponse() (response *CreateSnapshotResponse) {
	response = &CreateSnapshotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
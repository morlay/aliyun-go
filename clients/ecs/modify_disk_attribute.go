package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDiskAttributeRequest struct {
	DiskName             string `position:"Query" name:"DiskName"`
	DeleteAutoSnapshot   string `position:"Query" name:"DeleteAutoSnapshot"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	EnableAutoSnapshot   string `position:"Query" name:"EnableAutoSnapshot"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DeleteWithInstance   string `position:"Query" name:"DeleteWithInstance"`
}

func (r ModifyDiskAttributeRequest) Invoke(client *sdk.Client) (response *ModifyDiskAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDiskAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyDiskAttribute", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDiskAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDiskAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDiskAttributeResponse struct {
	RequestId string
}

package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyImageShareGroupPermissionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	AddGroup1            string `position:"Query" name:"AddGroup.1"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RemoveGroup1         string `position:"Query" name:"RemoveGroup.1"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyImageShareGroupPermissionRequest) Invoke(client *sdk.Client) (response *ModifyImageShareGroupPermissionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyImageShareGroupPermissionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyImageShareGroupPermission", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyImageShareGroupPermissionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyImageShareGroupPermissionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyImageShareGroupPermissionResponse struct {
	RequestId string
}

package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckDiskEnableAutoSnapshotValidationRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CheckDiskEnableAutoSnapshotValidationRequest) Invoke(client *sdk.Client) (response *CheckDiskEnableAutoSnapshotValidationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CheckDiskEnableAutoSnapshotValidationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CheckDiskEnableAutoSnapshotValidation", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CheckDiskEnableAutoSnapshotValidationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CheckDiskEnableAutoSnapshotValidationResponse

	err = client.DoAction(&req, &resp)
	return
}

type CheckDiskEnableAutoSnapshotValidationResponse struct {
	RequestId              string
	IsPermitted            string
	AutoSnapshotOccupation int
}

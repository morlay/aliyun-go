package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckDiskEnableAutoSnapshotValidationRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CheckDiskEnableAutoSnapshotValidationRequest) Invoke(client *sdk.Client) (resp *CheckDiskEnableAutoSnapshotValidationResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CheckDiskEnableAutoSnapshotValidation", "ecs", "")
	resp = &CheckDiskEnableAutoSnapshotValidationResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckDiskEnableAutoSnapshotValidationResponse struct {
	responses.BaseResponse
	RequestId              string
	IsPermitted            string
	AutoSnapshotOccupation int
}

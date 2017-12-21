package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateSnapshotRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	SnapshotName         string `position:"Query" name:"SnapshotName"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	DiskId               string `position:"Query" name:"DiskId"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
}

func (r CreateSnapshotRequest) Invoke(client *sdk.Client) (response *CreateSnapshotResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateSnapshotRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateSnapshot", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateSnapshotResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateSnapshotResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateSnapshotResponse struct {
	RequestId  string
	SnapshotId string
}

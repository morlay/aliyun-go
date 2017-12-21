package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDiskRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Description          string `position:"Query" name:"Description"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	DiskName             string `position:"Query" name:"DiskName"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	DiskCategory         string `position:"Query" name:"DiskCategory"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	Size                 int    `position:"Query" name:"Size"`
	Encrypted            string `position:"Query" name:"Encrypted"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
}

func (r CreateDiskRequest) Invoke(client *sdk.Client) (response *CreateDiskResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateDiskRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateDisk", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateDiskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateDiskResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateDiskResponse struct {
	RequestId string
	DiskId    string
}

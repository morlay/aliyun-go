package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateVolumeRequest struct {
	requests.RpcRequest
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	VolumeName           string `position:"Query" name:"VolumeName"`
	VolumeEncrypted      string `position:"Query" name:"VolumeEncrypted"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	VolumeCategory       string `position:"Query" name:"VolumeCategory"`
	Size                 int    `position:"Query" name:"Size"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
}

func (req *CreateVolumeRequest) Invoke(client *sdk.Client) (resp *CreateVolumeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateVolume", "ecs", "")
	resp = &CreateVolumeResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateVolumeResponse struct {
	responses.BaseResponse
	RequestId common.String
	VolumeId  common.String
}

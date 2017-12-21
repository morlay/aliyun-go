package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateImageRequest struct {
	requests.RpcRequest
	DiskDeviceMappings   *CreateImageDiskDeviceMappingList `position:"Query" type:"Repeated" name:"DiskDeviceMapping"`
	Tag4Value            string                            `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64                             `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string                            `position:"Query" name:"SnapshotId"`
	Tag2Key              string                            `position:"Query" name:"Tag.2.Key"`
	ClientToken          string                            `position:"Query" name:"ClientToken"`
	Description          string                            `position:"Query" name:"Description"`
	Tag3Key              string                            `position:"Query" name:"Tag.3.Key"`
	Platform             string                            `position:"Query" name:"Platform"`
	Tag1Value            string                            `position:"Query" name:"Tag.1.Value"`
	ImageName            string                            `position:"Query" name:"ImageName"`
	Tag3Value            string                            `position:"Query" name:"Tag.3.Value"`
	Architecture         string                            `position:"Query" name:"Architecture"`
	Tag5Key              string                            `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string                            `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                            `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                             `position:"Query" name:"OwnerId"`
	Tag5Value            string                            `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string                            `position:"Query" name:"Tag.1.Key"`
	InstanceId           string                            `position:"Query" name:"InstanceId"`
	Tag2Value            string                            `position:"Query" name:"Tag.2.Value"`
	ImageVersion         string                            `position:"Query" name:"ImageVersion"`
	Tag4Key              string                            `position:"Query" name:"Tag.4.Key"`
}

func (req *CreateImageRequest) Invoke(client *sdk.Client) (resp *CreateImageResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateImage", "ecs", "")
	resp = &CreateImageResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateImageDiskDeviceMapping struct {
	Size       int    `name:"Size"`
	SnapshotId string `name:"SnapshotId"`
	Device     string `name:"Device"`
	DiskType   string `name:"DiskType"`
}

type CreateImageResponse struct {
	responses.BaseResponse
	RequestId string
	ImageId   string
}

type CreateImageDiskDeviceMappingList []CreateImageDiskDeviceMapping

func (list *CreateImageDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateImageDiskDeviceMapping)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ImportImageRequest struct {
	requests.RpcRequest
	DiskDeviceMappings   *ImportImageDiskDeviceMappingList `position:"Query" type:"Repeated" name:"DiskDeviceMapping"`
	ResourceOwnerId      int64                             `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                            `position:"Query" name:"ResourceOwnerAccount"`
	ImageName            string                            `position:"Query" name:"ImageName"`
	RoleName             string                            `position:"Query" name:"RoleName"`
	Description          string                            `position:"Query" name:"Description"`
	OSType               string                            `position:"Query" name:"OSType"`
	OwnerId              int64                             `position:"Query" name:"OwnerId"`
	Platform             string                            `position:"Query" name:"Platform"`
	Architecture         string                            `position:"Query" name:"Architecture"`
}

func (req *ImportImageRequest) Invoke(client *sdk.Client) (resp *ImportImageResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ImportImage", "ecs", "")
	resp = &ImportImageResponse{}
	err = client.DoAction(req, resp)
	return
}

type ImportImageDiskDeviceMapping struct {
	Format        string `name:"Format"`
	OSSBucket     string `name:"OSSBucket"`
	OSSObject     string `name:"OSSObject"`
	DiskImSize    int    `name:"DiskImSize"`
	DiskImageSize int    `name:"DiskImageSize"`
	Device        string `name:"Device"`
}

type ImportImageResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskId    common.String
	RegionId  common.String
	ImageId   common.String
}

type ImportImageDiskDeviceMappingList []ImportImageDiskDeviceMapping

func (list *ImportImageDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ImportImageDiskDeviceMapping)
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

package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImportImageRequest struct {
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

func (r ImportImageRequest) Invoke(client *sdk.Client) (response *ImportImageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ImportImageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ImportImage", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ImportImageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ImportImageResponse

	err = client.DoAction(&req, &resp)
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
	RequestId string
	TaskId    string
	RegionId  string
	ImageId   string
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

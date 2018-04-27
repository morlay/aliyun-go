package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResizeDiskRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                     `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string                    `position:"Query" name:"ClusterId"`
	DiskConfigs     *ResizeDiskDiskConfigList `position:"Query" type:"Repeated" name:"DiskConfig"`
}

func (req *ResizeDiskRequest) Invoke(client *sdk.Client) (resp *ResizeDiskResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ResizeDisk", "", "")
	resp = &ResizeDiskResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResizeDiskDiskConfig struct {
	HostGroupId string `name:"HostGroupId"`
	NewDiskSize int    `name:"NewDiskSize"`
}

type ResizeDiskResponse struct {
	responses.BaseResponse
	RequestId string
	ClusterId string
}

type ResizeDiskDiskConfigList []ResizeDiskDiskConfig

func (list *ResizeDiskDiskConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ResizeDiskDiskConfig)
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

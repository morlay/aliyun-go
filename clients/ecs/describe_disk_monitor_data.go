package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDiskMonitorDataRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	DiskId               string `position:"Query" name:"DiskId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDiskMonitorDataRequest) Invoke(client *sdk.Client) (resp *DescribeDiskMonitorDataResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDiskMonitorData", "ecs", "")
	resp = &DescribeDiskMonitorDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDiskMonitorDataResponse struct {
	responses.BaseResponse
	RequestId   common.String
	TotalCount  common.Integer
	MonitorData DescribeDiskMonitorDataDiskMonitorDataList
}

type DescribeDiskMonitorDataDiskMonitorData struct {
	DiskId       common.String
	IOPSRead     common.Integer
	IOPSWrite    common.Integer
	IOPSTotal    common.Integer
	BPSRead      common.Integer
	BPSWrite     common.Integer
	BPSTotal     common.Integer
	LatencyRead  common.Integer
	LatencyWrite common.Integer
	TimeStamp    common.String
}

type DescribeDiskMonitorDataDiskMonitorDataList []DescribeDiskMonitorDataDiskMonitorData

func (list *DescribeDiskMonitorDataDiskMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDiskMonitorDataDiskMonitorData)
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

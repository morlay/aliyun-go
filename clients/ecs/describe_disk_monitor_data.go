package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDiskMonitorDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	DiskId               string `position:"Query" name:"DiskId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDiskMonitorDataRequest) Invoke(client *sdk.Client) (response *DescribeDiskMonitorDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDiskMonitorDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDiskMonitorData", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDiskMonitorDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDiskMonitorDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDiskMonitorDataResponse struct {
	RequestId   string
	TotalCount  int
	MonitorData DescribeDiskMonitorDataDiskMonitorDataList
}

type DescribeDiskMonitorDataDiskMonitorData struct {
	DiskId       string
	IOPSRead     int
	IOPSWrite    int
	IOPSTotal    int
	BPSRead      int
	BPSWrite     int
	BPSTotal     int
	LatencyRead  int
	LatencyWrite int
	TimeStamp    string
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

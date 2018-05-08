package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSnapshotMonitorDataRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeSnapshotMonitorDataRequest) Invoke(client *sdk.Client) (resp *DescribeSnapshotMonitorDataResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSnapshotMonitorData", "ecs", "")
	resp = &DescribeSnapshotMonitorDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSnapshotMonitorDataResponse struct {
	responses.BaseResponse
	RequestId   common.String
	MonitorData DescribeSnapshotMonitorDataDataPointList
}

type DescribeSnapshotMonitorDataDataPoint struct {
	TimeStamp common.String
	Size      common.Long
}

type DescribeSnapshotMonitorDataDataPointList []DescribeSnapshotMonitorDataDataPoint

func (list *DescribeSnapshotMonitorDataDataPointList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotMonitorDataDataPoint)
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

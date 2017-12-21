package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSnapshotMonitorDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeSnapshotMonitorDataRequest) Invoke(client *sdk.Client) (response *DescribeSnapshotMonitorDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSnapshotMonitorDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSnapshotMonitorData", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSnapshotMonitorDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSnapshotMonitorDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSnapshotMonitorDataResponse struct {
	RequestId   string
	MonitorData DescribeSnapshotMonitorDataDataPointList
}

type DescribeSnapshotMonitorDataDataPoint struct {
	TimeStamp string
	Size      int64
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

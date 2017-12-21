package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRealtimeLogsRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	LogStartTime         string `position:"Query" name:"LogStartTime"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	LogEndTime           string `position:"Query" name:"LogEndTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Tags                 string `position:"Query" name:"Tags"`
	LogType              string `position:"Query" name:"LogType"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	PageSize             int    `position:"Query" name:"PageSize"`
}

func (req *DescribeRealtimeLogsRequest) Invoke(client *sdk.Client) (resp *DescribeRealtimeLogsResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeRealtimeLogs", "slb", "")
	resp = &DescribeRealtimeLogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRealtimeLogsResponse struct {
	responses.BaseResponse
	RequestId         string
	PageNumber        int
	PageSize          int
	TotalCount        int64
	Progress          string
	LBRealTimeLogsSet DescribeRealtimeLogsLBRealTimeLogList
}

type DescribeRealtimeLogsLBRealTimeLog struct {
	LogDetail string
}

type DescribeRealtimeLogsLBRealTimeLogList []DescribeRealtimeLogsLBRealTimeLog

func (list *DescribeRealtimeLogsLBRealTimeLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRealtimeLogsLBRealTimeLog)
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

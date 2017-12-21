package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeHistoryMonitorValuesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	IntervalForHistory   string `position:"Query" name:"IntervalForHistory"`
	MonitorKeys          string `position:"Query" name:"MonitorKeys"`
}

func (req *DescribeHistoryMonitorValuesRequest) Invoke(client *sdk.Client) (resp *DescribeHistoryMonitorValuesResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeHistoryMonitorValues", "redisa", "")
	resp = &DescribeHistoryMonitorValuesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeHistoryMonitorValuesResponse struct {
	responses.BaseResponse
	RequestId      string
	MonitorHistory string
}

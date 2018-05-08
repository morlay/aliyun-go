package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeHistoryMonitorValuesRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	EndTime              string `position:"Query" name:"EndTime"`
	IntervalForHistory   string `position:"Query" name:"IntervalForHistory"`
}

func (req *DescribeHistoryMonitorValuesRequest) Invoke(client *sdk.Client) (resp *DescribeHistoryMonitorValuesResponse, err error) {
	req.InitWithApiInfo("Ocs", "2015-03-01", "DescribeHistoryMonitorValues", "", "")
	resp = &DescribeHistoryMonitorValuesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeHistoryMonitorValuesResponse struct {
	responses.BaseResponse
	RequestId      common.String
	MonitorHistory common.String
}

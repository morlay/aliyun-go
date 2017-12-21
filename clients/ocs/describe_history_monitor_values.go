package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeHistoryMonitorValuesRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	EndTime              string `position:"Query" name:"EndTime"`
	IntervalForHistory   string `position:"Query" name:"IntervalForHistory"`
}

func (r DescribeHistoryMonitorValuesRequest) Invoke(client *sdk.Client) (response *DescribeHistoryMonitorValuesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeHistoryMonitorValuesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "DescribeHistoryMonitorValues", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeHistoryMonitorValuesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeHistoryMonitorValuesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeHistoryMonitorValuesResponse struct {
	RequestId      string
	MonitorHistory string
}

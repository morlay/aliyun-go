package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryDeviceStatisticsRequest struct {
	requests.RpcRequest
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *QueryDeviceStatisticsRequest) Invoke(client *sdk.Client) (resp *QueryDeviceStatisticsResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceStatistics", "", "")
	resp = &QueryDeviceStatisticsResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDeviceStatisticsResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
	Data         QueryDeviceStatisticsData
}

type QueryDeviceStatisticsData struct {
	DeviceCount common.Long
	OnlineCount common.Long
	ActiveCount common.Long
}

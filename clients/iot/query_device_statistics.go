package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         QueryDeviceStatisticsData
}

type QueryDeviceStatisticsData struct {
	DeviceCount int64
	OnlineCount int64
	ActiveCount int64
}

package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetAccountStatisticsRequest struct {
	requests.RpcRequest
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
}

func (req *GetAccountStatisticsRequest) Invoke(client *sdk.Client) (resp *GetAccountStatisticsResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "GetAccountStatistics", "vipaegis", "")
	resp = &GetAccountStatisticsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAccountStatisticsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Success   bool
	Message   common.String
	Data      GetAccountStatisticsData
}

type GetAccountStatisticsData struct {
	RemoteLogin  common.Integer
	CrackSuccess common.Integer
}

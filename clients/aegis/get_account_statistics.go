package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetAccountStatisticsRequest struct {
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
}

func (r GetAccountStatisticsRequest) Invoke(client *sdk.Client) (response *GetAccountStatisticsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetAccountStatisticsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("aegis", "2016-11-11", "GetAccountStatistics", "", "")

	resp := struct {
		*responses.BaseResponse
		GetAccountStatisticsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetAccountStatisticsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetAccountStatisticsResponse struct {
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      GetAccountStatisticsData
}

type GetAccountStatisticsData struct {
	RemoteLogin  int
	CrackSuccess int
}

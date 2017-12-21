package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetCrackStatisticsRequest struct {
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
}

func (r GetCrackStatisticsRequest) Invoke(client *sdk.Client) (response *GetCrackStatisticsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetCrackStatisticsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("aegis", "2016-11-11", "GetCrackStatistics", "", "")

	resp := struct {
		*responses.BaseResponse
		GetCrackStatisticsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetCrackStatisticsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetCrackStatisticsResponse struct {
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      GetCrackStatisticsData
}

type GetCrackStatisticsData struct {
	Intercepted int
}

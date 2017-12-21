package httpdns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetResolveStatisticsRequest struct {
	DomainName   string `position:"Query" name:"DomainName"`
	TimeSpan     int    `position:"Query" name:"TimeSpan"`
	ProtocolName string `position:"Query" name:"ProtocolName"`
	Granularity  string `position:"Query" name:"Granularity"`
}

func (r GetResolveStatisticsRequest) Invoke(client *sdk.Client) (response *GetResolveStatisticsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetResolveStatisticsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Httpdns", "2016-02-01", "GetResolveStatistics", "", "")

	resp := struct {
		*responses.BaseResponse
		GetResolveStatisticsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetResolveStatisticsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetResolveStatisticsResponse struct {
	RequestId  string
	DataPoints GetResolveStatisticsDataPointList
}

type GetResolveStatisticsDataPoint struct {
	Time  int
	Count int
}

type GetResolveStatisticsDataPointList []GetResolveStatisticsDataPoint

func (list *GetResolveStatisticsDataPointList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResolveStatisticsDataPoint)
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

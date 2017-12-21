package httpdns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetResolveStatisticsRequest struct {
	requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	TimeSpan     int    `position:"Query" name:"TimeSpan"`
	ProtocolName string `position:"Query" name:"ProtocolName"`
	Granularity  string `position:"Query" name:"Granularity"`
}

func (req *GetResolveStatisticsRequest) Invoke(client *sdk.Client) (resp *GetResolveStatisticsResponse, err error) {
	req.InitWithApiInfo("Httpdns", "2016-02-01", "GetResolveStatistics", "", "")
	resp = &GetResolveStatisticsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetResolveStatisticsResponse struct {
	responses.BaseResponse
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

package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCdnMonitorDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeCdnMonitorDataRequest) Invoke(client *sdk.Client) (response *DescribeCdnMonitorDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCdnMonitorDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnMonitorData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCdnMonitorDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCdnMonitorDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCdnMonitorDataResponse struct {
	RequestId       string
	DomainName      string
	MonitorInterval int64
	StartTime       string
	EndTime         string
	MonitorDatas    DescribeCdnMonitorDataCDNMonitorDataList
}

type DescribeCdnMonitorDataCDNMonitorData struct {
	TimeStamp         string
	QueryPerSecond    string
	BytesPerSecond    string
	BytesHitRate      string
	RequestHitRate    string
	AverageObjectSize string
}

type DescribeCdnMonitorDataCDNMonitorDataList []DescribeCdnMonitorDataCDNMonitorData

func (list *DescribeCdnMonitorDataCDNMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnMonitorDataCDNMonitorData)
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

package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCdnMonitorDataRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCdnMonitorDataRequest) Invoke(client *sdk.Client) (resp *DescribeCdnMonitorDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnMonitorData", "", "")
	resp = &DescribeCdnMonitorDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCdnMonitorDataResponse struct {
	responses.BaseResponse
	RequestId       common.String
	DomainName      common.String
	MonitorInterval common.Long
	StartTime       common.String
	EndTime         common.String
	MonitorDatas    DescribeCdnMonitorDataCDNMonitorDataList
}

type DescribeCdnMonitorDataCDNMonitorData struct {
	TimeStamp         common.String
	QueryPerSecond    common.String
	BytesPerSecond    common.String
	BytesHitRate      common.String
	RequestHitRate    common.String
	AverageObjectSize common.String
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

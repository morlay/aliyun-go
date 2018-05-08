package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainsUsageByDayRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainsUsageByDayRequest) Invoke(client *sdk.Client) (resp *DescribeDomainsUsageByDayResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainsUsageByDay", "", "")
	resp = &DescribeDomainsUsageByDayResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainsUsageByDayResponse struct {
	responses.BaseResponse
	RequestId    common.String
	DomainName   common.String
	DataInterval common.String
	StartTime    common.String
	EndTime      common.String
	UsageByDays  DescribeDomainsUsageByDayUsageByDayList
	UsageTotal   DescribeDomainsUsageByDayUsageTotal
}

type DescribeDomainsUsageByDayUsageByDay struct {
	TimeStamp      common.String
	Qps            common.String
	BytesHitRate   common.String
	RequestHitRate common.String
	MaxBps         common.String
	MaxBpsTime     common.String
	MaxSrcBps      common.String
	MaxSrcBpsTime  common.String
	TotalAccess    common.String
	TotalTraffic   common.String
}

type DescribeDomainsUsageByDayUsageTotal struct {
	BytesHitRate   common.String
	RequestHitRate common.String
	MaxBps         common.String
	MaxBpsTime     common.String
	MaxSrcBps      common.String
	MaxSrcBpsTime  common.String
	TotalAccess    common.String
	TotalTraffic   common.String
}

type DescribeDomainsUsageByDayUsageByDayList []DescribeDomainsUsageByDayUsageByDay

func (list *DescribeDomainsUsageByDayUsageByDayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsUsageByDayUsageByDay)
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

package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeExtensiveDomainDataRequest struct {
	requests.RpcRequest
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	ExtensiveDomain string `position:"Query" name:"ExtensiveDomain"`
	PageSize        int    `position:"Query" name:"PageSize"`
	EndTime         string `position:"Query" name:"EndTime"`
	StartTime       string `position:"Query" name:"StartTime"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeExtensiveDomainDataRequest) Invoke(client *sdk.Client) (resp *DescribeExtensiveDomainDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeExtensiveDomainData", "", "")
	resp = &DescribeExtensiveDomainDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeExtensiveDomainDataResponse struct {
	responses.BaseResponse
	RequestId       string
	ExtensiveDomain string
	DataInterval    string
	StartTime       string
	EndTime         string
	PageNumber      string
	TotalCount      string
	PageSize        string
	DataPerInterval DescribeExtensiveDomainDataUsageDataList
}

type DescribeExtensiveDomainDataUsageData struct {
	ExactDomain string
	TimeStamp   string
	Acc         string
	Flow        string
}

type DescribeExtensiveDomainDataUsageDataList []DescribeExtensiveDomainDataUsageData

func (list *DescribeExtensiveDomainDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeExtensiveDomainDataUsageData)
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

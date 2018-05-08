package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId       common.String
	ExtensiveDomain common.String
	DataInterval    common.String
	StartTime       common.String
	EndTime         common.String
	PageNumber      common.String
	TotalCount      common.String
	PageSize        common.String
	DataPerInterval DescribeExtensiveDomainDataUsageDataList
}

type DescribeExtensiveDomainDataUsageData struct {
	ExactDomain common.String
	TimeStamp   common.String
	Acc         common.String
	Flow        common.String
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

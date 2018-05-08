package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeTopDomainsByFlowRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Limit         int64  `position:"Query" name:"Limit"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeTopDomainsByFlowRequest) Invoke(client *sdk.Client) (resp *DescribeTopDomainsByFlowResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeTopDomainsByFlow", "", "")
	resp = &DescribeTopDomainsByFlowResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTopDomainsByFlowResponse struct {
	responses.BaseResponse
	RequestId         common.String
	StartTime         common.String
	EndTime           common.String
	DomainCount       common.Long
	DomainOnlineCount common.Long
	TopDomains        DescribeTopDomainsByFlowTopDomainList
}

type DescribeTopDomainsByFlowTopDomain struct {
	DomainName     common.String
	Rank           common.Long
	TotalTraffic   common.String
	TrafficPercent common.String
	MaxBps         common.Long
	MaxBpsTime     common.String
	TotalAccess    common.Long
}

type DescribeTopDomainsByFlowTopDomainList []DescribeTopDomainsByFlowTopDomain

func (list *DescribeTopDomainsByFlowTopDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTopDomainsByFlowTopDomain)
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

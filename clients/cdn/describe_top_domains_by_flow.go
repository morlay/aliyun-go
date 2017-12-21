package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTopDomainsByFlowRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Limit         int64  `position:"Query" name:"Limit"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeTopDomainsByFlowRequest) Invoke(client *sdk.Client) (response *DescribeTopDomainsByFlowResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeTopDomainsByFlowRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeTopDomainsByFlow", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeTopDomainsByFlowResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeTopDomainsByFlowResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeTopDomainsByFlowResponse struct {
	RequestId         string
	StartTime         string
	EndTime           string
	DomainCount       int64
	DomainOnlineCount int64
	TopDomains        DescribeTopDomainsByFlowTopDomainList
}

type DescribeTopDomainsByFlowTopDomain struct {
	DomainName     string
	Rank           int64
	TotalTraffic   string
	TrafficPercent string
	MaxBps         int64
	MaxBpsTime     string
	TotalAccess    int64
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

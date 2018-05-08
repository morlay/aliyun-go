package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainPathDataRequest struct {
	requests.RpcRequest
	StartTime     string `position:"Query" name:"StartTime"`
	PageNumber    int    `position:"Query" name:"PageNumber"`
	Path          string `position:"Query" name:"Path"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PageSize      int    `position:"Query" name:"PageSize"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *DescribeDomainPathDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainPathDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainPathData", "", "")
	resp = &DescribeDomainPathDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainPathDataResponse struct {
	responses.BaseResponse
	DomainName          common.String
	StartTime           common.String
	EndTime             common.String
	PageSize            common.Integer
	PageNumber          common.Integer
	DataInterval        common.String
	TotalCount          common.Integer
	PathDataPerInterval DescribeDomainPathDataUsageDataList
}

type DescribeDomainPathDataUsageData struct {
	Traffic common.Integer
	Acc     common.Integer
	Path    common.String
	Time    common.String
}

type DescribeDomainPathDataUsageDataList []DescribeDomainPathDataUsageData

func (list *DescribeDomainPathDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainPathDataUsageData)
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

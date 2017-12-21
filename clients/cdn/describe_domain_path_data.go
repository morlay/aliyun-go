package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	DomainName          string
	StartTime           string
	EndTime             string
	PageSize            int
	PageNumber          int
	DataInterval        string
	TotalCount          int
	PathDataPerInterval DescribeDomainPathDataUsageDataList
}

type DescribeDomainPathDataUsageData struct {
	Traffic int
	Acc     int
	Path    string
	Time    string
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

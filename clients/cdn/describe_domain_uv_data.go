package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainUvDataRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainUvDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainUvDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainUvData", "", "")
	resp = &DescribeDomainUvDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainUvDataResponse struct {
	responses.BaseResponse
	RequestId      common.String
	DomainName     common.String
	DataInterval   common.String
	StartTime      common.String
	EndTime        common.String
	UvDataInterval DescribeDomainUvDataUsageDataList
}

type DescribeDomainUvDataUsageData struct {
	TimeStamp common.String
	Value     common.String
}

type DescribeDomainUvDataUsageDataList []DescribeDomainUvDataUsageData

func (list *DescribeDomainUvDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainUvDataUsageData)
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

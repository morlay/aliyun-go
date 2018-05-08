package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainSrcBpsDataRequest struct {
	requests.RpcRequest
	StartTime  string `position:"Query" name:"StartTime"`
	FixTimeGap string `position:"Query" name:"FixTimeGap"`
	TimeMerge  string `position:"Query" name:"TimeMerge"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	Interval   string `position:"Query" name:"Interval"`
}

func (req *DescribeDomainSrcBpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainSrcBpsDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainSrcBpsData", "", "")
	resp = &DescribeDomainSrcBpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainSrcBpsDataResponse struct {
	responses.BaseResponse
	RequestId             common.String
	DomainName            common.String
	StartTime             common.String
	EndTime               common.String
	DataInterval          common.String
	SrcBpsDataPerInterval DescribeDomainSrcBpsDataDataModuleList
}

type DescribeDomainSrcBpsDataDataModule struct {
	TimeStamp common.String
	Value     common.String
}

type DescribeDomainSrcBpsDataDataModuleList []DescribeDomainSrcBpsDataDataModule

func (list *DescribeDomainSrcBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainSrcBpsDataDataModule)
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

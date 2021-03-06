package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainSrcFlowDataRequest struct {
	requests.RpcRequest
	StartTime  string `position:"Query" name:"StartTime"`
	FixTimeGap string `position:"Query" name:"FixTimeGap"`
	TimeMerge  string `position:"Query" name:"TimeMerge"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	Interval   string `position:"Query" name:"Interval"`
}

func (req *DescribeDomainSrcFlowDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainSrcFlowDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainSrcFlowData", "", "")
	resp = &DescribeDomainSrcFlowDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainSrcFlowDataResponse struct {
	responses.BaseResponse
	RequestId              common.String
	DomainName             common.String
	StartTime              common.String
	EndTime                common.String
	DataInterval           common.String
	SrcFlowDataPerInterval DescribeDomainSrcFlowDataDataModuleList
}

type DescribeDomainSrcFlowDataDataModule struct {
	TimeStamp common.String
	Value     common.String
}

type DescribeDomainSrcFlowDataDataModuleList []DescribeDomainSrcFlowDataDataModule

func (list *DescribeDomainSrcFlowDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainSrcFlowDataDataModule)
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

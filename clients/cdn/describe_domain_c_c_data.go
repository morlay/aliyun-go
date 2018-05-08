package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainCCDataRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainCCDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainCCDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainCCData", "", "")
	resp = &DescribeDomainCCDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainCCDataResponse struct {
	responses.BaseResponse
	RequestId    common.String
	DomainName   common.String
	DataInterval common.String
	StartTime    common.String
	EndTime      common.String
	CCDataList   DescribeDomainCCDataCCDatasList
}

type DescribeDomainCCDataCCDatas struct {
	TimeStamp common.String
	Count     common.String
}

type DescribeDomainCCDataCCDatasList []DescribeDomainCCDataCCDatas

func (list *DescribeDomainCCDataCCDatasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCCDataCCDatas)
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

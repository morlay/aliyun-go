package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId    string
	DomainName   string
	DataInterval string
	StartTime    string
	EndTime      string
	CCDataList   DescribeDomainCCDataCCDatasList
}

type DescribeDomainCCDataCCDatas struct {
	TimeStamp string
	Count     string
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

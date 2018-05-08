package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainRealTimeDataRequest struct {
	requests.RpcRequest
	Field         string `position:"Query" name:"Field"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainRealTimeDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainRealTimeDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainRealTimeData", "", "")
	resp = &DescribeDomainRealTimeDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainRealTimeDataResponse struct {
	responses.BaseResponse
	RequestId       common.String
	DomainName      common.String
	Field           common.String
	StartTime       common.String
	EndTime         common.String
	DataPerInterval DescribeDomainRealTimeDataDataModuleList
}

type DescribeDomainRealTimeDataDataModule struct {
	TimeStamp common.String
	Value     common.String
}

type DescribeDomainRealTimeDataDataModuleList []DescribeDomainRealTimeDataDataModule

func (list *DescribeDomainRealTimeDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRealTimeDataDataModule)
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

package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type BatchDescribeDomainBpsDataRequest struct {
	requests.RpcRequest
	StartTime  string `position:"Query" name:"StartTime"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *BatchDescribeDomainBpsDataRequest) Invoke(client *sdk.Client) (resp *BatchDescribeDomainBpsDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "BatchDescribeDomainBpsData", "", "")
	resp = &BatchDescribeDomainBpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type BatchDescribeDomainBpsDataResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageNumber common.Integer
	PageSize   common.Integer
	TotalCount common.Integer
	BpsDatas   BatchDescribeDomainBpsDataDataModuleList
}

type BatchDescribeDomainBpsDataDataModule struct {
	Timestamp  common.String
	L1Bps      common.Float
	L1InnerBps common.Float
	L1OutBps   common.Float
	DomainName common.String
}

type BatchDescribeDomainBpsDataDataModuleList []BatchDescribeDomainBpsDataDataModule

func (list *BatchDescribeDomainBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]BatchDescribeDomainBpsDataDataModule)
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

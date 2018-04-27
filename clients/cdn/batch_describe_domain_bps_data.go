package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	PageNumber int
	PageSize   int
	TotalCount int
	BpsDatas   BatchDescribeDomainBpsDataDataModuleList
}

type BatchDescribeDomainBpsDataDataModule struct {
	Timestamp  string
	L1Bps      float32
	L1InnerBps float32
	L1OutBps   float32
	DomainName string
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

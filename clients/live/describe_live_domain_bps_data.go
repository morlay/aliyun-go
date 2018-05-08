package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveDomainBpsDataRequest struct {
	requests.RpcRequest
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	Interval       string `position:"Query" name:"Interval"`
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveDomainBpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeLiveDomainBpsDataResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainBpsData", "live", "")
	resp = &DescribeLiveDomainBpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveDomainBpsDataResponse struct {
	responses.BaseResponse
	RequestId          common.String
	DomainName         common.String
	StartTime          common.String
	EndTime            common.String
	DataInterval       common.String
	BpsDataPerInterval DescribeLiveDomainBpsDataDataModuleList
}

type DescribeLiveDomainBpsDataDataModule struct {
	TimeStamp     common.String
	BpsValue      common.String
	HttpBpsValue  common.String
	HttpsBpsValue common.String
}

type DescribeLiveDomainBpsDataDataModuleList []DescribeLiveDomainBpsDataDataModule

func (list *DescribeLiveDomainBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveDomainBpsDataDataModule)
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

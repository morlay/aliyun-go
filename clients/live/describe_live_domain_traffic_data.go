package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveDomainTrafficDataRequest struct {
	requests.RpcRequest
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	Interval       string `position:"Query" name:"Interval"`
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveDomainTrafficDataRequest) Invoke(client *sdk.Client) (resp *DescribeLiveDomainTrafficDataResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainTrafficData", "live", "")
	resp = &DescribeLiveDomainTrafficDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveDomainTrafficDataResponse struct {
	responses.BaseResponse
	RequestId              common.String
	DomainName             common.String
	StartTime              common.String
	EndTime                common.String
	DataInterval           common.String
	TrafficDataPerInterval DescribeLiveDomainTrafficDataDataModuleList
}

type DescribeLiveDomainTrafficDataDataModule struct {
	TimeStamp         common.String
	TrafficValue      common.String
	HttpTrafficValue  common.String
	HttpsTrafficValue common.String
}

type DescribeLiveDomainTrafficDataDataModuleList []DescribeLiveDomainTrafficDataDataModule

func (list *DescribeLiveDomainTrafficDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveDomainTrafficDataDataModule)
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

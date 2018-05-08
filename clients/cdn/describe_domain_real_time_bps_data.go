package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainRealTimeBpsDataRequest struct {
	requests.RpcRequest
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainRealTimeBpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainRealTimeBpsDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainRealTimeBpsData", "", "")
	resp = &DescribeDomainRealTimeBpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainRealTimeBpsDataResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      DescribeDomainRealTimeBpsDataBpsModelList
}

type DescribeDomainRealTimeBpsDataBpsModel struct {
	Bps       common.Float
	TimeStamp common.String
}

type DescribeDomainRealTimeBpsDataBpsModelList []DescribeDomainRealTimeBpsDataBpsModel

func (list *DescribeDomainRealTimeBpsDataBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRealTimeBpsDataBpsModel)
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

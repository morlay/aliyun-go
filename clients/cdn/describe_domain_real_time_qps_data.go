package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainRealTimeQpsDataRequest struct {
	requests.RpcRequest
	LocationNameEn string `position:"Query" name:"LocationNameEn"`
	IspNameEn      string `position:"Query" name:"IspNameEn"`
	StartTime      string `position:"Query" name:"StartTime"`
	DomainName     string `position:"Query" name:"DomainName"`
	EndTime        string `position:"Query" name:"EndTime"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainRealTimeQpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainRealTimeQpsDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainRealTimeQpsData", "", "")
	resp = &DescribeDomainRealTimeQpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainRealTimeQpsDataResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      DescribeDomainRealTimeQpsDataQpsModelList
}

type DescribeDomainRealTimeQpsDataQpsModel struct {
	Qps       common.Float
	TimeStamp common.String
}

type DescribeDomainRealTimeQpsDataQpsModelList []DescribeDomainRealTimeQpsDataQpsModel

func (list *DescribeDomainRealTimeQpsDataQpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRealTimeQpsDataQpsModel)
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

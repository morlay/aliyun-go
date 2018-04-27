package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainRealTimeByteHitRateDataRequest struct {
	requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainRealTimeByteHitRateDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainRealTimeByteHitRateDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainRealTimeByteHitRateData", "", "")
	resp = &DescribeDomainRealTimeByteHitRateDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainRealTimeByteHitRateDataResponse struct {
	responses.BaseResponse
	RequestId string
	Data      DescribeDomainRealTimeByteHitRateDataByteHitRateDataModelList
}

type DescribeDomainRealTimeByteHitRateDataByteHitRateDataModel struct {
	ByteHitRate float32
	TimeStamp   string
}

type DescribeDomainRealTimeByteHitRateDataByteHitRateDataModelList []DescribeDomainRealTimeByteHitRateDataByteHitRateDataModel

func (list *DescribeDomainRealTimeByteHitRateDataByteHitRateDataModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRealTimeByteHitRateDataByteHitRateDataModel)
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

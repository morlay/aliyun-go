package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainRealTimeReqHitRateDataRequest struct {
	requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainRealTimeReqHitRateDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainRealTimeReqHitRateDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainRealTimeReqHitRateData", "", "")
	resp = &DescribeDomainRealTimeReqHitRateDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainRealTimeReqHitRateDataResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      DescribeDomainRealTimeReqHitRateDataReqHitRateDataModelList
}

type DescribeDomainRealTimeReqHitRateDataReqHitRateDataModel struct {
	ReqHitRate common.Float
	TimeStamp  common.String
}

type DescribeDomainRealTimeReqHitRateDataReqHitRateDataModelList []DescribeDomainRealTimeReqHitRateDataReqHitRateDataModel

func (list *DescribeDomainRealTimeReqHitRateDataReqHitRateDataModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRealTimeReqHitRateDataReqHitRateDataModel)
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

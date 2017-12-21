package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainHttpsDataRequest struct {
	requests.RpcRequest
	DomainType    string `position:"Query" name:"DomainType"`
	FixTimeGap    string `position:"Query" name:"FixTimeGap"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	TimeMerge     string `position:"Query" name:"TimeMerge"`
	DomainNames   string `position:"Query" name:"DomainNames"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	Cls           string `position:"Query" name:"Cls"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainHttpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeDomainHttpsDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainHttpsData", "", "")
	resp = &DescribeDomainHttpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainHttpsDataResponse struct {
	responses.BaseResponse
	RequestId            string
	DomainNames          string
	DataInterval         string
	HttpsStatisticsInfos DescribeDomainHttpsDataHttpsStatisticsInfoList
}

type DescribeDomainHttpsDataHttpsStatisticsInfo struct {
	Time               string
	L1HttpsBps         float32
	L1HttpsInnerBps    float32
	L1HttpsOutBps      float32
	L1HttpsQps         int64
	L1HttpsInnerQps    int64
	L1HttpsOutQps      int64
	L1HttpsTtraf       int64
	HttpsSrcBps        int64
	HttpsSrcTraf       int64
	L1HttpsInnerTraf   int64
	L1HttpsOutTraf     int64
	HttpsByteHitRate   float32
	HttpsReqHitRate    float32
	L1HttpsHitRate     float32
	L1HttpsInner_acc   float32
	L1HttpsOut_acc     float32
	L1HttpsTacc        float32
	L1DyHttpsBps       float32
	L1DyHttpsInnerBps  float32
	L1DyHttpsOutBps    float32
	L1StHttpsBps       float32
	L1StHttpsInnerBps  float32
	L1StHttpsOutBps    float32
	L1DyHttpsTraf      float32
	L1DyHttpsInnerTraf float32
	L1DyHttpsOutTraf   float32
	L1StHttpsTraf      float32
	L1StHttpsInnerTraf float32
	L1StHttpsOutTraf   float32
	L1DyHttpsQps       float32
	L1DyHttpsInnerQps  float32
	L1DyHttpsOutQps    float32
	L1StHttpsQps       float32
	L1StHttpsInnerQps  float32
	L1StHttpsOutQps    float32
	L1DyHttpsAcc       float32
	L1DyHttpsInnerAcc  float32
	L1DyHttpsOutAcc    float32
	L1StHttpsAcc       float32
	L1StHttpsInnerAcc  float32
	L1StHttpsOutAcc    float32
}

type DescribeDomainHttpsDataHttpsStatisticsInfoList []DescribeDomainHttpsDataHttpsStatisticsInfo

func (list *DescribeDomainHttpsDataHttpsStatisticsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainHttpsDataHttpsStatisticsInfo)
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

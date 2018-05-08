package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId            common.String
	DomainNames          common.String
	DataInterval         common.String
	HttpsStatisticsInfos DescribeDomainHttpsDataHttpsStatisticsInfoList
}

type DescribeDomainHttpsDataHttpsStatisticsInfo struct {
	Time               common.String
	L1HttpsBps         common.Float
	L1HttpsInnerBps    common.Float
	L1HttpsOutBps      common.Float
	L1HttpsQps         common.Long
	L1HttpsInnerQps    common.Long
	L1HttpsOutQps      common.Long
	L1HttpsTtraf       common.Long
	HttpsSrcBps        common.Long
	HttpsSrcTraf       common.Long
	L1HttpsInnerTraf   common.Long
	L1HttpsOutTraf     common.Long
	HttpsByteHitRate   common.Float
	HttpsReqHitRate    common.Float
	L1HttpsHitRate     common.Float
	L1HttpsInner_acc   common.Float
	L1HttpsOut_acc     common.Float
	L1HttpsTacc        common.Float
	L1DyHttpsBps       common.Float
	L1DyHttpsInnerBps  common.Float
	L1DyHttpsOutBps    common.Float
	L1StHttpsBps       common.Float
	L1StHttpsInnerBps  common.Float
	L1StHttpsOutBps    common.Float
	L1DyHttpsTraf      common.Float
	L1DyHttpsInnerTraf common.Float
	L1DyHttpsOutTraf   common.Float
	L1StHttpsTraf      common.Float
	L1StHttpsInnerTraf common.Float
	L1StHttpsOutTraf   common.Float
	L1DyHttpsQps       common.Float
	L1DyHttpsInnerQps  common.Float
	L1DyHttpsOutQps    common.Float
	L1StHttpsQps       common.Float
	L1StHttpsInnerQps  common.Float
	L1StHttpsOutQps    common.Float
	L1DyHttpsAcc       common.Float
	L1DyHttpsInnerAcc  common.Float
	L1DyHttpsOutAcc    common.Float
	L1StHttpsAcc       common.Float
	L1StHttpsInnerAcc  common.Float
	L1StHttpsOutAcc    common.Float
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

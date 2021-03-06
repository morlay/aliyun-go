package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamBitRateDataRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamBitRateDataRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamBitRateDataResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamBitRateData", "live", "")
	resp = &DescribeLiveStreamBitRateDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamBitRateDataResponse struct {
	responses.BaseResponse
	RequestId                common.String
	FrameRateAndBitRateInfos DescribeLiveStreamBitRateDataFrameRateAndBitRateInfoList
}

type DescribeLiveStreamBitRateDataFrameRateAndBitRateInfo struct {
	StreamUrl      common.String
	VideoFrameRate common.Float
	AudioFrameRate common.Float
	BitRate        common.Float
	Time           common.String
}

type DescribeLiveStreamBitRateDataFrameRateAndBitRateInfoList []DescribeLiveStreamBitRateDataFrameRateAndBitRateInfo

func (list *DescribeLiveStreamBitRateDataFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamBitRateDataFrameRateAndBitRateInfo)
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

package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamsFrameRateAndBitRateDataRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamsFrameRateAndBitRateDataRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamsFrameRateAndBitRateDataResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamsFrameRateAndBitRateData", "live", "")
	resp = &DescribeLiveStreamsFrameRateAndBitRateDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamsFrameRateAndBitRateDataResponse struct {
	responses.BaseResponse
	RequestId                common.String
	FrameRateAndBitRateInfos DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfoList
}

type DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfo struct {
	StreamUrl      common.String
	VideoFrameRate common.Float
	AudioFrameRate common.Float
	BitRate        common.Float
	Time           common.String
}

type DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfoList []DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfo

func (list *DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfo)
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

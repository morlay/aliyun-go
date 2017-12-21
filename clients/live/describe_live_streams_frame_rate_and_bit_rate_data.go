package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamsFrameRateAndBitRateDataRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r DescribeLiveStreamsFrameRateAndBitRateDataRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamsFrameRateAndBitRateDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamsFrameRateAndBitRateDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamsFrameRateAndBitRateData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamsFrameRateAndBitRateDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamsFrameRateAndBitRateDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamsFrameRateAndBitRateDataResponse struct {
	RequestId                string
	FrameRateAndBitRateInfos DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfoList
}

type DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfo struct {
	StreamUrl      string
	VideoFrameRate float32
	AudioFrameRate float32
	BitRate        float32
	Time           string
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

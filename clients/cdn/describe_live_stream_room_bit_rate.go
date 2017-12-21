package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamRoomBitRateRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamRoomBitRateRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamRoomBitRateResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRoomBitRate", "", "")
	resp = &DescribeLiveStreamRoomBitRateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamRoomBitRateResponse struct {
	responses.BaseResponse
	RequestId                string
	FrameRateAndBitRateInfos DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfoList
}

type DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfo struct {
	StreamUrl      string
	VideoFrameRate float32
	AudioFrameRate float32
	BitRate        float32
	Time           string
}

type DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfoList []DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfo

func (list *DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfo)
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

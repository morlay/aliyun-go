package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamFrameLossRatioRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamFrameLossRatioRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamFrameLossRatioResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamFrameLossRatio", "", "")
	resp = &DescribeLiveStreamFrameLossRatioResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamFrameLossRatioResponse struct {
	responses.BaseResponse
	RequestId           string
	FrameLossRatioInfos DescribeLiveStreamFrameLossRatioFrameLossRatioInfoList
}

type DescribeLiveStreamFrameLossRatioFrameLossRatioInfo struct {
	StreamUrl      string
	FrameLossRatio float32
	Time           string
}

type DescribeLiveStreamFrameLossRatioFrameLossRatioInfoList []DescribeLiveStreamFrameLossRatioFrameLossRatioInfo

func (list *DescribeLiveStreamFrameLossRatioFrameLossRatioInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamFrameLossRatioFrameLossRatioInfo)
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

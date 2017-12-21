package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamPushDataRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveStreamPushDataRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamPushDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamPushData", "", "")
	resp = &DescribeLiveStreamPushDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamPushDataResponse struct {
	responses.BaseResponse
	RequestId           string
	PushStreamModelList DescribeLiveStreamPushDataPushStreamModelList
}

type DescribeLiveStreamPushDataPushStreamModel struct {
	Time          string
	Stream        string
	FrameRate     float32
	BitRate       float32
	FrameLossRate float32
	ServerAddr    string
	ClientAddr    string
}

type DescribeLiveStreamPushDataPushStreamModelList []DescribeLiveStreamPushDataPushStreamModel

func (list *DescribeLiveStreamPushDataPushStreamModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamPushDataPushStreamModel)
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

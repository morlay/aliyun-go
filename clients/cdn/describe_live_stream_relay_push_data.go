package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamRelayPushDataRequest struct {
	RelayDomain   string `position:"Query" name:"RelayDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveStreamRelayPushDataRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamRelayPushDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamRelayPushDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRelayPushData", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamRelayPushDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamRelayPushDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamRelayPushDataResponse struct {
	RequestId                string
	RelayPushDetailModelList DescribeLiveStreamRelayPushDataRelayPushDetailModelList
}

type DescribeLiveStreamRelayPushDataRelayPushDetailModel struct {
	Time          string
	Stream        string
	FrameRate     float32
	BitRate       float32
	FrameLossRate float32
	ServerAddr    string
	ClientAddr    string
}

type DescribeLiveStreamRelayPushDataRelayPushDetailModelList []DescribeLiveStreamRelayPushDataRelayPushDetailModel

func (list *DescribeLiveStreamRelayPushDataRelayPushDetailModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRelayPushDataRelayPushDetailModel)
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

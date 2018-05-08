package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamRelayPushDataRequest struct {
	requests.RpcRequest
	RelayDomain   string `position:"Query" name:"RelayDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveStreamRelayPushDataRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamRelayPushDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRelayPushData", "", "")
	resp = &DescribeLiveStreamRelayPushDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamRelayPushDataResponse struct {
	responses.BaseResponse
	RequestId                common.String
	RelayPushDetailModelList DescribeLiveStreamRelayPushDataRelayPushDetailModelList
}

type DescribeLiveStreamRelayPushDataRelayPushDetailModel struct {
	Time          common.String
	Stream        common.String
	FrameRate     common.Float
	BitRate       common.Float
	FrameLossRate common.Float
	ServerAddr    common.String
	ClientAddr    common.String
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

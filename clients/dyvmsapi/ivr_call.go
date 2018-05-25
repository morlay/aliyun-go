package dyvmsapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type IvrCallRequest struct {
	requests.RpcRequest
	ByeCode              string                 `position:"Query" name:"ByeCode"`
	MenuKeyMaps          *IvrCallMenuKeyMapList `position:"Query" type:"Repeated" name:"MenuKeyMap"`
	ResourceOwnerId      int64                  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                 `position:"Query" name:"ResourceOwnerAccount"`
	StartTtsParams       string                 `position:"Query" name:"StartTtsParams"`
	PlayTimes            int64                  `position:"Query" name:"PlayTimes"`
	OwnerId              int64                  `position:"Query" name:"OwnerId"`
	Timeout              int                    `position:"Query" name:"Timeout"`
	StartCode            string                 `position:"Query" name:"StartCode"`
	CalledNumber         string                 `position:"Query" name:"CalledNumber"`
	CalledShowNumber     string                 `position:"Query" name:"CalledShowNumber"`
	OutId                string                 `position:"Query" name:"OutId"`
	ByeTtsParams         string                 `position:"Query" name:"ByeTtsParams"`
}

func (req *IvrCallRequest) Invoke(client *sdk.Client) (resp *IvrCallResponse, err error) {
	req.InitWithApiInfo("Dyvmsapi", "2017-05-25", "IvrCall", "", "")
	resp = &IvrCallResponse{}
	err = client.DoAction(req, resp)
	return
}

type IvrCallMenuKeyMap struct {
	Key       string `name:"Key"`
	Code      string `name:"Code"`
	TtsParams string `name:"TtsParams"`
}

type IvrCallResponse struct {
	responses.BaseResponse
	RequestId common.String
	CallId    common.String
	Code      common.String
	Message   common.String
}

type IvrCallMenuKeyMapList []IvrCallMenuKeyMap

func (list *IvrCallMenuKeyMapList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IvrCallMenuKeyMap)
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

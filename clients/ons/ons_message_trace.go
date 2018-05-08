package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsMessageTraceRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	MsgId        string `position:"Query" name:"MsgId"`
}

func (req *OnsMessageTraceRequest) Invoke(client *sdk.Client) (resp *OnsMessageTraceResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMessageTrace", "", "")
	resp = &OnsMessageTraceResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMessageTraceResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
	Data      OnsMessageTraceMessageTrackList
}

type OnsMessageTraceMessageTrack struct {
	ConsumerGroup common.String
	TrackType     common.String
	ExceptionDesc common.String
}

type OnsMessageTraceMessageTrackList []OnsMessageTraceMessageTrack

func (list *OnsMessageTraceMessageTrackList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageTraceMessageTrack)
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

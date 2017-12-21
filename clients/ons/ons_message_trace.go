package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMessageTraceRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	MsgId        string `position:"Query" name:"MsgId"`
}

func (r OnsMessageTraceRequest) Invoke(client *sdk.Client) (response *OnsMessageTraceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMessageTraceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMessageTrace", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMessageTraceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMessageTraceResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMessageTraceResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsMessageTraceMessageTrackList
}

type OnsMessageTraceMessageTrack struct {
	ConsumerGroup string
	TrackType     string
	ExceptionDesc string
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

package dyvmsapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SmartCallRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ActionCodeBreak      string `position:"Query" name:"ActionCodeBreak"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RecordFlag           string `position:"Query" name:"RecordFlag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Speed                int    `position:"Query" name:"Speed"`
	Volume               int    `position:"Query" name:"Volume"`
	DynamicId            string `position:"Query" name:"DynamicId"`
	CalledNumber         string `position:"Query" name:"CalledNumber"`
	VoiceCode            string `position:"Query" name:"VoiceCode"`
	MuteTime             int    `position:"Query" name:"MuteTime"`
	CalledShowNumber     string `position:"Query" name:"CalledShowNumber"`
	OutId                string `position:"Query" name:"OutId"`
	AsrModelId           string `position:"Query" name:"AsrModelId"`
	PauseTime            int    `position:"Query" name:"PauseTime"`
}

func (req *SmartCallRequest) Invoke(client *sdk.Client) (resp *SmartCallResponse, err error) {
	req.InitWithApiInfo("Dyvmsapi", "2017-05-25", "SmartCall", "", "")
	resp = &SmartCallResponse{}
	err = client.DoAction(req, resp)
	return
}

type SmartCallResponse struct {
	responses.BaseResponse
	RequestId common.String
	CallId    common.String
	Code      common.String
	Message   common.String
}

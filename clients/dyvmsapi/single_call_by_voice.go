package dyvmsapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SingleCallByVoiceRequest struct {
	requests.RpcRequest
	Volume               int    `position:"Query" name:"Volume"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CalledNumber         string `position:"Query" name:"CalledNumber"`
	VoiceCode            string `position:"Query" name:"VoiceCode"`
	CalledShowNumber     string `position:"Query" name:"CalledShowNumber"`
	PlayTimes            int    `position:"Query" name:"PlayTimes"`
	OutId                string `position:"Query" name:"OutId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Speed                int    `position:"Query" name:"Speed"`
}

func (req *SingleCallByVoiceRequest) Invoke(client *sdk.Client) (resp *SingleCallByVoiceResponse, err error) {
	req.InitWithApiInfo("Dyvmsapi", "2017-05-25", "SingleCallByVoice", "", "")
	resp = &SingleCallByVoiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type SingleCallByVoiceResponse struct {
	responses.BaseResponse
	RequestId common.String
	CallId    common.String
	Code      common.String
	Message   common.String
}

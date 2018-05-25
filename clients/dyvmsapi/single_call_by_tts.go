package dyvmsapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SingleCallByTtsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	TtsCode              string `position:"Query" name:"TtsCode"`
	PlayTimes            int    `position:"Query" name:"PlayTimes"`
	TtsParam             string `position:"Query" name:"TtsParam"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Speed                int    `position:"Query" name:"Speed"`
	Volume               int    `position:"Query" name:"Volume"`
	CalledNumber         string `position:"Query" name:"CalledNumber"`
	CalledShowNumber     string `position:"Query" name:"CalledShowNumber"`
	OutId                string `position:"Query" name:"OutId"`
}

func (req *SingleCallByTtsRequest) Invoke(client *sdk.Client) (resp *SingleCallByTtsResponse, err error) {
	req.InitWithApiInfo("Dyvmsapi", "2017-05-25", "SingleCallByTts", "", "")
	resp = &SingleCallByTtsResponse{}
	err = client.DoAction(req, resp)
	return
}

type SingleCallByTtsResponse struct {
	responses.BaseResponse
	RequestId common.String
	CallId    common.String
	Code      common.String
	Message   common.String
}

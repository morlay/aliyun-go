package dyvmsapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ClickToDialRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RecordFlag           string `position:"Query" name:"RecordFlag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CallerShowNumber     string `position:"Query" name:"CallerShowNumber"`
	SessionTimeout       int    `position:"Query" name:"SessionTimeout"`
	CalledNumber         string `position:"Query" name:"CalledNumber"`
	CalledShowNumber     string `position:"Query" name:"CalledShowNumber"`
	OutId                string `position:"Query" name:"OutId"`
	AsrFlag              string `position:"Query" name:"AsrFlag"`
	AsrModelId           string `position:"Query" name:"AsrModelId"`
	CallerNumber         string `position:"Query" name:"CallerNumber"`
}

func (req *ClickToDialRequest) Invoke(client *sdk.Client) (resp *ClickToDialResponse, err error) {
	req.InitWithApiInfo("Dyvmsapi", "2017-05-25", "ClickToDial", "", "")
	resp = &ClickToDialResponse{}
	err = client.DoAction(req, resp)
	return
}

type ClickToDialResponse struct {
	responses.BaseResponse
	RequestId common.String
	CallId    common.String
	Code      common.String
	Message   common.String
}

package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetectPornRequest struct {
	requests.RpcRequest
	SrcType string `position:"Query" name:"SrcType"`
	Engine  string `position:"Query" name:"Engine"`
	Project string `position:"Query" name:"Project"`
	Url     string `position:"Query" name:"Url"`
}

func (req *DetectPornRequest) Invoke(client *sdk.Client) (resp *DetectPornResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DetectPorn", "imm", "")
	resp = &DetectPornResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetectPornResponse struct {
	responses.BaseResponse
	RequestId string
	Score     float32
}

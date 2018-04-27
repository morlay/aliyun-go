package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopParagraphRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *StopParagraphRequest) Invoke(client *sdk.Client) (resp *StopParagraphResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "StopParagraph", "", "")
	resp = &StopParagraphResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopParagraphResponse struct {
	responses.BaseResponse
	RequestId string
}

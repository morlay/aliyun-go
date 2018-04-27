package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateParagraphRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	Text            string `position:"Query" name:"Text"`
}

func (req *CreateParagraphRequest) Invoke(client *sdk.Client) (resp *CreateParagraphResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateParagraph", "", "")
	resp = &CreateParagraphResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateParagraphResponse struct {
	responses.BaseResponse
	RequestId string
	Id        string
}

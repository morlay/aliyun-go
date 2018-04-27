package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveParagraphRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	Id              string `position:"Query" name:"Id"`
	Text            string `position:"Query" name:"Text"`
}

func (req *SaveParagraphRequest) Invoke(client *sdk.Client) (resp *SaveParagraphResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "SaveParagraph", "", "")
	resp = &SaveParagraphResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveParagraphResponse struct {
	responses.BaseResponse
	RequestId string
}

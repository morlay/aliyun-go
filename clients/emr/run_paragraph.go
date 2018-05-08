package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RunParagraphRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	Id              string `position:"Query" name:"Id"`
	Text            string `position:"Query" name:"Text"`
}

func (req *RunParagraphRequest) Invoke(client *sdk.Client) (resp *RunParagraphResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RunParagraph", "", "")
	resp = &RunParagraphResponse{}
	err = client.DoAction(req, resp)
	return
}

type RunParagraphResponse struct {
	responses.BaseResponse
	RequestId common.String
}

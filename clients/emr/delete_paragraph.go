package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteParagraphRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DeleteParagraphRequest) Invoke(client *sdk.Client) (resp *DeleteParagraphResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteParagraph", "", "")
	resp = &DeleteParagraphResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteParagraphResponse struct {
	responses.BaseResponse
	RequestId string
}

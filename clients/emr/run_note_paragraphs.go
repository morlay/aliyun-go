package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RunNoteParagraphsRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
}

func (req *RunNoteParagraphsRequest) Invoke(client *sdk.Client) (resp *RunNoteParagraphsResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RunNoteParagraphs", "", "")
	resp = &RunNoteParagraphsResponse{}
	err = client.DoAction(req, resp)
	return
}

type RunNoteParagraphsResponse struct {
	responses.BaseResponse
	RequestId string
}

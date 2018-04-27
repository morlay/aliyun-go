package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteNoteRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DeleteNoteRequest) Invoke(client *sdk.Client) (resp *DeleteNoteResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteNote", "", "")
	resp = &DeleteNoteResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteNoteResponse struct {
	responses.BaseResponse
	RequestId string
}

package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateNoteRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	Type            string `position:"Query" name:"Type"`
}

func (req *CreateNoteRequest) Invoke(client *sdk.Client) (resp *CreateNoteResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateNote", "", "")
	resp = &CreateNoteResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateNoteResponse struct {
	responses.BaseResponse
	RequestId common.String
	Id        common.String
	Paragraph common.String
}

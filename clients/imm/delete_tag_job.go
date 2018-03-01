package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteTagJobRequest struct {
	requests.RpcRequest
	JobId          string `position:"Query" name:"JobId"`
	Project        string `position:"Query" name:"Project"`
	ClearIndexData string `position:"Query" name:"ClearIndexData"`
}

func (req *DeleteTagJobRequest) Invoke(client *sdk.Client) (resp *DeleteTagJobResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DeleteTagJob", "imm", "")
	resp = &DeleteTagJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteTagJobResponse struct {
	responses.BaseResponse
	RequestId string
}

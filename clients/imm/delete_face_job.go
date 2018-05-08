package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteFaceJobRequest struct {
	requests.RpcRequest
	JobId          string `position:"Query" name:"JobId"`
	Project        string `position:"Query" name:"Project"`
	ClearIndexData string `position:"Query" name:"ClearIndexData"`
}

func (req *DeleteFaceJobRequest) Invoke(client *sdk.Client) (resp *DeleteFaceJobResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DeleteFaceJob", "imm", "")
	resp = &DeleteFaceJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFaceJobResponse struct {
	responses.BaseResponse
	RequestId common.String
}

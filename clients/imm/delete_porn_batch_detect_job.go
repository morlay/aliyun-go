package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeletePornBatchDetectJobRequest struct {
	requests.RpcRequest
	JobId   string `position:"Query" name:"JobId"`
	Project string `position:"Query" name:"Project"`
}

func (req *DeletePornBatchDetectJobRequest) Invoke(client *sdk.Client) (resp *DeletePornBatchDetectJobResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DeletePornBatchDetectJob", "imm", "")
	resp = &DeletePornBatchDetectJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeletePornBatchDetectJobResponse struct {
	responses.BaseResponse
	RequestId common.String
}

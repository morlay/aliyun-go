package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteOfficeConversionTaskRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	TaskId  string `position:"Query" name:"TaskId"`
}

func (req *DeleteOfficeConversionTaskRequest) Invoke(client *sdk.Client) (resp *DeleteOfficeConversionTaskResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DeleteOfficeConversionTask", "imm", "")
	resp = &DeleteOfficeConversionTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteOfficeConversionTaskResponse struct {
	responses.BaseResponse
	RequestId common.String
}

package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetOpsCommandResultRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	EndCursor       int64  `position:"Query" name:"EndCursor"`
	StartCursor     int64  `position:"Query" name:"StartCursor"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	TaskId          int64  `position:"Query" name:"TaskId"`
}

func (req *GetOpsCommandResultRequest) Invoke(client *sdk.Client) (resp *GetOpsCommandResultResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetOpsCommandResult", "", "")
	resp = &GetOpsCommandResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetOpsCommandResultResponse struct {
	responses.BaseResponse
	RequestId common.String
	Content   common.String
	Cursor    common.Long
	Finished  bool
}

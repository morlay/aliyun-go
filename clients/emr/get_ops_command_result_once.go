package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetOpsCommandResultOnceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	TaskId          int64  `position:"Query" name:"TaskId"`
}

func (req *GetOpsCommandResultOnceRequest) Invoke(client *sdk.Client) (resp *GetOpsCommandResultOnceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetOpsCommandResultOnce", "", "")
	resp = &GetOpsCommandResultOnceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetOpsCommandResultOnceResponse struct {
	responses.BaseResponse
	RequestId string
	Content   string
}

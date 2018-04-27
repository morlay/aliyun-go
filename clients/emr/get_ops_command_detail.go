package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetOpsCommandDetailRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	OpsCommandName  string `position:"Query" name:"OpsCommandName"`
}

func (req *GetOpsCommandDetailRequest) Invoke(client *sdk.Client) (resp *GetOpsCommandDetailResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetOpsCommandDetail", "", "")
	resp = &GetOpsCommandDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetOpsCommandDetailResponse struct {
	responses.BaseResponse
	RequestId   string
	Id          int64
	Name        string
	Description string
	TargetType  string
	ServiceName string
	Category    string
	Params      string
}

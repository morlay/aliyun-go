package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	Id          common.Long
	Name        common.String
	Description common.String
	TargetType  common.String
	ServiceName common.String
	Category    common.String
	Params      common.String
}

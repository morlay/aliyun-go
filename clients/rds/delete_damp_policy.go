package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteDampPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	PolicyName           string `position:"Query" name:"PolicyName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteDampPolicyRequest) Invoke(client *sdk.Client) (resp *DeleteDampPolicyResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DeleteDampPolicy", "rds", "")
	resp = &DeleteDampPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDampPolicyResponse struct {
	responses.BaseResponse
	RequestId common.String
}

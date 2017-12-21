package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDampPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	PolicyName           string `position:"Query" name:"PolicyName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteDampPolicyRequest) Invoke(client *sdk.Client) (response *DeleteDampPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteDampPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DeleteDampPolicy", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DeleteDampPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteDampPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteDampPolicyResponse struct {
	RequestId string
}

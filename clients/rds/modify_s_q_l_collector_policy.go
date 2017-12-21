package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySQLCollectorPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	SQLCollectorStatus   string `position:"Query" name:"SQLCollectorStatus"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifySQLCollectorPolicyRequest) Invoke(client *sdk.Client) (response *ModifySQLCollectorPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifySQLCollectorPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifySQLCollectorPolicy", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifySQLCollectorPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifySQLCollectorPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifySQLCollectorPolicyResponse struct {
	RequestId string
}

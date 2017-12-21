package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddBuDBInstanceRelationRequest struct {
	BusinessUnit         string `position:"Query" name:"BusinessUnit"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r AddBuDBInstanceRelationRequest) Invoke(client *sdk.Client) (response *AddBuDBInstanceRelationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddBuDBInstanceRelationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "AddBuDBInstanceRelation", "rds", "")

	resp := struct {
		*responses.BaseResponse
		AddBuDBInstanceRelationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddBuDBInstanceRelationResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddBuDBInstanceRelationResponse struct {
	RequestId      string
	BusinessUnit   string
	DBInstanceName string
}

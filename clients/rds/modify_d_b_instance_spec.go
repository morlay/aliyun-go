package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    int    `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PayType              string `position:"Query" name:"PayType"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
}

func (r ModifyDBInstanceSpecRequest) Invoke(client *sdk.Client) (response *ModifyDBInstanceSpecResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBInstanceSpecRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceSpec", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBInstanceSpecResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBInstanceSpecResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBInstanceSpecResponse struct {
	RequestId string
}

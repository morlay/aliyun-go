package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceSSLRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyDBInstanceSSLRequest) Invoke(client *sdk.Client) (response *ModifyDBInstanceSSLResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBInstanceSSLRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceSSL", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBInstanceSSLResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBInstanceSSLResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBInstanceSSLResponse struct {
	RequestId string
}

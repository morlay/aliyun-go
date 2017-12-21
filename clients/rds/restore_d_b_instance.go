package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RestoreDBInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r RestoreDBInstanceRequest) Invoke(client *sdk.Client) (response *RestoreDBInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RestoreDBInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "RestoreDBInstance", "rds", "")

	resp := struct {
		*responses.BaseResponse
		RestoreDBInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RestoreDBInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RestoreDBInstanceResponse struct {
	RequestId string
}

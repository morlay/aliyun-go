package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RestoreDBInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BackupId             int    `position:"Query" name:"BackupId"`
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
	req.InitWithApiInfo("Dds", "2015-12-01", "RestoreDBInstance", "dds", "")

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

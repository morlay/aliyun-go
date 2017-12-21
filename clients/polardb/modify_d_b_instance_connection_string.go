package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceConnectionStringRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix  string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	DBInstanceId            string `position:"Query" name:"DBInstanceId"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
}

func (r ModifyDBInstanceConnectionStringRequest) Invoke(client *sdk.Client) (response *ModifyDBInstanceConnectionStringResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBInstanceConnectionStringRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBInstanceConnectionString", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBInstanceConnectionStringResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBInstanceConnectionStringResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBInstanceConnectionStringResponse struct {
	RequestId           string
	DBInstanceId        string
	OldConnectionString string
	OldPort             string
	NewConnectionString string
	NewPort             string
	IPType              string
}

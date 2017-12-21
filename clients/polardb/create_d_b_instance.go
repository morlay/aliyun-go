package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDBInstanceRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	SecurityIPList        string `position:"Query" name:"SecurityIPList"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string `position:"Query" name:"PrivateIpAddress"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
}

func (r CreateDBInstanceRequest) Invoke(client *sdk.Client) (response *CreateDBInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateDBInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "CreateDBInstance", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		CreateDBInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateDBInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateDBInstanceResponse struct {
	RequestId    string
	DBClusterId  string
	DBInstanceId string
	OrderId      string
}

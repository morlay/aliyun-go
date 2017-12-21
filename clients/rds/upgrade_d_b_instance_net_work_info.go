package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpgradeDBInstanceNetWorkInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r UpgradeDBInstanceNetWorkInfoRequest) Invoke(client *sdk.Client) (response *UpgradeDBInstanceNetWorkInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpgradeDBInstanceNetWorkInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "UpgradeDBInstanceNetWorkInfo", "rds", "")

	resp := struct {
		*responses.BaseResponse
		UpgradeDBInstanceNetWorkInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpgradeDBInstanceNetWorkInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpgradeDBInstanceNetWorkInfoResponse struct {
	RequestId string
}

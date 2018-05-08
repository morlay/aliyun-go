package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RestoreDBInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *RestoreDBInstanceRequest) Invoke(client *sdk.Client) (resp *RestoreDBInstanceResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "RestoreDBInstance", "rds", "")
	resp = &RestoreDBInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RestoreDBInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}

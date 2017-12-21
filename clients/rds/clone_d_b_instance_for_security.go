package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CloneDBInstanceForSecurityRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    int    `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	TargetAliBid         string `position:"Query" name:"TargetAliBid"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	TargetAliUid         string `position:"Query" name:"TargetAliUid"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	PayType              string `position:"Query" name:"PayType"`
}

func (req *CloneDBInstanceForSecurityRequest) Invoke(client *sdk.Client) (resp *CloneDBInstanceForSecurityResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CloneDBInstanceForSecurity", "rds", "")
	resp = &CloneDBInstanceForSecurityResponse{}
	err = client.DoAction(req, resp)
	return
}

type CloneDBInstanceForSecurityResponse struct {
	responses.BaseResponse
	RequestId        string
	DBInstanceId     string
	OrderId          string
	ConnectionString string
	Port             string
}

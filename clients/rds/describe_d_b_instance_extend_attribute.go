package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDBInstanceExtendAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBInstanceExtendAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceExtendAttributeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceExtendAttribute", "rds", "")
	resp = &DescribeDBInstanceExtendAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstanceExtendAttributeResponse struct {
	responses.BaseResponse
	RequestId                         common.String
	CanTempUpgrade                    bool
	TempUpgradeTimeStart              common.String
	TempUpgradeTimeEnd                common.String
	TempUpgradeRecoveryTime           common.String
	TempUpgradeRecoveryClass          common.String
	TempUpgradeRecoveryCpu            common.Integer
	TempUpgradeRecoveryMemory         common.Integer
	TempUpgradeRecoveryMaxIOPS        common.String
	TempUpgradeRecoveryMaxConnections common.String
}

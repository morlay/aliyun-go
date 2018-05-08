package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CloneDBInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	RestoreTime           string `position:"Query" name:"RestoreTime"`
	Period                string `position:"Query" name:"Period"`
	DBInstanceStorage     int    `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	BackupId              string `position:"Query" name:"BackupId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	UsedTime              string `position:"Query" name:"UsedTime"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string `position:"Query" name:"PrivateIpAddress"`
	ResourceGroupId       string `position:"Query" name:"ResourceGroupId"`
	VPCId                 string `position:"Query" name:"VPCId"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	DBInstanceId          string `position:"Query" name:"DBInstanceId"`
	PayType               string `position:"Query" name:"PayType"`
	InstanceNetworkType   string `position:"Query" name:"InstanceNetworkType"`
}

func (req *CloneDBInstanceRequest) Invoke(client *sdk.Client) (resp *CloneDBInstanceResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CloneDBInstance", "rds", "")
	resp = &CloneDBInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CloneDBInstanceResponse struct {
	responses.BaseResponse
	RequestId        common.String
	DBInstanceId     common.String
	OrderId          common.String
	ConnectionString common.String
	Port             common.String
}

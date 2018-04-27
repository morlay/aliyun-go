package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDBInstanceRequest struct {
	requests.RpcRequest
	ConnectionMode        string `position:"Query" name:"ConnectionMode"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     int    `position:"Query" name:"DBInstanceStorage"`
	SystemDBCharset       string `position:"Query" name:"SystemDBCharset"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	EngineVersion         string `position:"Query" name:"EngineVersion"`
	ResourceGroupId       string `position:"Query" name:"ResourceGroupId"`
	Engine                string `position:"Query" name:"Engine"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	DBInstanceNetType     string `position:"Query" name:"DBInstanceNetType"`
	Period                string `position:"Query" name:"Period"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	UsedTime              string `position:"Query" name:"UsedTime"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	SecurityIPList        string `position:"Query" name:"SecurityIPList"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string `position:"Query" name:"PrivateIpAddress"`
	VPCId                 string `position:"Query" name:"VPCId"`
	TunnelId              string `position:"Query" name:"TunnelId"`
	ZoneId                string `position:"Query" name:"ZoneId"`
	PayType               string `position:"Query" name:"PayType"`
	InstanceNetworkType   string `position:"Query" name:"InstanceNetworkType"`
}

func (req *CreateDBInstanceRequest) Invoke(client *sdk.Client) (resp *CreateDBInstanceResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateDBInstance", "rds", "")
	resp = &CreateDBInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateDBInstanceResponse struct {
	responses.BaseResponse
	RequestId        string
	DBInstanceId     string
	OrderId          string
	ConnectionString string
	Port             string
}

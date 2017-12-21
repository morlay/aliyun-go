package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CloneDBInstanceRequest struct {
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

func (r CloneDBInstanceRequest) Invoke(client *sdk.Client) (response *CloneDBInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CloneDBInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CloneDBInstance", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CloneDBInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CloneDBInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CloneDBInstanceResponse struct {
	RequestId        string
	DBInstanceId     string
	OrderId          string
	ConnectionString string
	Port             string
}

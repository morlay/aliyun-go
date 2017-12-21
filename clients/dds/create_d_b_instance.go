package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDBInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     int    `position:"Query" name:"DBInstanceStorage"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	CouponNo              string `position:"Query" name:"CouponNo"`
	EngineVersion         string `position:"Query" name:"EngineVersion"`
	NetworkType           string `position:"Query" name:"NetworkType"`
	StorageEngine         string `position:"Query" name:"StorageEngine"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	Engine                string `position:"Query" name:"Engine"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	BusinessInfo          string `position:"Query" name:"BusinessInfo"`
	Period                int    `position:"Query" name:"Period"`
	RestoreTime           string `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	SrcDBInstanceId       string `position:"Query" name:"SrcDBInstanceId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	BackupId              string `position:"Query" name:"BackupId"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	SecurityIPList        string `position:"Query" name:"SecurityIPList"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	AccountPassword       string `position:"Query" name:"AccountPassword"`
	VpcId                 string `position:"Query" name:"VpcId"`
	ZoneId                string `position:"Query" name:"ZoneId"`
	ChargeType            string `position:"Query" name:"ChargeType"`
}

func (req *CreateDBInstanceRequest) Invoke(client *sdk.Client) (resp *CreateDBInstanceResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "CreateDBInstance", "dds", "")
	resp = &CreateDBInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateDBInstanceResponse struct {
	responses.BaseResponse
	RequestId    string
	OrderId      string
	DBInstanceId string
}

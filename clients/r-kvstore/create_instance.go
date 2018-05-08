package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	NodeType             string `position:"Query" name:"NodeType"`
	CouponNo             string `position:"Query" name:"CouponNo"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	InstanceClass        string `position:"Query" name:"InstanceClass"`
	Capacity             int64  `position:"Query" name:"Capacity"`
	Password             string `position:"Query" name:"Password"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
	Period               string `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	SrcDBInstanceId      string `position:"Query" name:"SrcDBInstanceId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Token                string `position:"Query" name:"Token"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	InstanceName         string `position:"Query" name:"InstanceName"`
	VpcId                string `position:"Query" name:"VpcId"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	ChargeType           string `position:"Query" name:"ChargeType"`
	Config               string `position:"Query" name:"Config"`
}

func (req *CreateInstanceRequest) Invoke(client *sdk.Client) (resp *CreateInstanceResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateInstance", "redisa", "")
	resp = &CreateInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateInstanceResponse struct {
	responses.BaseResponse
	RequestId        common.String
	InstanceId       common.String
	InstanceName     common.String
	ConnectionDomain common.String
	Port             common.Integer
	UserName         common.String
	InstanceStatus   common.String
	RegionId         common.String
	Capacity         common.Long
	QPS              common.Long
	Bandwidth        common.Long
	Connections      common.Long
	ZoneId           common.String
	Config           common.String
	ChargeType       common.String
	EndTime          common.String
	NodeType         common.String
	NetworkType      common.String
	VpcId            common.String
	VSwitchId        common.String
	PrivateIpAddr    common.String
}


package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type CreateInstanceRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
NodeType string `position:"Query" name:"NodeType"`
CouponNo string `position:"Query" name:"CouponNo"`
NetworkType string `position:"Query" name:"NetworkType"`
InstanceClass string `position:"Query" name:"InstanceClass"`
Capacity int64 `position:"Query" name:"Capacity"`
Password string `position:"Query" name:"Password"`
SecurityToken string `position:"Query" name:"SecurityToken"`
InstanceType string `position:"Query" name:"InstanceType"`
BusinessInfo string `position:"Query" name:"BusinessInfo"`
Period string `position:"Query" name:"Period"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
SrcDBInstanceId string `position:"Query" name:"SrcDBInstanceId"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
BackupId string `position:"Query" name:"BackupId"`
OwnerId int64 `position:"Query" name:"OwnerId"`
Token string `position:"Query" name:"Token"`
VSwitchId string `position:"Query" name:"VSwitchId"`
InstanceName string `position:"Query" name:"InstanceName"`
VpcId string `position:"Query" name:"VpcId"`
ZoneId string `position:"Query" name:"ZoneId"`
ChargeType string `position:"Query" name:"ChargeType"`
Config string `position:"Query" name:"Config"`
}

func (r CreateInstanceRequest) Invoke(client *sdk.Client) (response *CreateInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateInstance", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		CreateInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.CreateInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateInstanceResponse struct {
RequestId string
InstanceId string
InstanceName string
ConnectionDomain string
Port int
UserName string
InstanceStatus string
RegionId string
Capacity int64
QPS int64
Bandwidth int64
Connections int64
ZoneId string
Config string
ChargeType string
EndTime string
NodeType string
NetworkType string
VpcId string
VSwitchId string
PrivateIpAddr string
}


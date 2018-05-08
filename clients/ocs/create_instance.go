package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateInstanceRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceName      string `position:"Query" name:"OcsInstanceName"`
	Password             string `position:"Query" name:"Password"`
	Capacity             int64  `position:"Query" name:"Capacity"`
	Region               string `position:"Query" name:"Region"`
	IzNo                 string `position:"Query" name:"IzNo"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	VpcId                string `position:"Query" name:"VpcId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	PrivateIp            string `position:"Query" name:"PrivateIp"`
}

func (req *CreateInstanceRequest) Invoke(client *sdk.Client) (resp *CreateInstanceResponse, err error) {
	req.InitWithApiInfo("Ocs", "2015-03-01", "CreateInstance", "", "")
	resp = &CreateInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateInstanceResponse struct {
	responses.BaseResponse
	RequestId   common.String
	OcsInstance CreateInstanceOcsInstance
}

type CreateInstanceOcsInstance struct {
	OcsInstanceId     common.String
	OcsInstanceName   common.String
	Capacity          common.Long
	Qps               common.Long
	Bandwidth         common.Long
	Connections       common.Long
	ConnectionDomain  common.String
	Port              common.Integer
	UserName          common.String
	RegionId          common.String
	OcsInstanceStatus common.String
	GmtCreated        common.String
	EndTime           common.String
	ChargeType        common.String
	IzId              common.String
	NetworkType       common.String
	VpcId             common.String
	VSwitchId         common.String
	PrivateIp         common.String
}

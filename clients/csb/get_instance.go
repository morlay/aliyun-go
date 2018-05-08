package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetInstanceRequest struct {
	requests.RpcRequest
	CsbId int64 `position:"Query" name:"CsbId"`
}

func (req *GetInstanceRequest) Invoke(client *sdk.Client) (resp *GetInstanceResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "GetInstance", "CSB", "")
	resp = &GetInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetInstanceResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      GetInstanceData
}

type GetInstanceData struct {
	Instance GetInstanceInstance
}

type GetInstanceInstance struct {
	ApprLevel            common.Integer
	ApprUser1            common.String
	ApprUser2            common.String
	BrokerVpcId          common.String
	BrokerVpcName        common.String
	ClientVpcId          common.String
	ClientVpcName        common.String
	ClusterMembers       common.Integer
	CredentialGroup      common.Long
	CsbAccountId         common.String
	CsbId                common.Long
	DbStatus             common.Integer
	Description          common.String
	FrontStatus          common.String
	GmtCreate            common.Long
	GmtModified          common.Long
	Id                   common.Long
	InstanceCategory     common.Integer
	InstanceType         common.Integer
	IpList               common.String
	IsImported           bool
	IsPublic             bool
	Name                 common.String
	OwnerId              common.String
	SentinelCtlStr       common.String
	SentinelCtrl         common.Long
	SentinelGridInterval common.Integer
	SentinelQps          common.Long
	Status               common.String
	StatusCode           common.Integer
	TenantId             common.String
	Testable             bool
	UserId               common.String
	Visible              bool
	VpcName              common.String
}

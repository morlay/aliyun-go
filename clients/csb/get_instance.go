package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      int
	Message   string
	RequestId string
	Data      GetInstanceData
}

type GetInstanceData struct {
	Instance GetInstanceInstance
}

type GetInstanceInstance struct {
	ApprLevel            int
	ApprUser1            string
	ApprUser2            string
	BrokerVpcId          string
	BrokerVpcName        string
	ClientVpcId          string
	ClientVpcName        string
	ClusterMembers       int
	CredentialGroup      int64
	CsbAccountId         string
	CsbId                int64
	DbStatus             int
	Description          string
	FrontStatus          string
	GmtCreate            int64
	GmtModified          int64
	Id                   int64
	InstanceCategory     int
	InstanceType         int
	IpList               string
	IsImported           bool
	IsPublic             bool
	Name                 string
	OwnerId              string
	SentinelCtlStr       string
	SentinelCtrl         int64
	SentinelGridInterval int
	SentinelQps          int64
	Status               string
	StatusCode           int
	TenantId             string
	Testable             bool
	UserId               string
	Visible              bool
	VpcName              string
}

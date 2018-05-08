package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SummaryRequest struct {
	requests.RpcRequest
	JstOwnerId  int64  `position:"Query" name:"JstOwnerId"`
	InstanceIds string `position:"Query" name:"InstanceIds"`
}

func (req *SummaryRequest) Invoke(client *sdk.Client) (resp *SummaryResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "Summary", "", "")
	resp = &SummaryResponse{}
	err = client.DoAction(req, resp)
	return
}

type SummaryResponse struct {
	responses.BaseResponse
	RequestId         common.String
	Status            common.Long
	AbnormalHostCount common.Long
	Ddos              SummaryDdos
	BruteForce        SummaryBruteForce
	Webshell          SummaryWebshell
	RemoteLogin       SummaryRemoteLogin
	WebAttack         SummaryWebAttack
	WebLeak           SummaryWebLeak
}

type SummaryDdos struct {
	Count     common.Long
	HostCount common.Long
}

type SummaryBruteForce struct {
	Count     common.Long
	HostCount common.Long
}

type SummaryWebshell struct {
	Count     common.Long
	HostCount common.Long
}

type SummaryRemoteLogin struct {
	Count     common.Long
	HostCount common.Long
}

type SummaryWebAttack struct {
	Count     common.Long
	HostCount common.Long
}

type SummaryWebLeak struct {
	Count     common.Long
	HostCount common.Long
}

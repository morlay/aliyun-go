package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId         string
	Status            int64
	AbnormalHostCount int64
	Ddos              SummaryDdos
	BruteForce        SummaryBruteForce
	Webshell          SummaryWebshell
	RemoteLogin       SummaryRemoteLogin
	WebAttack         SummaryWebAttack
	WebLeak           SummaryWebLeak
}

type SummaryDdos struct {
	Count     int64
	HostCount int64
}

type SummaryBruteForce struct {
	Count     int64
	HostCount int64
}

type SummaryWebshell struct {
	Count     int64
	HostCount int64
}

type SummaryRemoteLogin struct {
	Count     int64
	HostCount int64
}

type SummaryWebAttack struct {
	Count     int64
	HostCount int64
}

type SummaryWebLeak struct {
	Count     int64
	HostCount int64
}

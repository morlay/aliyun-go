package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SummaryRequest struct {
	JstOwnerId  int64  `position:"Query" name:"JstOwnerId"`
	InstanceIds string `position:"Query" name:"InstanceIds"`
}

func (r SummaryRequest) Invoke(client *sdk.Client) (response *SummaryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SummaryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "Summary", "", "")

	resp := struct {
		*responses.BaseResponse
		SummaryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SummaryResponse

	err = client.DoAction(&req, &resp)
	return
}

type SummaryResponse struct {
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

package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsEmpowerCreateRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	EmpowerUser  string `position:"Query" name:"EmpowerUser"`
	Topic        string `position:"Query" name:"Topic"`
	Relation     int    `position:"Query" name:"Relation"`
}

func (req *OnsEmpowerCreateRequest) Invoke(client *sdk.Client) (resp *OnsEmpowerCreateResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsEmpowerCreate", "", "")
	resp = &OnsEmpowerCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsEmpowerCreateResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
}

package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsEmpowerDeleteRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	EmpowerUser  string `position:"Query" name:"EmpowerUser"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsEmpowerDeleteRequest) Invoke(client *sdk.Client) (resp *OnsEmpowerDeleteResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsEmpowerDelete", "", "")
	resp = &OnsEmpowerDeleteResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsEmpowerDeleteResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
}

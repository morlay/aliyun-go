package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsEmpowerCreateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	EmpowerUser  string `position:"Query" name:"EmpowerUser"`
	Topic        string `position:"Query" name:"Topic"`
	Relation     int    `position:"Query" name:"Relation"`
}

func (r OnsEmpowerCreateRequest) Invoke(client *sdk.Client) (response *OnsEmpowerCreateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsEmpowerCreateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsEmpowerCreate", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsEmpowerCreateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsEmpowerCreateResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsEmpowerCreateResponse struct {
	RequestId string
	HelpUrl   string
}

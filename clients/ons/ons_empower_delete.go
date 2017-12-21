package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsEmpowerDeleteRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	EmpowerUser  string `position:"Query" name:"EmpowerUser"`
	Topic        string `position:"Query" name:"Topic"`
}

func (r OnsEmpowerDeleteRequest) Invoke(client *sdk.Client) (response *OnsEmpowerDeleteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsEmpowerDeleteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsEmpowerDelete", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsEmpowerDeleteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsEmpowerDeleteResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsEmpowerDeleteResponse struct {
	RequestId string
	HelpUrl   string
}

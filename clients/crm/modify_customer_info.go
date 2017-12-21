package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCustomerInfoRequest struct {
	KpId    int64  `position:"Query" name:"KpId"`
	Website string `position:"Query" name:"Website"`
	Biz     string `position:"Query" name:"Biz"`
}

func (r ModifyCustomerInfoRequest) Invoke(client *sdk.Client) (response *ModifyCustomerInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyCustomerInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Crm", "2015-03-24", "ModifyCustomerInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyCustomerInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyCustomerInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyCustomerInfoResponse struct {
	Success       bool
	ResultCode    string
	ResultMessage string
}

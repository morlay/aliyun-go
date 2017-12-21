package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindCustomerInfoRequest struct {
	KpId int64 `position:"Query" name:"KpId"`
}

func (r FindCustomerInfoRequest) Invoke(client *sdk.Client) (response *FindCustomerInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		FindCustomerInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Crm", "2015-03-24", "FindCustomerInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		FindCustomerInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.FindCustomerInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type FindCustomerInfoResponse struct {
	Success       bool
	ResultCode    string
	ResultMessage string
	Data          FindCustomerInfoData
}

type FindCustomerInfoData struct {
	Website string
	Biz     string
}

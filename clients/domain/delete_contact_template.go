package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteContactTemplateRequest struct {
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	Lang              string `position:"Query" name:"Lang"`
	ContactTemplateId int64  `position:"Query" name:"ContactTemplateId"`
}

func (r DeleteContactTemplateRequest) Invoke(client *sdk.Client) (response *DeleteContactTemplateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteContactTemplateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "DeleteContactTemplate", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteContactTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteContactTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteContactTemplateResponse struct {
	RequestId string
	Success   bool
}

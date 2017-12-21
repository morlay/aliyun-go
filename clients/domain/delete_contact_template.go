package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteContactTemplateRequest struct {
	requests.RpcRequest
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	Lang              string `position:"Query" name:"Lang"`
	ContactTemplateId int64  `position:"Query" name:"ContactTemplateId"`
}

func (req *DeleteContactTemplateRequest) Invoke(client *sdk.Client) (resp *DeleteContactTemplateResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "DeleteContactTemplate", "", "")
	resp = &DeleteContactTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteContactTemplateResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}

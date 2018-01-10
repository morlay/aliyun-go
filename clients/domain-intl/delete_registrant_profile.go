package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRegistrantProfileRequest struct {
	requests.RpcRequest
	UserClientIp        string `position:"Query" name:"UserClientIp"`
	RegistrantProfileId int64  `position:"Query" name:"RegistrantProfileId"`
	Lang                string `position:"Query" name:"Lang"`
}

func (req *DeleteRegistrantProfileRequest) Invoke(client *sdk.Client) (resp *DeleteRegistrantProfileResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "DeleteRegistrantProfile", "domain", "")
	resp = &DeleteRegistrantProfileResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteRegistrantProfileResponse struct {
	responses.BaseResponse
	RequestId string
}

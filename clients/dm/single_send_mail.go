package dm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SingleSendMailRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	AddressType          int    `position:"Query" name:"AddressType"`
	TagName              string `position:"Query" name:"TagName"`
	ReplyToAddress       string `position:"Query" name:"ReplyToAddress"`
	ToAddress            string `position:"Query" name:"ToAddress"`
	Subject              string `position:"Query" name:"Subject"`
	HtmlBody             string `position:"Query" name:"HtmlBody"`
	TextBody             string `position:"Query" name:"TextBody"`
	FromAlias            string `position:"Query" name:"FromAlias"`
	ReplyAddress         string `position:"Query" name:"ReplyAddress"`
	ReplyAddressAlias    string `position:"Query" name:"ReplyAddressAlias"`
	ClickTrace           string `position:"Query" name:"ClickTrace"`
}

func (req *SingleSendMailRequest) Invoke(client *sdk.Client) (resp *SingleSendMailResponse, err error) {
	req.InitWithApiInfo("Dm", "2015-11-23", "SingleSendMail", "", "")
	resp = &SingleSendMailResponse{}
	err = client.DoAction(req, resp)
	return
}

type SingleSendMailResponse struct {
	responses.BaseResponse
	RequestId string
	EnvId     string
}

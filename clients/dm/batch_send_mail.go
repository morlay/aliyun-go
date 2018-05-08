package dm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type BatchSendMailRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	TemplateName         string `position:"Query" name:"TemplateName"`
	AccountName          string `position:"Query" name:"AccountName"`
	ReceiversName        string `position:"Query" name:"ReceiversName"`
	AddressType          int    `position:"Query" name:"AddressType"`
	TagName              string `position:"Query" name:"TagName"`
	ReplyAddress         string `position:"Query" name:"ReplyAddress"`
	ReplyAddressAlias    string `position:"Query" name:"ReplyAddressAlias"`
	ClickTrace           string `position:"Query" name:"ClickTrace"`
}

func (req *BatchSendMailRequest) Invoke(client *sdk.Client) (resp *BatchSendMailResponse, err error) {
	req.InitWithApiInfo("Dm", "2015-11-23", "BatchSendMail", "", "")
	resp = &BatchSendMailResponse{}
	err = client.DoAction(req, resp)
	return
}

type BatchSendMailResponse struct {
	responses.BaseResponse
	RequestId common.String
	EnvId     common.String
}

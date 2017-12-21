package dm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BatchSendMailRequest struct {
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

func (r BatchSendMailRequest) Invoke(client *sdk.Client) (response *BatchSendMailResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BatchSendMailRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dm", "2015-11-23", "BatchSendMail", "", "")

	resp := struct {
		*responses.BaseResponse
		BatchSendMailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BatchSendMailResponse

	err = client.DoAction(&req, &resp)
	return
}

type BatchSendMailResponse struct {
	RequestId string
	EnvId     string
}

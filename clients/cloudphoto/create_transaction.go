package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateTransactionRequest struct {
	requests.RpcRequest
	Ext       string `position:"Query" name:"Ext"`
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Force     string `position:"Query" name:"Force"`
	Md5       string `position:"Query" name:"Md.5"`
}

func (req *CreateTransactionRequest) Invoke(client *sdk.Client) (resp *CreateTransactionResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreateTransaction", "cloudphoto", "")
	resp = &CreateTransactionResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateTransactionResponse struct {
	responses.BaseResponse
	Code        string
	Message     string
	RequestId   string
	Action      string
	Transaction CreateTransactionTransaction
}

type CreateTransactionTransaction struct {
	Upload CreateTransactionUpload
}

type CreateTransactionUpload struct {
	Bucket          string
	FileId          string
	OssEndpoint     string
	ObjectKey       string
	SessionId       string
	AccessKeyId     string
	AccessKeySecret string
	StsToken        string
}

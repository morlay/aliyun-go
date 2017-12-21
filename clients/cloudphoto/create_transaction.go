package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateTransactionRequest struct {
	Ext       string `position:"Query" name:"Ext"`
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Force     string `position:"Query" name:"Force"`
	Md5       string `position:"Query" name:"Md.5"`
}

func (r CreateTransactionRequest) Invoke(client *sdk.Client) (response *CreateTransactionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateTransactionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreateTransaction", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		CreateTransactionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateTransactionResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateTransactionResponse struct {
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

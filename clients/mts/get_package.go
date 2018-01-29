package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPackageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Data                 string `position:"Query" name:"Data"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *GetPackageRequest) Invoke(client *sdk.Client) (resp *GetPackageResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "GetPackage", "mts", "")
	resp = &GetPackageResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPackageResponse struct {
	responses.BaseResponse
	RequestId   string
	CertPackage string
}

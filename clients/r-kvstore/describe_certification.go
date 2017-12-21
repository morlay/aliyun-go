
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeCertificationRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
Parameters string `position:"Query" name:"Parameters"`
}

func (r DescribeCertificationRequest) Invoke(client *sdk.Client) (response *DescribeCertificationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCertificationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeCertification", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCertificationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeCertificationResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCertificationResponse struct {
RequestId string
NoCertification bool
}


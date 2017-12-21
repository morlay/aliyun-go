package ossadmin

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseOssInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Region               string `position:"Query" name:"Region"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

func (r ReleaseOssInstanceRequest) Invoke(client *sdk.Client) (response *ReleaseOssInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReleaseOssInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("OssAdmin", "2015-03-02", "ReleaseOssInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		ReleaseOssInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReleaseOssInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReleaseOssInstanceResponse struct {
	RequestId string
}

package ossadmin

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopOssInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Region               string `position:"Query" name:"Region"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

func (r StopOssInstanceRequest) Invoke(client *sdk.Client) (response *StopOssInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StopOssInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("OssAdmin", "2015-03-02", "StopOssInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		StopOssInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StopOssInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type StopOssInstanceResponse struct {
	RequestId string
}

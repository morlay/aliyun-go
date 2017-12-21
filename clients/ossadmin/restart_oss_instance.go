package ossadmin

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RestartOssInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Region               string `position:"Query" name:"Region"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

func (r RestartOssInstanceRequest) Invoke(client *sdk.Client) (response *RestartOssInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RestartOssInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("OssAdmin", "2015-03-02", "RestartOssInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		RestartOssInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RestartOssInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RestartOssInstanceResponse struct {
	RequestId string
}

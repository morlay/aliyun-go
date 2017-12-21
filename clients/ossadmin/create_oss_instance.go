package ossadmin

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateOssInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Region               string `position:"Query" name:"Region"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

func (r CreateOssInstanceRequest) Invoke(client *sdk.Client) (response *CreateOssInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateOssInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("OssAdmin", "2015-03-02", "CreateOssInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateOssInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateOssInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateOssInstanceResponse struct {
	RequestId      string
	AliUid         int64
	InstanceId     string
	InstacneStatus string
	InstanceName   string
	StartTime      string
	EndTime        string
}

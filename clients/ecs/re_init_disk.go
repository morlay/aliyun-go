package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReInitDiskRequest struct {
	ResourceOwnerId             int64  `position:"Query" name:"ResourceOwnerId"`
	Password                    string `position:"Query" name:"Password"`
	ResourceOwnerAccount        string `position:"Query" name:"ResourceOwnerAccount"`
	AutoStartInstance           string `position:"Query" name:"AutoStartInstance"`
	OwnerAccount                string `position:"Query" name:"OwnerAccount"`
	DiskId                      string `position:"Query" name:"DiskId"`
	SecurityEnhancementStrategy string `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string `position:"Query" name:"KeyPairName"`
	OwnerId                     int64  `position:"Query" name:"OwnerId"`
}

func (r ReInitDiskRequest) Invoke(client *sdk.Client) (response *ReInitDiskResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReInitDiskRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ReInitDisk", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ReInitDiskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReInitDiskResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReInitDiskResponse struct {
	RequestId string
}

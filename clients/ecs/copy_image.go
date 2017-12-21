package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CopyImageRequest struct {
	Tag4Value              string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId                string `position:"Query" name:"ImageId"`
	Tag2Key                string `position:"Query" name:"Tag.2.Key"`
	Tag5Key                string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	DestinationImageName   string `position:"Query" name:"DestinationImageName"`
	DestinationRegionId    string `position:"Query" name:"DestinationRegionId"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	Tag3Key                string `position:"Query" name:"Tag.3.Key"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	Tag5Value              string `position:"Query" name:"Tag.5.Value"`
	Tag1Key                string `position:"Query" name:"Tag.1.Key"`
	Tag1Value              string `position:"Query" name:"Tag.1.Value"`
	Encrypted              string `position:"Query" name:"Encrypted"`
	Tag2Value              string `position:"Query" name:"Tag.2.Value"`
	Tag4Key                string `position:"Query" name:"Tag.4.Key"`
	DestinationDescription string `position:"Query" name:"DestinationDescription"`
	Tag3Value              string `position:"Query" name:"Tag.3.Value"`
}

func (r CopyImageRequest) Invoke(client *sdk.Client) (response *CopyImageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CopyImageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CopyImage", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CopyImageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CopyImageResponse

	err = client.DoAction(&req, &resp)
	return
}

type CopyImageResponse struct {
	RequestId string
	ImageId   string
}

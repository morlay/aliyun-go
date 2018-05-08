package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ReplaceSystemDiskRequest struct {
	requests.RpcRequest
	ResourceOwnerId             int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId                     string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount        string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken                 string `position:"Query" name:"ClientToken"`
	OwnerAccount                string `position:"Query" name:"OwnerAccount"`
	SecurityEnhancementStrategy string `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string `position:"Query" name:"KeyPairName"`
	OwnerId                     int64  `position:"Query" name:"OwnerId"`
	Platform                    string `position:"Query" name:"Platform"`
	Password                    string `position:"Query" name:"Password"`
	InstanceId                  string `position:"Query" name:"InstanceId"`
	SystemDiskSize              int    `position:"Query" name:"SystemDiskSize"`
	DiskId                      string `position:"Query" name:"DiskId"`
	UseAdditionalService        string `position:"Query" name:"UseAdditionalService"`
	Architecture                string `position:"Query" name:"Architecture"`
}

func (req *ReplaceSystemDiskRequest) Invoke(client *sdk.Client) (resp *ReplaceSystemDiskResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ReplaceSystemDisk", "ecs", "")
	resp = &ReplaceSystemDiskResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReplaceSystemDiskResponse struct {
	responses.BaseResponse
	RequestId common.String
	DiskId    common.String
}

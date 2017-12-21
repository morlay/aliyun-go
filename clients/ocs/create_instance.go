package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceName      string `position:"Query" name:"OcsInstanceName"`
	Password             string `position:"Query" name:"Password"`
	Capacity             int64  `position:"Query" name:"Capacity"`
	Region               string `position:"Query" name:"Region"`
	IzNo                 string `position:"Query" name:"IzNo"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	VpcId                string `position:"Query" name:"VpcId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	PrivateIp            string `position:"Query" name:"PrivateIp"`
}

func (r CreateInstanceRequest) Invoke(client *sdk.Client) (response *CreateInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "CreateInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateInstanceResponse struct {
	RequestId   string
	OcsInstance CreateInstanceOcsInstance
}

type CreateInstanceOcsInstance struct {
	OcsInstanceId     string
	OcsInstanceName   string
	Capacity          int64
	Qps               int64
	Bandwidth         int64
	Connections       int64
	ConnectionDomain  string
	Port              int
	UserName          string
	RegionId          string
	OcsInstanceStatus string
	GmtCreated        string
	EndTime           string
	ChargeType        string
	IzId              string
	NetworkType       string
	VpcId             string
	VSwitchId         string
	PrivateIp         string
}

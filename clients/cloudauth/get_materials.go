package cloudauth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetMaterialsRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Biz             string `position:"Query" name:"Biz"`
	TicketId        string `position:"Query" name:"TicketId"`
}

func (req *GetMaterialsRequest) Invoke(client *sdk.Client) (resp *GetMaterialsResponse, err error) {
	req.InitWithApiInfo("Cloudauth", "2017-11-17", "GetMaterials", "cloudauth", "")
	resp = &GetMaterialsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetMaterialsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      GetMaterialsData
}

type GetMaterialsData struct {
	Name                 common.String
	IdentificationNumber common.String
	IdCardType           common.String
	IdCardExpiry         common.String
	Address              common.String
	Sex                  common.String
	IdCardFrontPic       common.String
	IdCardBackPic        common.String
	FacePic              common.String
}

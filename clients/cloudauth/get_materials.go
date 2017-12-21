package cloudauth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      GetMaterialsData
}

type GetMaterialsData struct {
	Name                 string
	IdentificationNumber string
	IdCardType           string
	IdCardExpiry         string
	Address              string
	Sex                  string
	IdCardFrontPic       string
	IdCardBackPic        string
	FacePic              string
}

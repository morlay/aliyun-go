package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetOrderRequest struct {
	requests.RpcRequest
	OrderId     int64  `position:"Query" name:"OrderId"`
	ServiceName string `position:"Query" name:"ServiceName"`
}

func (req *GetOrderRequest) Invoke(client *sdk.Client) (resp *GetOrderResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "GetOrder", "CSB", "")
	resp = &GetOrderResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetOrderResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      GetOrderData
}

type GetOrderData struct {
	Order GetOrderOrder
}

type GetOrderOrder struct {
	Alias                 common.String
	CredentialGroupId     common.Long
	CsbId                 common.Long
	DauthGroupName        common.String
	GmtCreate             common.Long
	GmtModified           common.Long
	GroupName             common.String
	Id                    common.Long
	ProjectName           common.String
	ServiceId             common.Long
	ServiceName           common.String
	ServiceStatus         common.Integer
	ServiceVersion        common.String
	StatisticName         common.String
	Status                common.Integer
	StrictWhiteListJson   common.String
	UserId                common.String
	ErrorTypeCatagoryList GetOrderErrorTypeCatagoryList
	StrictWhiteList       GetOrderStrictWhiteListList
	Service               GetOrderService
	SlaInfo               GetOrderSlaInfo
	Total                 GetOrderTotal
}

type GetOrderErrorTypeCatagory struct {
	Total    common.Integer
	ErrorNum common.Integer
	Name     common.String
}

type GetOrderService struct {
	AccessParamsJSON    common.String
	Active              bool
	Alias               common.String
	AllVisiable         bool
	ConsumeTypesJSON    common.String
	CreateTime          common.Long
	CsbId               common.Long
	ErrDefJSON          common.String
	Id                  common.Long
	InterfaceName       common.String
	OldVersion          common.String
	OttFlag             bool
	OwnerId             common.String
	PrincipalName       common.String
	ProjectId           common.String
	ProjectName         common.String
	ProvideType         common.String
	SSL                 bool
	Scope               common.String
	ServiceName         common.String
	ServiceProviderType common.String
	ServiceVersion      common.String
	SkipAuth            bool
	StatisticName       common.String
	Status              common.Integer
	UserId              common.Long
	ValidConsumeTypes   bool
	ValidProvideType    bool
}

type GetOrderSlaInfo struct {
	Qph common.String
	Qps common.String
}

type GetOrderTotal struct {
	ErrorNum common.Integer
	Total    common.Integer
}

type GetOrderErrorTypeCatagoryList []GetOrderErrorTypeCatagory

func (list *GetOrderErrorTypeCatagoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetOrderErrorTypeCatagory)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type GetOrderStrictWhiteListList []common.String

func (list *GetOrderStrictWhiteListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

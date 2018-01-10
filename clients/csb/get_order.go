package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      int
	Message   string
	RequestId string
	Data      GetOrderData
}

type GetOrderData struct {
	Order GetOrderOrder
}

type GetOrderOrder struct {
	Alias                 string
	CredentialGroupId     int64
	CsbId                 int64
	DauthGroupName        string
	GmtCreate             int64
	GmtModified           int64
	GroupName             string
	Id                    int64
	ProjectName           string
	ServiceId             int64
	ServiceName           string
	ServiceStatus         int
	ServiceVersion        string
	StatisticName         string
	Status                int
	StrictWhiteListJson   string
	UserId                string
	ErrorTypeCatagoryList GetOrderErrorTypeCatagoryList
	StrictWhiteList       GetOrderStrictWhiteListList
	Service               GetOrderService
	SlaInfo               GetOrderSlaInfo
	Total                 GetOrderTotal
}

type GetOrderErrorTypeCatagory struct {
	Total    int
	ErrorNum int
	Name     string
}

type GetOrderService struct {
	AccessParamsJSON    string
	Active              bool
	Alias               string
	AllVisiable         bool
	ConsumeTypesJSON    string
	CreateTime          int64
	CsbId               int64
	ErrDefJSON          string
	Id                  int64
	InterfaceName       string
	OldVersion          string
	OttFlag             bool
	OwnerId             string
	PrincipalName       string
	ProjectId           string
	ProjectName         string
	ProvideType         string
	SSL                 bool
	Scope               string
	ServiceName         string
	ServiceProviderType string
	ServiceVersion      string
	SkipAuth            bool
	StatisticName       string
	Status              int
	UserId              int64
	ValidConsumeTypes   bool
	ValidProvideType    bool
}

type GetOrderSlaInfo struct {
	Qph string
	Qps string
}

type GetOrderTotal struct {
	ErrorNum int
	Total    int
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

type GetOrderStrictWhiteListList []string

func (list *GetOrderStrictWhiteListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

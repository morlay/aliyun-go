package crm

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *CrmClient) FindCustomerInfo(req *FindCustomerInfoArgs) (resp *FindCustomerInfoResponse, err error) {
	resp = &FindCustomerInfoResponse{}
	err = c.Request("FindCustomerInfo", req, resp)
	return
}

type FindCustomerInfoData struct {
	Website string
	Biz     string
}
type FindCustomerInfoArgs struct {
	KpId int64
}
type FindCustomerInfoResponse struct {
	Success       core.Bool
	ResultCode    string
	ResultMessage string
	Data          FindCustomerInfoData
}

func (c *CrmClient) QueryEnumConfigByType(req *QueryEnumConfigByTypeArgs) (resp *QueryEnumConfigByTypeResponse, err error) {
	resp = &QueryEnumConfigByTypeResponse{}
	err = c.Request("QueryEnumConfigByType", req, resp)
	return
}

type QueryEnumConfigByTypeBizCategory struct {
	EnumName  string
	EnumValue string
}
type QueryEnumConfigByTypeArgs struct {
	Type string
}
type QueryEnumConfigByTypeResponse struct {
	Success       core.Bool
	ResultCode    string
	ResultMessage string
	Data          QueryEnumConfigByTypeBizCategoryList
}

type QueryEnumConfigByTypeBizCategoryList []QueryEnumConfigByTypeBizCategory

func (list *QueryEnumConfigByTypeBizCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEnumConfigByTypeBizCategory)
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

func (c *CrmClient) FindContacter(req *FindContacterArgs) (resp *FindContacterResponse, err error) {
	resp = &FindContacterResponse{}
	err = c.Request("FindContacter", req, resp)
	return
}

type FindContacterData struct {
	ContacterId       int64
	KpId              int64
	CustomerId        int64
	ContacterName     string
	ContacterEmail    string
	ContacterMobile   string
	ContacterPosition string
}
type FindContacterArgs struct {
	ContacterId int64
}
type FindContacterResponse struct {
	Success       core.Bool
	ResultCode    string
	ResultMessage string
	Data          FindContacterData
}

func (c *CrmClient) ModifyContacter(req *ModifyContacterArgs) (resp *ModifyContacterResponse, err error) {
	resp = &ModifyContacterResponse{}
	err = c.Request("ModifyContacter", req, resp)
	return
}

type ModifyContacterArgs struct {
	ContacterId       int64
	ContacterName     string
	ContacterEmail    string
	ContacterMobile   string
	ContacterPosition string
}
type ModifyContacterResponse struct {
	Success       core.Bool
	ResultCode    string
	ResultMessage string
}

func (c *CrmClient) DeleteContacter(req *DeleteContacterArgs) (resp *DeleteContacterResponse, err error) {
	resp = &DeleteContacterResponse{}
	err = c.Request("DeleteContacter", req, resp)
	return
}

type DeleteContacterArgs struct {
	ContacterId int64
}
type DeleteContacterResponse struct {
	Success       core.Bool
	ResultCode    string
	ResultMessage string
}

func (c *CrmClient) ModifyCustomerInfo(req *ModifyCustomerInfoArgs) (resp *ModifyCustomerInfoResponse, err error) {
	resp = &ModifyCustomerInfoResponse{}
	err = c.Request("ModifyCustomerInfo", req, resp)
	return
}

type ModifyCustomerInfoArgs struct {
	KpId    int64
	Website string
	Biz     string
}
type ModifyCustomerInfoResponse struct {
	Success       core.Bool
	ResultCode    string
	ResultMessage string
}

func (c *CrmClient) FindAllContacter(req *FindAllContacterArgs) (resp *FindAllContacterResponse, err error) {
	resp = &FindAllContacterResponse{}
	err = c.Request("FindAllContacter", req, resp)
	return
}

type FindAllContacterContacterInfo struct {
	ContacterId       int64
	KpId              int64
	CustomerId        int64
	ContacterName     string
	ContacterEmail    string
	ContacterMobile   string
	ContacterPosition string
}
type FindAllContacterArgs struct {
	KpId int64
}
type FindAllContacterResponse struct {
	Success       core.Bool
	ResultCode    string
	ResultMessage string
	Data          FindAllContacterContacterInfoList
}

type FindAllContacterContacterInfoList []FindAllContacterContacterInfo

func (list *FindAllContacterContacterInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindAllContacterContacterInfo)
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

func (c *CrmClient) CreateContacter(req *CreateContacterArgs) (resp *CreateContacterResponse, err error) {
	resp = &CreateContacterResponse{}
	err = c.Request("CreateContacter", req, resp)
	return
}

type CreateContacterData struct {
	ContacterId int64
}
type CreateContacterArgs struct {
	KpId              int64
	ContacterName     string
	ContacterEmail    string
	ContacterMobile   string
	ContacterPosition string
}
type CreateContacterResponse struct {
	Success       core.Bool
	ResultCode    string
	ResultMessage string
	Data          CreateContacterData
}

func (c *CrmClient) FindBizCategoryConfig(req *FindBizCategoryConfigArgs) (resp *FindBizCategoryConfigResponse, err error) {
	resp = &FindBizCategoryConfigResponse{}
	err = c.Request("FindBizCategoryConfig", req, resp)
	return
}

type FindBizCategoryConfigBizCategory struct {
	Name       string
	Code       string
	IsCheck    core.Bool
	MainBiz    core.Bool
	Other      string
	SubConfigs FindBizCategoryConfigBizSubCategoryList
}

type FindBizCategoryConfigBizSubCategory struct {
	Name    string
	Code    string
	IsCheck core.Bool
	MainBiz core.Bool
	Other   string
}
type FindBizCategoryConfigArgs struct {
	KpId int64
}
type FindBizCategoryConfigResponse struct {
	Success       core.Bool
	ResultCode    string
	ResultMessage string
	Data          FindBizCategoryConfigBizCategoryList
}

type FindBizCategoryConfigBizSubCategoryList []FindBizCategoryConfigBizSubCategory

func (list *FindBizCategoryConfigBizSubCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindBizCategoryConfigBizSubCategory)
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

type FindBizCategoryConfigBizCategoryList []FindBizCategoryConfigBizCategory

func (list *FindBizCategoryConfigBizCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindBizCategoryConfigBizCategory)
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

func (c *CrmClient) FindContacterTest(req *FindContacterTestArgs) (resp *FindContacterTestResponse, err error) {
	resp = &FindContacterTestResponse{}
	err = c.Request("FindContacterTest", req, resp)
	return
}

type FindContacterTestData struct {
	ContacterId       int64
	KpId              int64
	CustomerId        int64
	ContacterName     string
	ContacterEmail    string
	ContacterMobile   string
	ContacterPosition string
}
type FindContacterTestArgs struct {
	ContacterId int64
}
type FindContacterTestResponse struct {
	Success       core.Bool
	ResultCode    string
	ResultMessage string
	Data          FindContacterTestData
}

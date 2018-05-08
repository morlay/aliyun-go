package crm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type FindBizCategoryConfigRequest struct {
	requests.RpcRequest
	KpId int64 `position:"Query" name:"KpId"`
}

func (req *FindBizCategoryConfigRequest) Invoke(client *sdk.Client) (resp *FindBizCategoryConfigResponse, err error) {
	req.InitWithApiInfo("Crm", "2015-03-24", "FindBizCategoryConfig", "", "")
	resp = &FindBizCategoryConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindBizCategoryConfigResponse struct {
	responses.BaseResponse
	Success       bool
	ResultCode    common.String
	ResultMessage common.String
	Data          FindBizCategoryConfigBizCategoryList
}

type FindBizCategoryConfigBizCategory struct {
	Name       common.String
	Code       common.String
	IsCheck    bool
	MainBiz    bool
	Other      common.String
	SubConfigs FindBizCategoryConfigBizSubCategoryList
}

type FindBizCategoryConfigBizSubCategory struct {
	Name    common.String
	Code    common.String
	IsCheck bool
	MainBiz bool
	Other   common.String
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

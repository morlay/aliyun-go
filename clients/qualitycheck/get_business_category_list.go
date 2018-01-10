package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetBusinessCategoryListRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *GetBusinessCategoryListRequest) Invoke(client *sdk.Client) (resp *GetBusinessCategoryListResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetBusinessCategoryList", "", "")
	resp = &GetBusinessCategoryListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetBusinessCategoryListResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      GetBusinessCategoryListBusinessCategoryBasicInfoList
}

type GetBusinessCategoryListBusinessCategoryBasicInfo struct {
	Bid          int
	ServiceType  int
	BusinessName string
}

type GetBusinessCategoryListBusinessCategoryBasicInfoList []GetBusinessCategoryListBusinessCategoryBasicInfo

func (list *GetBusinessCategoryListBusinessCategoryBasicInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetBusinessCategoryListBusinessCategoryBasicInfo)
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

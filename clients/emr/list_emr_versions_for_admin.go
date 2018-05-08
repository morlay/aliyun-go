package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListEmrVersionsForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
	PageSize        int   `position:"Query" name:"PageSize"`
	PageNumber      int   `position:"Query" name:"PageNumber"`
}

func (req *ListEmrVersionsForAdminRequest) Invoke(client *sdk.Client) (resp *ListEmrVersionsForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListEmrVersionsForAdmin", "", "")
	resp = &ListEmrVersionsForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListEmrVersionsForAdminResponse struct {
	responses.BaseResponse
	RequestId               common.String
	PageNumber              common.Integer
	PageSize                common.Integer
	TotalCount              common.Integer
	MainVersionKeyValueList ListEmrVersionsForAdminMainVersionKeyValueList
}

type ListEmrVersionsForAdminMainVersionKeyValue struct {
	Id                   common.Integer
	Key                  common.String
	Value                common.String
	Status               common.Integer
	UtcCreateTimestamp   common.Long
	UtcModifiedTimestamp common.Long
}

type ListEmrVersionsForAdminMainVersionKeyValueList []ListEmrVersionsForAdminMainVersionKeyValue

func (list *ListEmrVersionsForAdminMainVersionKeyValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEmrVersionsForAdminMainVersionKeyValue)
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

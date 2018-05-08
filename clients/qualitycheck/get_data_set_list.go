package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetDataSetListRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *GetDataSetListRequest) Invoke(client *sdk.Client) (resp *GetDataSetListResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetDataSetList", "", "")
	resp = &GetDataSetListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetDataSetListResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Count     common.Integer
	Data      GetDataSetListDataSetList
}

type GetDataSetListDataSet struct {
	SetId         common.Long
	SetName       common.String
	SetDomain     common.String
	SetRoleArn    common.String
	SetBucketName common.String
	SetFolderName common.String
	ChannelType   common.Integer
	CreateType    common.Integer
}

type GetDataSetListDataSetList []GetDataSetListDataSet

func (list *GetDataSetListDataSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetDataSetListDataSet)
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

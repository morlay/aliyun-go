package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Code      string
	Message   string
	Count     int
	Data      GetDataSetListDataSetList
}

type GetDataSetListDataSet struct {
	SetId         int64
	SetName       string
	SetDomain     string
	SetRoleArn    string
	SetBucketName string
	SetFolderName string
	ChannelType   int
	CreateType    int
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

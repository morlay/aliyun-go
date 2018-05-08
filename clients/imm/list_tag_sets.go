package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListTagSetsRequest struct {
	requests.RpcRequest
	MaxKeys int    `position:"Query" name:"MaxKeys"`
	Marker  string `position:"Query" name:"Marker"`
	Project string `position:"Query" name:"Project"`
}

func (req *ListTagSetsRequest) Invoke(client *sdk.Client) (resp *ListTagSetsResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListTagSets", "imm", "")
	resp = &ListTagSetsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListTagSetsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	NextMarker common.String
	Sets       ListTagSetsSetsItemList
}

type ListTagSetsSetsItem struct {
	SetId      common.String
	Status     common.String
	Photos     common.Long
	CreateTime common.String
	ModifyTime common.String
}

type ListTagSetsSetsItemList []ListTagSetsSetsItem

func (list *ListTagSetsSetsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagSetsSetsItem)
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

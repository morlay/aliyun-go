package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListFaceSetsRequest struct {
	requests.RpcRequest
	Marker  string `position:"Query" name:"Marker"`
	Project string `position:"Query" name:"Project"`
}

func (req *ListFaceSetsRequest) Invoke(client *sdk.Client) (resp *ListFaceSetsResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListFaceSets", "imm", "")
	resp = &ListFaceSetsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFaceSetsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	NextMarker common.String
	Sets       ListFaceSetsSetsItemList
}

type ListFaceSetsSetsItem struct {
	SetId      common.String
	Status     common.String
	Photos     common.Long
	CreateTime common.String
	ModifyTime common.String
}

type ListFaceSetsSetsItemList []ListFaceSetsSetsItem

func (list *ListFaceSetsSetsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFaceSetsSetsItem)
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

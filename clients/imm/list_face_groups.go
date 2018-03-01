package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListFaceGroupsRequest struct {
	requests.RpcRequest
	MaxKeys int    `position:"Query" name:"MaxKeys"`
	Marker  int    `position:"Query" name:"Marker"`
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
}

func (req *ListFaceGroupsRequest) Invoke(client *sdk.Client) (resp *ListFaceGroupsResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListFaceGroups", "imm", "")
	resp = &ListFaceGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFaceGroupsResponse struct {
	responses.BaseResponse
	RequestId  string
	NextMarker int
	Groups     ListFaceGroupsGroupsItemList
}

type ListFaceGroupsGroupsItem struct {
	GroupId int
	FaceNum int
}

type ListFaceGroupsGroupsItemList []ListFaceGroupsGroupsItem

func (list *ListFaceGroupsGroupsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFaceGroupsGroupsItem)
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

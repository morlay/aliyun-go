package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GroupFacesRequest struct {
	requests.RpcRequest
	Project   string `position:"Query" name:"Project"`
	SetId     string `position:"Query" name:"SetId"`
	Operation string `position:"Query" name:"Operation"`
}

func (req *GroupFacesRequest) Invoke(client *sdk.Client) (resp *GroupFacesResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GroupFaces", "imm", "")
	resp = &GroupFacesResponse{}
	err = client.DoAction(req, resp)
	return
}

type GroupFacesResponse struct {
	responses.BaseResponse
	RequestId common.String
	SetId     common.String
	HasMore   common.Integer
	Groups    GroupFacesGroupsItemList
}

type GroupFacesGroupsItem struct {
	FaceId        common.String
	GroupId       common.String
	UnGroupReason common.String
}

type GroupFacesGroupsItemList []GroupFacesGroupsItem

func (list *GroupFacesGroupsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GroupFacesGroupsItem)
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

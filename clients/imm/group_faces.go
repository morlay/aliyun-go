package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	SetId     string
	HasMore   int
	Groups    GroupFacesGroupsList
}

type GroupFacesGroups struct {
	FaceId        string
	GroupId       string
	UnGroupReason string
}

type GroupFacesGroupsList []GroupFacesGroups

func (list *GroupFacesGroupsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GroupFacesGroups)
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

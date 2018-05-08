package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetMyGroupsRequest struct {
	requests.RpcRequest
	SelectContactGroups string `position:"Query" name:"SelectContactGroups"`
	InstanceId          string `position:"Query" name:"InstanceId"`
	GroupId             int64  `position:"Query" name:"GroupId"`
	Type                string `position:"Query" name:"Type"`
	GroupName           string `position:"Query" name:"GroupName"`
	BindUrl             string `position:"Query" name:"BindUrl"`
}

func (req *GetMyGroupsRequest) Invoke(client *sdk.Client) (resp *GetMyGroupsResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "GetMyGroups", "cms", "")
	resp = &GetMyGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetMyGroupsResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorCode    common.Integer
	ErrorMessage common.String
	Group        GetMyGroupsGroup
}

type GetMyGroupsGroup struct {
	GroupId       common.Long
	GroupName     common.String
	ServiceId     common.Long
	BindUrl       common.String
	Type          common.String
	ContactGroups GetMyGroupsContactGroupList
}

type GetMyGroupsContactGroup struct {
	Name common.String
}

type GetMyGroupsContactGroupList []GetMyGroupsContactGroup

func (list *GetMyGroupsContactGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetMyGroupsContactGroup)
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

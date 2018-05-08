package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type FindProjectListRequest struct {
	requests.RpcRequest
	ProjectName string `position:"Query" name:"ProjectName"`
	CsbId       int64  `position:"Query" name:"CsbId"`
	PageNum     int    `position:"Query" name:"PageNum"`
}

func (req *FindProjectListRequest) Invoke(client *sdk.Client) (resp *FindProjectListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "FindProjectList", "CSB", "")
	resp = &FindProjectListResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindProjectListResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      FindProjectListData
}

type FindProjectListData struct {
	CurrentPage common.Integer
	PageNumber  common.Integer
	Total       common.Integer
	ProjectList FindProjectListProjectList
}

type FindProjectListProject struct {
	ApiNum               common.Integer
	CsbId                common.Long
	DeleteFlag           common.Integer
	Description          common.String
	GmtCreate            common.Long
	GmtModified          common.Long
	Id                   common.Long
	InterfaceJarLocation common.String
	InterfaceJarName     common.String
	JarFileKey           common.String
	OwnerId              common.String
	ProjectName          common.String
	ProjectOwnerEmail    common.String
	ProjectOwnerName     common.String
	ProjectOwnerPhoneNum common.String
	Status               common.Integer
	UserId               common.String
}

type FindProjectListProjectList []FindProjectListProject

func (list *FindProjectListProjectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindProjectListProject)
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

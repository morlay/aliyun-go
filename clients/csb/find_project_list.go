package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      int
	Message   string
	RequestId string
	Data      FindProjectListData
}

type FindProjectListData struct {
	CurrentPage int
	PageNumber  int
	Total       int
	ProjectList FindProjectListProjectList
}

type FindProjectListProject struct {
	ApiNum               int
	CsbId                int64
	DeleteFlag           int
	Description          string
	GmtCreate            int64
	GmtModified          int64
	Id                   int64
	InterfaceJarLocation string
	InterfaceJarName     string
	JarFileKey           string
	OwnerId              string
	ProjectName          string
	ProjectOwnerEmail    string
	ProjectOwnerName     string
	ProjectOwnerPhoneNum string
	Status               int
	UserId               string
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

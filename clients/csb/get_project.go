package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetProjectRequest struct {
	requests.RpcRequest
	ProjectName string `position:"Query" name:"ProjectName"`
	CsbId       int64  `position:"Query" name:"CsbId"`
}

func (req *GetProjectRequest) Invoke(client *sdk.Client) (resp *GetProjectResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "GetProject", "CSB", "")
	resp = &GetProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetProjectResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      GetProjectData
}

type GetProjectData struct {
	ProjectList GetProjectProjectList
}

type GetProjectProject struct {
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

type GetProjectProjectList []GetProjectProject

func (list *GetProjectProjectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProjectProject)
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

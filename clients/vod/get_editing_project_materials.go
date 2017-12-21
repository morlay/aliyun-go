package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetEditingProjectMaterialsRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Type                 string `position:"Query" name:"Type"`
	ProjectId            string `position:"Query" name:"ProjectId"`
}

func (r GetEditingProjectMaterialsRequest) Invoke(client *sdk.Client) (response *GetEditingProjectMaterialsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetEditingProjectMaterialsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "GetEditingProjectMaterials", "", "")

	resp := struct {
		*responses.BaseResponse
		GetEditingProjectMaterialsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetEditingProjectMaterialsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetEditingProjectMaterialsResponse struct {
	RequestId    string
	MaterialList GetEditingProjectMaterialsMaterialList
}

type GetEditingProjectMaterialsMaterial struct {
	MaterialId   string
	Title        string
	Tags         string
	Status       string
	Size         int64
	Duration     float32
	Description  string
	CreationTime string
	ModifiedTime string
	CoverURL     string
	CateId       int
	CateName     string
	Source       string
	Snapshots    GetEditingProjectMaterialsSnapshotList
	Sprites      GetEditingProjectMaterialsSpriteList
}

type GetEditingProjectMaterialsMaterialList []GetEditingProjectMaterialsMaterial

func (list *GetEditingProjectMaterialsMaterialList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetEditingProjectMaterialsMaterial)
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

type GetEditingProjectMaterialsSnapshotList []string

func (list *GetEditingProjectMaterialsSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type GetEditingProjectMaterialsSpriteList []string

func (list *GetEditingProjectMaterialsSpriteList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

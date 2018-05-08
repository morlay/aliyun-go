package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetEditingProjectMaterialsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Type                 string `position:"Query" name:"Type"`
	ProjectId            string `position:"Query" name:"ProjectId"`
}

func (req *GetEditingProjectMaterialsRequest) Invoke(client *sdk.Client) (resp *GetEditingProjectMaterialsResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetEditingProjectMaterials", "vod", "")
	resp = &GetEditingProjectMaterialsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetEditingProjectMaterialsResponse struct {
	responses.BaseResponse
	RequestId    common.String
	MaterialList GetEditingProjectMaterialsMaterialList
}

type GetEditingProjectMaterialsMaterial struct {
	MaterialId   common.String
	Title        common.String
	Tags         common.String
	Status       common.String
	Size         common.Long
	Duration     common.Float
	Description  common.String
	CreationTime common.String
	ModifiedTime common.String
	CoverURL     common.String
	CateId       common.Integer
	CateName     common.String
	Source       common.String
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

type GetEditingProjectMaterialsSnapshotList []common.String

func (list *GetEditingProjectMaterialsSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type GetEditingProjectMaterialsSpriteList []common.String

func (list *GetEditingProjectMaterialsSpriteList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

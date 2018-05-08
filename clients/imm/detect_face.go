package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DetectFaceRequest struct {
	requests.RpcRequest
	SrcUris string `position:"Query" name:"SrcUris"`
	Project string `position:"Query" name:"Project"`
}

func (req *DetectFaceRequest) Invoke(client *sdk.Client) (resp *DetectFaceResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DetectFace", "imm", "")
	resp = &DetectFaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetectFaceResponse struct {
	responses.BaseResponse
	RequestId      common.String
	SuccessDetails DetectFaceSuccessDetailsItemList
	FailDetails    DetectFaceFailDetailsItemList
	SrcUris        DetectFaceSrcUriList
}

type DetectFaceSuccessDetailsItem struct {
	SrcUri  common.String
	PhotoId common.String
	Faces   DetectFaceFacesItemList
}

type DetectFaceFacesItem struct {
	FaceId        common.String
	FaceRectangle DetectFaceFaceRectangle
	FaceAttribute DetectFaceFaceAttribute
}

type DetectFaceFaceRectangle struct {
	Top    common.String
	Left   common.String
	Width  common.String
	Height common.String
}

type DetectFaceFaceAttribute struct {
	Gender      DetectFaceGender
	Age         DetectFaceAge
	Headpose    DetectFaceHeadpose
	Eyestatus   DetectFaceEyestatus
	Blur        DetectFaceBlur
	Facequality DetectFaceFacequality
}

type DetectFaceGender struct {
	Value common.String
}

type DetectFaceAge struct {
	Value common.Integer
}

type DetectFaceHeadpose struct {
	Pitch_angle common.Float
	Roll_angle  common.Float
	Yaw_angle   common.Float
}

type DetectFaceEyestatus struct {
	Left_eye_status  DetectFaceLeft_eye_status
	Right_eye_status DetectFaceRight_eye_status
}

type DetectFaceLeft_eye_status struct {
	Normal_glass_eye_open  common.Float
	No_glass_eye_close     common.Float
	Occlusion              common.Float
	No_glass_eye_open      common.Float
	Normal_glass_eye_close common.Float
	Dark_glasses           common.Float
}

type DetectFaceRight_eye_status struct {
	Normal_glass_eye_open  common.Float
	No_glass_eye_close     common.Float
	Occlusion              common.Float
	No_glass_eye_open      common.Float
	Normal_glass_eye_close common.Float
	Dark_glasses           common.Float
}

type DetectFaceBlur struct {
	Blurness DetectFaceBlurness
}

type DetectFaceBlurness struct {
	Value     common.Float
	Threshold common.Float
}

type DetectFaceFacequality struct {
	Value     common.Float
	Threshold common.Float
}

type DetectFaceFailDetailsItem struct {
	SrcUri common.String
	Reason common.String
}

type DetectFaceSuccessDetailsItemList []DetectFaceSuccessDetailsItem

func (list *DetectFaceSuccessDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectFaceSuccessDetailsItem)
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

type DetectFaceFailDetailsItemList []DetectFaceFailDetailsItem

func (list *DetectFaceFailDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectFaceFailDetailsItem)
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

type DetectFaceSrcUriList []common.String

func (list *DetectFaceSrcUriList) UnmarshalJSON(data []byte) error {
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

type DetectFaceFacesItemList []DetectFaceFacesItem

func (list *DetectFaceFacesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectFaceFacesItem)
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

// Code generated by thriftgo (0.3.13). DO NOT EDIT.

package user_server

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cutejiuges/disk_back/kitex_gen/base"
	"strings"
)

type UserInfoMeta struct {
	UserName *string `thrift:"user_name,1,optional" frugal:"1,optional,string" json:"user_name,omitempty"`
	Password string  `thrift:"password,2,required" frugal:"2,required,string" json:"password"`
	Email    string  `thrift:"email,3,required" frugal:"3,required,string" json:"email"`
	Phone    *string `thrift:"phone,4,optional" frugal:"4,optional,string" json:"phone,omitempty"`
	Profile  *string `thrift:"profile,5,optional" frugal:"5,optional,string" json:"profile,omitempty"`
}

func NewUserInfoMeta() *UserInfoMeta {
	return &UserInfoMeta{}
}

func (p *UserInfoMeta) InitDefault() {
}

var UserInfoMeta_UserName_DEFAULT string

func (p *UserInfoMeta) GetUserName() (v string) {
	if !p.IsSetUserName() {
		return UserInfoMeta_UserName_DEFAULT
	}
	return *p.UserName
}

func (p *UserInfoMeta) GetPassword() (v string) {
	return p.Password
}

func (p *UserInfoMeta) GetEmail() (v string) {
	return p.Email
}

var UserInfoMeta_Phone_DEFAULT string

func (p *UserInfoMeta) GetPhone() (v string) {
	if !p.IsSetPhone() {
		return UserInfoMeta_Phone_DEFAULT
	}
	return *p.Phone
}

var UserInfoMeta_Profile_DEFAULT string

func (p *UserInfoMeta) GetProfile() (v string) {
	if !p.IsSetProfile() {
		return UserInfoMeta_Profile_DEFAULT
	}
	return *p.Profile
}
func (p *UserInfoMeta) SetUserName(val *string) {
	p.UserName = val
}
func (p *UserInfoMeta) SetPassword(val string) {
	p.Password = val
}
func (p *UserInfoMeta) SetEmail(val string) {
	p.Email = val
}
func (p *UserInfoMeta) SetPhone(val *string) {
	p.Phone = val
}
func (p *UserInfoMeta) SetProfile(val *string) {
	p.Profile = val
}

var fieldIDToName_UserInfoMeta = map[int16]string{
	1: "user_name",
	2: "password",
	3: "email",
	4: "phone",
	5: "profile",
}

func (p *UserInfoMeta) IsSetUserName() bool {
	return p.UserName != nil
}

func (p *UserInfoMeta) IsSetPhone() bool {
	return p.Phone != nil
}

func (p *UserInfoMeta) IsSetProfile() bool {
	return p.Profile != nil
}

func (p *UserInfoMeta) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetPassword bool = false
	var issetEmail bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetPassword = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
				issetEmail = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetPassword {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetEmail {
		fieldId = 3
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserInfoMeta[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_UserInfoMeta[fieldId]))
}

func (p *UserInfoMeta) ReadField1(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.UserName = _field
	return nil
}
func (p *UserInfoMeta) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Password = _field
	return nil
}
func (p *UserInfoMeta) ReadField3(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Email = _field
	return nil
}
func (p *UserInfoMeta) ReadField4(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Phone = _field
	return nil
}
func (p *UserInfoMeta) ReadField5(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Profile = _field
	return nil
}

func (p *UserInfoMeta) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UserInfoMeta"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserInfoMeta) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetUserName() {
		if err = oprot.WriteFieldBegin("user_name", thrift.STRING, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.UserName); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UserInfoMeta) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("password", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Password); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *UserInfoMeta) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("email", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Email); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *UserInfoMeta) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetPhone() {
		if err = oprot.WriteFieldBegin("phone", thrift.STRING, 4); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Phone); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *UserInfoMeta) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetProfile() {
		if err = oprot.WriteFieldBegin("profile", thrift.STRING, 5); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Profile); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *UserInfoMeta) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserInfoMeta(%+v)", *p)

}

func (p *UserInfoMeta) DeepEqual(ano *UserInfoMeta) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.UserName) {
		return false
	}
	if !p.Field2DeepEqual(ano.Password) {
		return false
	}
	if !p.Field3DeepEqual(ano.Email) {
		return false
	}
	if !p.Field4DeepEqual(ano.Phone) {
		return false
	}
	if !p.Field5DeepEqual(ano.Profile) {
		return false
	}
	return true
}

func (p *UserInfoMeta) Field1DeepEqual(src *string) bool {

	if p.UserName == src {
		return true
	} else if p.UserName == nil || src == nil {
		return false
	}
	if strings.Compare(*p.UserName, *src) != 0 {
		return false
	}
	return true
}
func (p *UserInfoMeta) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Password, src) != 0 {
		return false
	}
	return true
}
func (p *UserInfoMeta) Field3DeepEqual(src string) bool {

	if strings.Compare(p.Email, src) != 0 {
		return false
	}
	return true
}
func (p *UserInfoMeta) Field4DeepEqual(src *string) bool {

	if p.Phone == src {
		return true
	} else if p.Phone == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Phone, *src) != 0 {
		return false
	}
	return true
}
func (p *UserInfoMeta) Field5DeepEqual(src *string) bool {

	if p.Profile == src {
		return true
	} else if p.Profile == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Profile, *src) != 0 {
		return false
	}
	return true
}

type UserSignUpRequest struct {
	UserInfo *UserInfoMeta `thrift:"user_info,1,required" frugal:"1,required,UserInfoMeta" json:"user_info"`
	Base     *base.Base    `thrift:"base,255,optional" frugal:"255,optional,base.Base" json:"base,omitempty"`
}

func NewUserSignUpRequest() *UserSignUpRequest {
	return &UserSignUpRequest{}
}

func (p *UserSignUpRequest) InitDefault() {
}

var UserSignUpRequest_UserInfo_DEFAULT *UserInfoMeta

func (p *UserSignUpRequest) GetUserInfo() (v *UserInfoMeta) {
	if !p.IsSetUserInfo() {
		return UserSignUpRequest_UserInfo_DEFAULT
	}
	return p.UserInfo
}

var UserSignUpRequest_Base_DEFAULT *base.Base

func (p *UserSignUpRequest) GetBase() (v *base.Base) {
	if !p.IsSetBase() {
		return UserSignUpRequest_Base_DEFAULT
	}
	return p.Base
}
func (p *UserSignUpRequest) SetUserInfo(val *UserInfoMeta) {
	p.UserInfo = val
}
func (p *UserSignUpRequest) SetBase(val *base.Base) {
	p.Base = val
}

var fieldIDToName_UserSignUpRequest = map[int16]string{
	1:   "user_info",
	255: "base",
}

func (p *UserSignUpRequest) IsSetUserInfo() bool {
	return p.UserInfo != nil
}

func (p *UserSignUpRequest) IsSetBase() bool {
	return p.Base != nil
}

func (p *UserSignUpRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetUserInfo bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetUserInfo = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetUserInfo {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserSignUpRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_UserSignUpRequest[fieldId]))
}

func (p *UserSignUpRequest) ReadField1(iprot thrift.TProtocol) error {
	_field := NewUserInfoMeta()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.UserInfo = _field
	return nil
}
func (p *UserSignUpRequest) ReadField255(iprot thrift.TProtocol) error {
	_field := base.NewBase()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Base = _field
	return nil
}

func (p *UserSignUpRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UserSignUpRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserSignUpRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("user_info", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.UserInfo.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UserSignUpRequest) writeField255(oprot thrift.TProtocol) (err error) {
	if p.IsSetBase() {
		if err = oprot.WriteFieldBegin("base", thrift.STRUCT, 255); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Base.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *UserSignUpRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserSignUpRequest(%+v)", *p)

}

func (p *UserSignUpRequest) DeepEqual(ano *UserSignUpRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.UserInfo) {
		return false
	}
	if !p.Field255DeepEqual(ano.Base) {
		return false
	}
	return true
}

func (p *UserSignUpRequest) Field1DeepEqual(src *UserInfoMeta) bool {

	if !p.UserInfo.DeepEqual(src) {
		return false
	}
	return true
}
func (p *UserSignUpRequest) Field255DeepEqual(src *base.Base) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}

type UserSignUpResponse struct {
	AccountId int64          `thrift:"account_id,1,required" frugal:"1,required,i64" json:"account_id"`
	BaseResp  *base.BaseResp `thrift:"base_resp,255,optional" frugal:"255,optional,base.BaseResp" json:"base_resp,omitempty"`
}

func NewUserSignUpResponse() *UserSignUpResponse {
	return &UserSignUpResponse{}
}

func (p *UserSignUpResponse) InitDefault() {
}

func (p *UserSignUpResponse) GetAccountId() (v int64) {
	return p.AccountId
}

var UserSignUpResponse_BaseResp_DEFAULT *base.BaseResp

func (p *UserSignUpResponse) GetBaseResp() (v *base.BaseResp) {
	if !p.IsSetBaseResp() {
		return UserSignUpResponse_BaseResp_DEFAULT
	}
	return p.BaseResp
}
func (p *UserSignUpResponse) SetAccountId(val int64) {
	p.AccountId = val
}
func (p *UserSignUpResponse) SetBaseResp(val *base.BaseResp) {
	p.BaseResp = val
}

var fieldIDToName_UserSignUpResponse = map[int16]string{
	1:   "account_id",
	255: "base_resp",
}

func (p *UserSignUpResponse) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *UserSignUpResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetAccountId bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetAccountId = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetAccountId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserSignUpResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_UserSignUpResponse[fieldId]))
}

func (p *UserSignUpResponse) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.AccountId = _field
	return nil
}
func (p *UserSignUpResponse) ReadField255(iprot thrift.TProtocol) error {
	_field := base.NewBaseResp()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.BaseResp = _field
	return nil
}

func (p *UserSignUpResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UserSignUpResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserSignUpResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("account_id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.AccountId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UserSignUpResponse) writeField255(oprot thrift.TProtocol) (err error) {
	if p.IsSetBaseResp() {
		if err = oprot.WriteFieldBegin("base_resp", thrift.STRUCT, 255); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.BaseResp.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *UserSignUpResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserSignUpResponse(%+v)", *p)

}

func (p *UserSignUpResponse) DeepEqual(ano *UserSignUpResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.AccountId) {
		return false
	}
	if !p.Field255DeepEqual(ano.BaseResp) {
		return false
	}
	return true
}

func (p *UserSignUpResponse) Field1DeepEqual(src int64) bool {

	if p.AccountId != src {
		return false
	}
	return true
}
func (p *UserSignUpResponse) Field255DeepEqual(src *base.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}

type UserService interface {
	UserSignUp(ctx context.Context, req *UserSignUpRequest) (r *UserSignUpResponse, err error)
}

type UserServiceUserSignUpArgs struct {
	Req *UserSignUpRequest `thrift:"req,1,required" frugal:"1,required,UserSignUpRequest" json:"req"`
}

func NewUserServiceUserSignUpArgs() *UserServiceUserSignUpArgs {
	return &UserServiceUserSignUpArgs{}
}

func (p *UserServiceUserSignUpArgs) InitDefault() {
}

var UserServiceUserSignUpArgs_Req_DEFAULT *UserSignUpRequest

func (p *UserServiceUserSignUpArgs) GetReq() (v *UserSignUpRequest) {
	if !p.IsSetReq() {
		return UserServiceUserSignUpArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *UserServiceUserSignUpArgs) SetReq(val *UserSignUpRequest) {
	p.Req = val
}

var fieldIDToName_UserServiceUserSignUpArgs = map[int16]string{
	1: "req",
}

func (p *UserServiceUserSignUpArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserServiceUserSignUpArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetReq bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetReq = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetReq {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserServiceUserSignUpArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_UserServiceUserSignUpArgs[fieldId]))
}

func (p *UserServiceUserSignUpArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewUserSignUpRequest()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *UserServiceUserSignUpArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UserSignUp_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserServiceUserSignUpArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UserServiceUserSignUpArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceUserSignUpArgs(%+v)", *p)

}

func (p *UserServiceUserSignUpArgs) DeepEqual(ano *UserServiceUserSignUpArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *UserServiceUserSignUpArgs) Field1DeepEqual(src *UserSignUpRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type UserServiceUserSignUpResult struct {
	Success *UserSignUpResponse `thrift:"success,0,optional" frugal:"0,optional,UserSignUpResponse" json:"success,omitempty"`
}

func NewUserServiceUserSignUpResult() *UserServiceUserSignUpResult {
	return &UserServiceUserSignUpResult{}
}

func (p *UserServiceUserSignUpResult) InitDefault() {
}

var UserServiceUserSignUpResult_Success_DEFAULT *UserSignUpResponse

func (p *UserServiceUserSignUpResult) GetSuccess() (v *UserSignUpResponse) {
	if !p.IsSetSuccess() {
		return UserServiceUserSignUpResult_Success_DEFAULT
	}
	return p.Success
}
func (p *UserServiceUserSignUpResult) SetSuccess(x interface{}) {
	p.Success = x.(*UserSignUpResponse)
}

var fieldIDToName_UserServiceUserSignUpResult = map[int16]string{
	0: "success",
}

func (p *UserServiceUserSignUpResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserServiceUserSignUpResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserServiceUserSignUpResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UserServiceUserSignUpResult) ReadField0(iprot thrift.TProtocol) error {
	_field := NewUserSignUpResponse()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *UserServiceUserSignUpResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UserSignUp_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserServiceUserSignUpResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *UserServiceUserSignUpResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceUserSignUpResult(%+v)", *p)

}

func (p *UserServiceUserSignUpResult) DeepEqual(ano *UserServiceUserSignUpResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *UserServiceUserSignUpResult) Field0DeepEqual(src *UserSignUpResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

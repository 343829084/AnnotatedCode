//Package propertyf comment
// This file war generated by tars2go 1.1
// Generated from PropertyF.tars
package propertyf

import (
	"context"
	"fmt"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
)

//PropertyF struct
type PropertyF struct {
	s m.Servant
}

//ReportPropMsg is the proxy function for the method defined in the tars file, with the context
func (_obj *PropertyF) ReportPropMsg(Statmsg map[StatPropMsgHead]StatPropMsgBody, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.WriteHead(codec.MAP, 1)
	if err != nil {
		return ret, err
	}
	err = _os.Write_int32(int32(len(Statmsg)), 0)
	if err != nil {
		return ret, err
	}
	for k1, v1 := range Statmsg {

		err = k1.WriteBlock(_os, 0)
		if err != nil {
			return ret, err
		}

		err = v1.WriteBlock(_os, 1)
		if err != nil {
			return ret, err
		}
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "reportPropMsg", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//ReportPropMsgWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *PropertyF) ReportPropMsgWithContext(ctx context.Context, Statmsg map[StatPropMsgHead]StatPropMsgBody, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.WriteHead(codec.MAP, 1)
	if err != nil {
		return ret, err
	}
	err = _os.Write_int32(int32(len(Statmsg)), 0)
	if err != nil {
		return ret, err
	}
	for k2, v2 := range Statmsg {

		err = k2.WriteBlock(_os, 0)
		if err != nil {
			return ret, err
		}

		err = v2.WriteBlock(_os, 1)
		if err != nil {
			return ret, err
		}
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "reportPropMsg", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SetServant sets servant for the service.
func (_obj *PropertyF) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *PropertyF) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

type _impPropertyF interface {
	ReportPropMsg(Statmsg map[StatPropMsgHead]StatPropMsgBody) (ret int32, err error)
}
type _impPropertyFWithContext interface {
	ReportPropMsg(ctx context.Context, Statmsg map[StatPropMsgHead]StatPropMsgBody) (ret int32, err error)
}

//Dispatch is used to call the server side implemnet for the method defined in the tars file. withContext shows using context or not.
func (_obj *PropertyF) Dispatch(ctx context.Context, _val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(tools.Int8ToByte(req.SBuffer))
	_os := codec.NewBuffer()
	switch req.SFuncName {
	case "reportPropMsg":
		var Statmsg map[StatPropMsgHead]StatPropMsgBody
		err, have = _is.SkipTo(codec.MAP, 1, true)
		if err != nil {
			return err
		}

		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		Statmsg = make(map[StatPropMsgHead]StatPropMsgBody)
		for i3, e3 := int32(0), length; i3 < e3; i3++ {
			var k3 StatPropMsgHead
			var v3 StatPropMsgBody

			err = k3.ReadBlock(_is, 0, false)
			if err != nil {
				return err
			}

			err = v3.ReadBlock(_is, 1, false)
			if err != nil {
				return err
			}

			Statmsg[k3] = v3
		}
		if withContext == false {
			_imp := _val.(_impPropertyF)
			ret, err := _imp.ReportPropMsg(Statmsg)
			if err != nil {
				return err
			}

			err = _os.Write_int32(ret, 0)
			if err != nil {
				return err
			}
		} else {
			_imp := _val.(_impPropertyFWithContext)
			ret, err := _imp.ReportPropMsg(ctx, Statmsg)
			if err != nil {
				return err
			}

			err = _os.Write_int32(ret, 0)
			if err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(ctx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(ctx)
	if ok && c != nil {
		_context = c
	}
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}
	_ = length
	_ = have
	_ = ty
	return nil
}
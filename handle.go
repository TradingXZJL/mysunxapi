package mysunxapi

import "fmt"

type SunxErrorRes struct {
	Code    int    `json:"code,omitempty"`    // 错误码
	Status  any    `json:"status"`            // 请求处理结果 ok , "error"
	Message string `json:"message,omitempty"` // 错误信息
	ErrMsg  string `json:"err-msg,omitempty"` // 错误信息
}

type SunxTimeRes struct {
	Ts int64 `json:"ts"` // 响应生成时间点，单位：毫秒
}

type SunxRestRes[T any] struct {
	SunxErrorRes
	SunxTimeRes
	Ch   string `json:"ch,omitempty"` // 数据所属主题
	Data T      `json:"data,omitempty"`
}

func (r *SunxRestRes[T]) UnmarshalJSON(data []byte) error {
	type Alias SunxRestRes[T]
	aux := &struct {
		Data  *T `json:"data,omitempty"`
		Tick  *T `json:"tick,omitempty"`
		Ticks *T `json:"ticks,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Data != nil {
		r.Data = *aux.Data
	}
	if aux.Tick != nil {
		r.Data = *aux.Tick
	}
	if aux.Ticks != nil {
		r.Data = *aux.Ticks
	}
	return nil
}

func handlerCommonRes[T any](body []byte) (*SunxRestRes[T], error) {
	res := &SunxRestRes[T]{}
	err := json.Unmarshal(body, res)
	if err != nil {
		log.Error("Rest返回值: ", string(body))
		// log.Error("Rest返回值解析失败: ", err)
		return nil, err
	}
	return res, nil
}

func (err *SunxErrorRes) handlerError() error {
	if (err.Code == 200 || err.Code == 0) &&
		(err.Status == "" || err.Status == nil || err.Status == "ok" || err.Status == 200) &&
		(err.Message == "Success" || err.Message == "") && err.ErrMsg == "" {
		return nil
	}
	return fmt.Errorf("request error: [code:%v][status:%v][message:%v][errMsg:%v]", err.Code, err.Status, err.Message, err.ErrMsg)
}

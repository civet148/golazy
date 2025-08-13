package svc

import (
	"github.com/civet148/sqlca/v3"
	"reflect"
	"test/internal/config"
)

type CodeMsg interface {
	Code() int32
	Msg() string
}

type Response struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type ServiceContext struct {
	DB     *sqlca.Engine
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	var opts = []sqlca.Option{
		sqlca.WithDebug(),
		sqlca.WithMaxConn(c.Orm.MaxConns),
		sqlca.WithIdleConn(c.Orm.IdleConns),
	}
	if c.Orm.Debug {
		opts = append(opts, sqlca.WithDebug())
	}
	if c.Orm.NodeId > 0 {
		opts = append(opts, sqlca.WithSnowFlake(&sqlca.SnowFlake{
			NodeId: c.Orm.NodeId,
		}))
	}
	db, err := sqlca.NewEngine(c.Orm.DSN, opts...)
	if err != nil {
		panic(err)
		return nil
	}
	return &ServiceContext{
		DB:     db,
		Config: c,
	}
}

func JsonResponse(rsp any, err error) any {
	if err != nil {
		switch e := err.(type) {
		case CodeMsg:
			return Response{
				Code: e.Code(),
				Msg:  e.Msg(),
				Data: struct{}{},
			}
		default:
			return Response{
				Code: -1,
				Msg:  err.Error(),
				Data: struct{}{},
			}
		}
	}
	makeEmptySlice(rsp)
	return Response{
		Code: 0,
		Msg:  "OK",
		Data: rsp,
	}
}

// check response object has empty slice or not
func makeEmptySlice(v any) {
	initSlices(reflect.ValueOf(v))
}

func initSlices(v reflect.Value) {
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return //ignore nil pointer
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return //ignore non-struct type
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Slice && field.IsNil() {
			field.Set(reflect.MakeSlice(field.Type(), 0, 0))
		}

		if field.Kind() == reflect.Struct {
			initSlices(field)
		} else if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			initSlices(field.Elem())
		} else if field.Kind() == reflect.Slice {
			for j := 0; j < field.Len(); j++ {
				elem := field.Index(j)
				if elem.Kind() == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
					initSlices(elem.Elem())
				}
			}
		}
	}
}

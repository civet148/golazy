package svc

import (
	"fmt"
	"github.com/civet148/sqlca/v3"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"test/internal/config"
)

type CodeMsg interface {
	Code() int32
	Msg() string
}

type response struct {
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
			return response{
				Code: e.Code(),
				Msg:  e.Msg(),
				Data: struct{}{},
			}
		default:
			return response{
				Code: -1,
				Msg:  err.Error(),
				Data: struct{}{},
			}
		}
	}
	makeEmptySlice(rsp)
	return response{
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

func ShouldBindParams(c *gin.Context, obj any) error {

	val := reflect.ValueOf(obj).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		var strTag string
		field := typ.Field(i)
		strFormTag := field.Tag.Get("form")
		strParamTag := field.Tag.Get("param")
		strJsonTag := field.Tag.Get("json")

		if strParamTag != "" {
			strTag = strParamTag
		} else if strFormTag != "" {
			strTag = strFormTag
		} else if strJsonTag != "" {
			strTag = strJsonTag
		} else {
			continue
		}

		paramValue := c.Param(strTag)
		if paramValue == "" {
			continue
		}

		switch field.Type.Kind() {
		case reflect.String:
			val.Field(i).SetString(paramValue)
		case reflect.Float32, reflect.Float64:
			uintValue, err := strconv.ParseFloat(paramValue, 10)
			if err != nil {
				return err
			}
			val.Field(i).SetFloat(uintValue)
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			uintValue, err := strconv.ParseInt(paramValue, 10, 64)
			if err != nil {
				return err
			}
			val.Field(i).SetInt(uintValue)
		case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
			uintValue, err := strconv.ParseUint(paramValue, 10, 64)
			if err != nil {
				return err
			}
			val.Field(i).SetUint(uintValue)
		default:
			return fmt.Errorf("unsupported type: %s", field.Type.Kind())
		}
	}
	return nil
}

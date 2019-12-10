package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-every-day/dktest/web/common"
	"go-every-day/dktest/web/controller"
	"go-every-day/dktest/web/utils"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"runtime"
	"time"
)

func InitDb(dsn, dblog string) (engine *xorm.Engine, err error) {
	//engine, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:6379)/test_db_name?charset=utf8")
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return
	}

	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)

	f := utils.NewXormLogger(dblog)
	if f == nil {
		panic("Init db log failed")
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
	if err != nil {
		engine.Logger().Error(err.Error())
		return
	} else {
		engine.Logger().Info("New Enging Ok")
	}

	//校验连接
	err = engine.Ping()
	if err != nil {
		engine.Logger().Error(err.Error())
		return
	} else {
		engine.Logger().Info("Ping Ok")
	}

	return
}

func PanicRecover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		cc := e.(*common.CustomContext)
		defer func() {
			host, _ := os.Hostname()
			if r := recover(); r != nil {
				var err error
				switch r := r.(type) {
				case error:
					err = r
				default:
					err = fmt.Errorf("%v", r)
				}
				stack := make([]byte, 1024)
				runtime.Stack(stack, true)
				cc.Blob(http.StatusInternalServerError, "", nil)
				fmt.Println(host, err.Error())
			}
		}()

		fmt.Println("PanicRecover func")
		return next(e)
	}
}

func RequestLantency(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		out := next(e)
		cc := e.(*common.CustomContext)
		latency := time.Now().Sub(cc.StartTime)
		fmt.Println("RequestLantency: ", latency)
		return out
	}
}

//解析一个json串是否含两端带空格的字段
func checkJsonWhiteSpace(str string) error {
	b := []byte(str)
	var r interface{} //结构未知，则解码接受变量用一个空接口变量
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	//解码完毕之后，r应该是map[string]interface{}类型变量，所以要进行类型断言
	data, ok := r.(map[string]interface{})
	if ok {
		//遍历所有成员，进行类型判断
		for key, val := range data {
			l := len(key)
			if l > 0 && (string(key[0]) == " " || string(key[l-1]) == " ") {
				return errors.New("Key: " + key + " . Contain white space")
			}
			switch tv := val.(type) { //注意，此处tv是一个明确了类型的值，而v是一个interface{}类型变量
			case string:
				l := len(tv)
				if l > 0 && (string(tv[0]) == " " || string(tv[l-1]) == " ") {
					return errors.New("Param: " + key + " . Contain white space")
				}
			case []interface{}:
				//fmt.Println(key, "is array:")
				for _, v1 := range tv {
					if s, ok := v1.(string); ok {
						//fmt.Println(k1, "is string", s)
						l := len(s)
						if l > 0 && (string(s[0]) == " " || string(s[l-1]) == " ") {
							return errors.New("Param: " + key + " Contain white space")
						}
					}
					if m, ok := v1.(map[string]interface{}); ok {
						b, err := json.Marshal(m)
						if err != nil {
							return err
						}
						err = checkJsonWhiteSpace(string(b))
						if err != nil {
							return err
						}
					}
					//fmt.Println(" ", k1, reflect.TypeOf(v1), v1)
				}
			case map[string]interface{}:
				//fmt.Println(key, "is struct:", tv)
				b, err := json.Marshal(tv)
				if err != nil {
					return err
				}
				err = checkJsonWhiteSpace(string(b))
				if err != nil {
					return err
				}
			default: //int, float, bool ...
				//IGNORE
				//fmt.Println("Ignore", key, reflect.TypeOf(val))
			}
		}
		//需要判断直接就是一个数组的情况啊，亲，怎么能忽略呢？
	} else if arrData, ok1 := r.([]interface{}); ok1 {
		//fmt.Println(key, "is array:")
		for _, v1 := range arrData {
			if s, ok := v1.(string); ok {
				//fmt.Println(k1, "is string", s)
				l := len(s)
				if l > 0 && (string(s[0]) == " " || string(s[l-1]) == " ") {
					return errors.New(fmt.Sprintf("Param %s or %s : Contain white space", s[0], s[l-1]))
				}
			}
			if m, ok := v1.(map[string]interface{}); ok {
				b, err := json.Marshal(m)
				if err != nil {
					return err
				}
				err = checkJsonWhiteSpace(string(b))
				if err != nil {
					return err
				}
			}
			//fmt.Println(" ", k1, reflect.TypeOf(v1), v1)
		}
	} else {
		return errors.New("Data format wrong")
	}

	return nil
}

//请求日志Middleware
func RequestLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*common.CustomContext)
		// clone 一个请求
		//cReq := cc.Request().Clone(cc.Request().Context())
		//body := cReq.Body

		body, _ := ioutil.ReadAll(cc.Request().Body)
		if len(body) != 0 {
			cc.Body = body
		} else {
			cc.Body = make([]byte, 0)
		}

		//*http.Request.Body 是 io.ReadCloser 类型，只能一次性读取完整，第二次就是空的
		//所以,在读到了body之后,需要再设置回来,否则action逻辑就无法再获取body了
		cc.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

		//cc.NewLogContextAccess("access message").
		//	Set("URI", cc.Request().URL.String()).
		//	Set("Param", string(body)).
		//	Print()

		fmt.Println("access message", "URI", cc.Request().URL.String(), "Param", string(body))
		return next(c)
	}
}

func CommonParamCheck(before echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*common.CustomContext)

		body, _ := ioutil.ReadAll(cc.Request().Body)
		cc.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

		if len(body) > 0 {
			if typ, ok := cc.Request().Header["Content-Type"]; ok {
				jsonFlag := false
				for _, t := range typ {
					if t == "application/json" {
						jsonFlag = true
						break
					}
				}
				if jsonFlag {
					err := checkJsonWhiteSpace(string(body))
					if err != nil {
						return cc.SendRsp(common.ERR_PARAM, err.Error())
					}
				}
			}
		}

		fmt.Println("CommonParamCheck func")
		return before(c)
	}
}

/**
 * @note
 * 运行选项
 */
func Flagset() *flag.FlagSet {
	flagSet := flag.NewFlagSet("disf-admin", flag.ExitOnError)

	flagSet.String("config", "", "path to config file")

	return flagSet
}

func getLogPath() string {
	_, filename, _, ok1 := runtime.Caller(0)
	if !ok1 {
		panic("No caller information")
	}
	dir := path.Dir(filename)

	logPath := dir + "/log"
	return logPath
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//解析运行参数 & 配置文件

	//新建Echo实例
	e := echo.New()

	//Db初始化
	engine, err := InitDb("root:123456@tcp(127.0.0.1:3306)/pushtx?charset=utf8", getLogPath()+"/xorm_log")
	if err != nil {
		panic("init db fail:" + err.Error())
	}

	// redisPool初始
	//redisPool := InitRedisPool(conf)

	//request id
	e.Use(middleware.RequestID())

	//创建自定义Context,将logger包入
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &common.CustomContext{
				Context:   c,
				DbEngine:  engine,
				ActionId:  c.QueryParam("actionId"),
				RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
				StartTime: time.Now(),
			}
			fmt.Println("new context")
			return h(cc)
		}
	})

	//recover
	e.Use(PanicRecover)

	//请求日志
	e.Use(RequestLog)

	//检查通用参数
	e.Use(CommonParamCheck)

	//request处理时间的统计以及返回值的统计
	e.Use(RequestLantency)

	//注册服务相关接口
	e.GET("/test/group/get", controller.GetGroups)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	//启动server
	e.Logger.Fatal(e.Start(":30000"))
}

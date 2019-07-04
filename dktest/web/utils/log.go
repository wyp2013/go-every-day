package utils
//
//import (
//	"fmt"
//	"path/filepath"
//	"git.xiaojukeji.com/disf/dlog"
//)
//
//const DefaultLogDir = "logs"
//
//type LogLevel string
//
//const (
//	LEVEL_DEBUG   LogLevel = "DEBUG"
//	LEVEL_ACCESS  LogLevel = "ACCESS"
//	LEVEL_INFO    LogLevel = "INFO"
//	LEVEL_ERROR   LogLevel = "ERR"
//)
//
////Disf logger句柄接口
//type DisfLogger interface {
//	Info(m map[string]interface{})
//	Error(m map[string]interface{})
//	Debug(m map[string]interface{})
//	LogDepth(s LogLevel, depth int, metaList []*dlog.LogMeta)
//	//Infof和Errorf兼容SetupFromDefaultDir
//	Errorf(format string, args ...interface{})
//	Infof(format string, args ...interface{})
//}
//
////dlogWrapper实现了DisfLogger接口
////封装了一个logger用于打印日志
//type dlogWrapper struct {
//	logger *dlog.Logger
//}
//
//var levelMap = map[LogLevel]dlog.Severity {
//	LEVEL_DEBUG: dlog.DEBUG,
//	LEVEL_ACCESS:  dlog.ACCESS,
//	LEVEL_INFO:  dlog.INFO,
//	LEVEL_ERROR: dlog.ERROR,
//}
//
//func (d *dlogWrapper) LogDepth(s LogLevel, depth int, metaList []*dlog.LogMeta) {
//	if level, ok := levelMap[s]; ok {
//		d.logger.LogSliceDepth(level, depth, metaList)
//		return
//	}
//	d.logger.LogSliceDepth(dlog.DEBUG, depth, metaList)
//}
//
//func (d *dlogWrapper) Info(m map[string]interface{}) {
//	d.logger.Info(m)
//}
//func (d *dlogWrapper) Error(m map[string]interface{}) {
//	d.logger.Error(m)
//}
//func (d *dlogWrapper) Debug(m map[string]interface{}) {
//	d.logger.Debug(m)
//}
////Infof和Errorf兼容SetupFromDefaultDir
//func (d *dlogWrapper) Infof(format string, args ...interface{}) {
//	d.logger.Infof(format, args...)
//}
//func (d *dlogWrapper) Errorf(format string, args ...interface{}) {
//	d.logger.Errorf(format, args...)
//}
//
//func NewDisfLogger(level LogLevel, dir string) (DisfLogger, error) {
//	if len(dir) == 0 {
//		dir = DefaultLogDir
//	}
//
//	backend, err := dlog.NewFileBackend(dir)
//	if err != nil {
//		abs, _ := filepath.Abs(dir)
//		fmt.Printf("Failed to setup log, dir=[%s] err=[%s]", abs, err.Error())
//		return nil, err
//	}
//
//	//日志切割不支持按小时,这块要改成依赖Odin做外部切割
//	//backend.SetRotateByHour(true)
//	//backend.SetKeepHours(24 * 3)
//
//	dlevel := dlog.INFO
//	switch level {
//	case LEVEL_DEBUG:
//		dlevel = dlog.DEBUG
//	case LEVEL_ACCESS:
//		dlevel = dlog.ACCESS
//	case LEVEL_ERROR:
//		dlevel = dlog.ERROR
//	}
//
//	return &dlogWrapper{logger: dlog.NewLogger(dlevel, backend)}, nil
//}
//
////LogContext接口,便于CustomContext打印日志
//type LogContext interface {
//	Get(key string) interface{}
//	Set(key string, value interface{}) LogContext
//	Print()
//}
//
////disfLogContext实现了LogContext接口
//type disfLogContext struct {
//	level  LogLevel
//	depth  int
//	logger DisfLogger
//	msg    string
//	data   map[string]*dlog.LogMeta
//	metaList []*dlog.LogMeta
//}
//
//func NewLogContext(s LogLevel, msg string, logger DisfLogger) LogContext {
//	return &disfLogContext{
//		level:    s,
//		msg:      msg,
//		data:     make(map[string]*dlog.LogMeta),
//		metaList: make([]*dlog.LogMeta, 0),
//		logger:   logger,
//	}
//}
//
//func NewLogContextDepth(s LogLevel, depth int, msg string, logger DisfLogger) LogContext {
//	return &disfLogContext{
//		level:    s,
//		depth:    depth,
//		msg:      msg,
//		data:     make(map[string]*dlog.LogMeta),
//		metaList: make([]*dlog.LogMeta, 0),
//		logger:   logger,
//	}
//}
//
//
//func (d *disfLogContext) Get(key string) interface{} {
//	return d.data[key]
//}
//
//func (d *disfLogContext) Set(key string, value interface{}) LogContext {
//	if _, ok := d.data[key]; !ok {
//		meta := &dlog.LogMeta{
//			Key:key,
//			Value:value,
//		}
//
//		d.data[key] = meta
//
//		d.metaList = append(d.metaList, meta)
//	} else {
//		d.data[key].Value = value
//	}
//
//	return d
//}
//
//func (d *disfLogContext) Print() {
//	d.Set("DISFMSG", d.msg)
//	d.logger.LogDepth(d.level, d.depth + 1, d.metaList)
//}
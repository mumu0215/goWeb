package myvalidate

import (
	"github.com/gin-gonic/gin/binding"
	"net"
	"reflect"
	"strconv"
	"strings"
	"sync"
	//"github.com/go-playground/validator/v10"
	"gopkg.in/go-playground/validator.v9"
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}
var _ binding.StructValidator = &DefaultValidator{}

func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}
func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// add any custom validations etc. here
		_ = v.validate.RegisterValidation("ip-port", validateIpPort)

	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func validateIpPort(fl validator.FieldLevel) bool {
	var (
		isValidIp   bool
		isValidPort bool
	)
	t := fl.Field().String()

	temp := strings.Split(t, ":")
	if len(temp) != 2 {
		return false
	}
	ip := net.ParseIP(temp[0])
	if ip != nil && ip.To4() != nil {
		isValidIp = true
	}
	port, err := strconv.Atoi(temp[1])
	if err != nil {
		return false
	}
	if port > 0 && port < 65536 {
		isValidPort = true
	}

	return isValidIp && isValidPort
}
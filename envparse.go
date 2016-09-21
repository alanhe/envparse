package envparse

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Version: 0.1.0

// A Param is the representation of a parameter.
type Param struct {
	name         string
	required     bool
	defaultValue string
}

// EnvParser the parser.
type EnvParser struct {
	params []*Param
	envs   map[string]string
}

func (e *EnvParser) paramValid(param *Param) bool {
	return len(param.name) > 0
}

func (e *EnvParser) paramExists(param *Param) bool {
	for _, p := range e.params {
		if p.name == param.name {
			return true
		}
	}
	return false
}

// Add an environment argument.
func (e *EnvParser) Add(param *Param) {
	if !e.paramValid(param) {
		panic(fmt.Errorf("Invalid param %v!", *param))
	}
	if e.paramExists(param) {
		panic(fmt.Errorf("Duplicated param: %s!", param.name))
	}
	e.params = append(e.params, param)
}

// Parse panics when the enviroment variables are not set correctly.
func (e *EnvParser) Parse(param *Param) {
	e.envs = make(map[string]string)
	errs := []error{}

	for _, p := range e.params {
		val := strings.TrimSpace(os.Getenv(p.name))
		if len(val) == 0 {
			val = p.defaultValue
		}
		if p.required && len(val) == 0 {
			errs = append(errs, fmt.Errorf("param %s is missing", p.name))
		} else {
			e.envs[p.name] = val
		}
	}

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		panic("Invalid env params!")
	}
}

func (e *EnvParser) isGetAllowed() {
	if e.envs == nil {
		panic("EnvParser#Parse() has not called!")
	}
}

// GetString gets the environment variable as string.
func (e *EnvParser) GetString(name string) string {
	e.isGetAllowed()
	return e.envs[name]
}

// GetInt gets the environment variable as int.
func (e *EnvParser) GetInt(name string) int {
	e.isGetAllowed()
	val, ok := e.envs[name]
	if !ok {
		return 0
	}
	ret, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return ret
}

// New EnvParser.
func New() *EnvParser {
	return &EnvParser{
		params: []*Param{},
	}
}

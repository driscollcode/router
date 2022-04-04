package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type Call interface {
	Request
	Error(response ...interface{}) Call
	PermanentRedirect(destination string) Call
	Redirect(destination string) Call
	Response(response ...interface{}) Call
	SetHeader(key, val string)
	Success(response ...interface{}) Call
}

type call struct {
	request
	response Response
}

func (c *call) SetHeader(key, val string) {
	if len(c.response.Headers) < 1 {
		c.response.Headers = make(map[string]string)
	}
	c.response.Headers[key] = val
}

func (c *call) Redirect(destination string) Call {
	c.response.StatusCode = http.StatusFound
	c.response.Redirect.DoRedirect = true
	c.response.Redirect.Destination = destination
	return c
}

func (c *call) PermanentRedirect(destination string) Call {
	c.response.StatusCode = http.StatusMovedPermanently
	c.response.Redirect.DoRedirect = true
	c.response.Redirect.Destination = destination
	return c
}

func (c *call) Success(response ...interface{}) Call {
	c.response.StatusCode = c.getStatusCode(http.StatusOK, response...)
	c.response.Content = append(c.response.Content, c.getResponseBody(response...)...)
	return c
}

func (c *call) Error(response ...interface{}) Call {
	c.response.StatusCode = c.getStatusCode(http.StatusBadRequest, response...)
	c.response.Content = append(c.response.Content, c.getResponseBody(response...)...)
	return c
}

func (c *call) Response(response ...interface{}) Call {
	if c.response.StatusCode < 1 {
		c.response.StatusCode = c.getStatusCode(http.StatusOK, response...)
	}
	c.response.Content = append(c.response.Content, c.getResponseBody(response...)...)
	return c
}

func (c *call) getStatusCode(defaultCode int, parts ...interface{}) int {
	if len(parts) < 1 {
		return defaultCode
	}

	userCode, err := strconv.Atoi(fmt.Sprintf("%v", parts[0]))
	if err != nil {
		return defaultCode
	}

	if userCode >= 100 && userCode <= 999 {
		return userCode
	}
	return defaultCode
}

func (c *call) getResponseBody(response ...interface{}) []byte {
	if response == nil {
		return nil
	}

	output := make([]byte, 0)
	for pos, piece := range response {
		if pos == 0 {
			if asInt, ok := piece.(int); ok && asInt >= 100 && asInt <= 999 {
				continue
			}
		}

		if len(output) > 0 {
			output = append(output, []byte("")...)
		}
		output = append(output, c.getContentAsByte(piece)...)
	}
	return output
}

func (c *call) getContentAsByte(content interface{}) []byte {
	switch reflect.ValueOf(content).Kind() {
	case reflect.Struct:
		if _, ok := content.(time.Time); ok {
			return []byte(fmt.Sprintf("%s", content.(time.Time).Format("2006-01-02 15:04:05")))
		}

		if bytes, err := json.Marshal(content); err == nil {
			return bytes
		}
	case reflect.Bool:
		return []byte(fmt.Sprint(content))
	case reflect.Slice:
		if bytes, ok := content.([]byte); ok {
			return bytes
		}
		return []byte(fmt.Sprint(content))
	case reflect.String:
		return []byte(content.(string))
	case reflect.Int:
		return []byte(strconv.Itoa(content.(int)))
	case reflect.Float64:
		return []byte(strconv.FormatFloat(content.(float64), 'f', -1, 64))
	}
	return nil
}

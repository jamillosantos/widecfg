package widecfg

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	ErrKeyNotFound    = errors.New("key not found")
	ErrValueWrongType = errors.New("value with wrong type")
)

type Config map[string]interface{}

var DefaultTimeFormat = time.RFC3339

// Get will find a given `key` into the Config.
func (c *Config) Get(key string) (interface{}, bool) {
	props := strings.Split(key, ".")
	value := (*c)
	for _, propName := range props {
		propValue, ok := value[propName]
		if !ok {
			break
		}
		switch propValueTyped := propValue.(type) {
		case map[string]interface{}:
			value = propValueTyped
		default:
			return propValueTyped, true
		}
	}
	// Could not find the Config.
	return nil, false
}

func (c *Config) GetString(key string) (string, error) {
	value, ok := c.Get(key)
	if !ok {
		return "", ErrKeyNotFound
	}
	switch typedValue := value.(type) {
	case string:
		return typedValue, nil
	case *string:
		return *typedValue, nil
	default:
		return "", ErrValueWrongType
	}
}

func (c *Config) GetInt(key string) (int, error) {
	value, ok := c.Get(key)
	if !ok {
		return 0, ErrKeyNotFound
	}
	switch typedValue := value.(type) {
	case int:
		return typedValue, nil
	case *int:
		return *typedValue, nil
	case string:
		return strconv.Atoi(typedValue)
	case *string:
		return strconv.Atoi(*typedValue)
	default:
		return 0, ErrValueWrongType
	}
}

func (c *Config) GetInt64(key string) (int64, error) {
	value, ok := c.Get(key)
	if !ok {
		return 0, ErrKeyNotFound
	}
	switch typedValue := value.(type) {
	case int64:
		return typedValue, nil
	case *int64:
		return *typedValue, nil
	case string:
		return strconv.ParseInt(typedValue, 10, 64)
	case *string:
		return strconv.ParseInt(*typedValue, 10, 64)
	default:
		return 0, ErrValueWrongType
	}
}

func (c *Config) GetUint(key string) (uint, error) {
	value, ok := c.Get(key)
	if !ok {
		return 0, ErrKeyNotFound
	}
	switch typedValue := value.(type) {
	case uint:
		return typedValue, nil
	case *uint:
		return *typedValue, nil
	case string:
		typedValueInt, err := strconv.ParseUint(typedValue, 10, 32)
		return uint(typedValueInt), err
	case *string:
		typedValueInt, err := strconv.ParseUint(*typedValue, 10, 32)
		return uint(typedValueInt), err
	default:
		return 0, ErrValueWrongType
	}
}

func (c *Config) GetUint64(key string) (uint64, error) {
	value, ok := c.Get(key)
	if !ok {
		return 0, ErrKeyNotFound
	}
	switch typedValue := value.(type) {
	case uint64:
		return typedValue, nil
	case *uint64:
		return *typedValue, nil
	case string:
		typedValueInt, err := strconv.ParseUint(typedValue, 10, 64)
		return uint64(typedValueInt), err
	case *string:
		typedValueInt, err := strconv.ParseUint(*typedValue, 10, 64)
		return uint64(typedValueInt), err
	default:
		return 0, ErrValueWrongType
	}
}

func (c *Config) GetFloat32(key string) (float32, error) {
	value, ok := c.Get(key)
	if !ok {
		return 0, ErrKeyNotFound
	}
	switch typedValue := value.(type) {
	case float32:
		return typedValue, nil
	case *float32:
		return *typedValue, nil
	case string:
		typedValueInt, err := strconv.ParseFloat(typedValue, 32)
		return float32(typedValueInt), err
	case *string:
		typedValueInt, err := strconv.ParseFloat(*typedValue, 32)
		return float32(typedValueInt), err
	default:
		return 0, ErrValueWrongType
	}
}

func (c *Config) GetFloat64(key string) (float64, error) {
	value, ok := c.Get(key)
	if !ok {
		return 0, ErrKeyNotFound
	}
	switch typedValue := value.(type) {
	case float64:
		return typedValue, nil
	case *float64:
		return *typedValue, nil
	case string:
		typedValueInt, err := strconv.ParseFloat(typedValue, 64)
		return float64(typedValueInt), err
	case *string:
		typedValueInt, err := strconv.ParseFloat(*typedValue, 64)
		return float64(typedValueInt), err
	default:
		return 0, ErrValueWrongType
	}
}

func (c *Config) GetBool(key string) (bool, error) {
	value, ok := c.Get(key)
	if !ok {
		return false, ErrKeyNotFound
	}
	switch typedValue := value.(type) {
	case bool:
		return typedValue, nil
	case *bool:
		return *typedValue, nil
	case string:
		return strconv.ParseBool(typedValue)
	case *string:
		return strconv.ParseBool(*typedValue)
	default:
		return false, ErrValueWrongType
	}
}

func (c *Config) GetTime(key string) (time.Time, error) {
	value, ok := c.Get(key)
	if !ok {
		return time.Time{}, ErrKeyNotFound
	}
	switch typedValue := value.(type) {
	case time.Time:
		return typedValue, nil
	case *time.Time:
		return *typedValue, nil
	case string:
		return time.Parse(DefaultTimeFormat, typedValue)
	case *string:
		return time.Parse(DefaultTimeFormat, *typedValue)
	default:
		return time.Time{}, ErrValueWrongType
	}
}

func (c *Config) GetDuration(key string) (time.Duration, error) {
	value, ok := c.Get(key)
	if !ok {
		return 0, ErrKeyNotFound
	}
	switch typedValue := value.(type) {
	case time.Duration:
		return typedValue, nil
	case *time.Duration:
		return *typedValue, nil
	case int:
		return time.Duration(typedValue), nil
	case *int:
		return time.Duration(*typedValue), nil
	case int64:
		return time.Duration(typedValue), nil
	case *int64:
		return time.Duration(*typedValue), nil
	case string:
		return time.ParseDuration(typedValue)
	case *string:
		return time.ParseDuration(*typedValue)
	default:
		return 0, ErrValueWrongType
	}
}

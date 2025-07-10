package utility

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"math/rand"
	"strings"
	"time"
)

func InArray[T int | string | int64 | int32 | uint](needle T, haystack []T) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
}

// ArrayDiff 差集
func ArrayDiff[T int | string | int64](dst, src []T) []T {
	m := make(map[T]bool)
	for _, item := range src {
		m[item] = true
	}
	var diff []T
	for _, item := range dst {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}

// ArrayUnique 唯一
func ArrayUnique[T int | string | int64 | int32 | uint](src []T) []T {
	if len(src) == 0 {
		return src
	}
	m := make(map[T]bool)
	var uniqueArr []T

	for _, item := range src {
		if _, ok := m[item]; !ok {
			m[item] = true

			uniqueArr = append(uniqueArr, item)
		}
	}
	return uniqueArr
}

// ArrayIntersect 交集
func ArrayIntersect[T int | string | int64](slice1, slice2 []T) []T {
	elemMap := make(map[T]bool)
	for _, item := range slice2 {
		elemMap[item] = true
	}
	var intersect []T
	for _, item := range slice1 {
		if _, ok := elemMap[item]; ok {
			intersect = append(intersect, item)
		}
	}
	return intersect
}

// ArrayChunk 分块
func ArrayChunk[T int | string | int64 | interface{}](slice1 []T, size int) [][]T {
	var chunks [][]T
	length := len(slice1)
	for i := 0; i < length; i += size {
		end := i + size
		if end > length {
			end = len(slice1)
		}
		chunks = append(chunks, slice1[i:end])
	}
	return chunks
}

// SearchMapKeyByValue 根据map的值值查询key
func SearchMapKeyByValue[T int | string | int64 | int32, T1 string](value T1, dataMap map[T]T1) (key T, found bool) {
	for key, val := range dataMap {
		if val == value {
			return key, true
		}
	}
	return
}

func CopyFields(src interface{}, to interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, to)
	if err != nil {
		return err
	}
	return nil
}

func RedisLock(ctx context.Context, cache *gredis.Redis, key string, value interface{}, timeout int64) bool {
	do, err := cache.Do(ctx, "set", key, value, "nx", "ex", timeout)
	if err != nil {
		g.Log().Errorf(ctx, "redis lock err, err=%+v", err)
		return false
	}
	return do.Bool()
}

// Snake2Camel snake转驼峰
func Snake2Camel(s string) string {
	var camelString string

	// Split the string by underscores
	words := strings.Split(s, "_")

	// Capitalize the first letter of each word
	for j, word := range words {
		words[j] = strings.Title(word)
	}

	// Join the words with spaces
	camelString = strings.Join(words, "")
	return camelString
}

func RandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	lettersLen := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(lettersLen)]
	}
	return string(b)
}

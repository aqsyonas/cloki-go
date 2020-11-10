package service

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/patrickmn/go-cache"
	"sort"
)

type LabelService struct {
	ServiceData
	GoCache *cache.Cache
}

func (ls *LabelService) GetLabels() string {

	if keys, exist := ls.GoCache.Get("__LABEL__"); exist {
		labelKeys := keys.([]string)
		reply := gabs.New()
		reply.Set("success", "status")
		sort.Slice(labelKeys, func(i, j int) bool {
			return labelKeys[i] < labelKeys[j]
		})
		reply.Set(labelKeys, "data")
		return reply.String()
	}
	reply := gabs.New()
	reply.Set("success", "status")
	reply.Set([]string{"__name__"}, "data")
	return reply.String()
}

func (ls *LabelService) LabelValsByKey(key string) string {

	if keys, exist := ls.GoCache.Get(key); exist {
		labelKeys := keys.([]string)
		reply := gabs.New()
		reply.Set("success", "status")
		sort.Slice(labelKeys, func(i, j int) bool {
			return labelKeys[i] < labelKeys[j]
		})
		reply.Set(labelKeys, "data")
		return reply.String()
	}
	reply := gabs.New()
	reply.Set("success", "status")
	reply.Set([]string{}, "data")
	return reply.String()
}

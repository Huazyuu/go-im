package umap

import "reflect"

// RefStructByTag 将结构体通过tag映射成 map[key]value
func RefStructByTag(data any, tag string) map[string]any {
	maps := make(map[string]any)
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		getTag, ok := field.Tag.Lookup(tag)
		if !ok {
			continue
		}
		val := v.Field(i)
		if val.IsZero() {
			continue
		}
		// 结构体类型 递归继续查找子类
		if field.Type.Kind() == reflect.Struct {
			newMap := RefStructByTag(val.Elem().Interface(), tag)
			maps[getTag] = newMap
			continue
		}
		// 指针类型
		if field.Type.Kind() == reflect.Ptr {
			if field.Type.Elem().Kind() == reflect.Struct {
				newMaps := RefStructByTag(val.Elem().Interface(), tag)
				maps[getTag] = newMaps
				continue
			}
			maps[getTag] = val.Elem().Interface()
			continue
		}
		maps[getTag] = val.Interface()
	}
	return maps
}

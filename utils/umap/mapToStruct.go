package umap

import "reflect"

// MapToStruct 函数用于将一个 map[string]any 类型的映射数据转换到指定的结构体指针指向的结构体中
// 该函数会根据结构体字段的 json 标签来匹配映射中的键，并将对应的值赋值给结构体的字段
// 参数 data 是包含键值对的映射，用于提供要赋值的数据
// 参数 ptr 是指向目标结构体的指针，函数会修改该指针指向的结构体内容
func MapToStruct(data map[string]any, ptr any) {
	// 获取指针指向的结构体的类型信息
	// reflect.TypeOf(ptr) 返回的是指针类型，使用 Elem() 方法获取指针指向的实际类型（即结构体类型）
	t := reflect.TypeOf(ptr).Elem()
	// 获取指针指向的结构体的实际值
	// reflect.ValueOf(ptr) 返回的是指针的值，使用 Elem() 方法获取指针指向的实际值（即结构体的值）
	v := reflect.ValueOf(ptr).Elem()

	// 遍历结构体的每个字段
	for i := 0; i < t.NumField(); i++ {
		// 获取当前字段的类型信息
		field := t.Field(i)
		// 从当前字段的标签中获取 json 标签的值
		tag := field.Tag.Get("json")
		// 如果 json 标签为空或者是 "-"，表示该字段不需要处理，跳过当前字段
		if tag == "" || tag == "-" {
			continue
		}

		// 尝试从映射中获取与 json 标签对应的键的值
		mapField, ok := data[tag]
		// 如果映射中不存在该键，跳过当前字段
		if !ok {
			continue
		}

		// 获取当前结构体字段的值
		val := v.Field(i)

		// 检查当前字段的类型是否为指针类型
		if field.Type.Kind() == reflect.Ptr {
			// 如果是指针类型，进一步检查指针指向的元素的类型
			switch field.Type.Elem().Kind() {
			// 当指针指向的元素类型为字符串类型时
			case reflect.String:
				// 获取映射中对应键的值的反射值
				mapFieldValue := reflect.ValueOf(mapField)
				// 检查映射中对应键的值的类型是否为字符串类型
				if mapFieldValue.Type().Kind() == reflect.String {
					// 将映射中的字符串值进行类型断言并赋值给 strVal
					strVal := mapField.(string)
					// 创建一个指向 strVal 的指针，并将该指针赋值给结构体的当前字段
					val.Set(reflect.ValueOf(&strVal))
				}
			}
		}
	}
}

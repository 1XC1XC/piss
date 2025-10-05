package Map

type Any map[any]any

func (Array Any) Set(Key any, Value any) {
	Array[Key] = Value
}

func (Array Any) Get(Key any) any {
	return Array[Key]
}

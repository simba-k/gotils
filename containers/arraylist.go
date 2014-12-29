package containers

type ArrayList struct {
	dat []interface{}
}

func NewArrayList(capacity int) ArrayList {
	list := ArrayList{
		make([]interface{}, 0, capacity),
	}
	return list
}

func (arr *ArrayList) Append(appention interface{}) {
	arr.dat = append(arr.dat, appention)
}

func (arr *ArrayList) Prepend(appention interface{}) {
	arr.dat = arr.dat[:len(arr.dat)+1]
	if len(arr.dat) > 1 {
		copy(arr.dat[1:len(arr.dat)], arr.dat[0:len(arr.dat)-1])
	}
	arr.dat[0] = appention
}

func (arr *ArrayList) Insert(insertion interface{}, idx int) {
	arr.dat = arr.dat[:len(arr.dat)+1]
	l := len(arr.dat)
	if len(arr.dat) > 1 {
		copy(arr.dat[idx+1:l], arr.dat[idx:l-1])
	}
	arr.dat[idx] = insertion
}

func (arr *ArrayList) Remove(idx int) {
	arr.dat = arr.dat[:len(arr.dat)+1]
	l := len(arr.dat)
	if len(arr.dat) > 1 {
		copy(arr.dat[idx+1:l], arr.dat[idx:l-1])
	}
	arr.dat[idx] = insertion
}

func (arr *ArrayList) Get(idx int) interface{} {
	return arr.dat[idx]
}

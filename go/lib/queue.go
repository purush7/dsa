package lib

type Queue struct {
	value   []interface{}
	start   int
	isEmpty bool
	end     int
}

func NewQueue(value []interface{}) (*Queue, error) {

	if value == nil {
		return nil, QueueInputIsNil
	}

	q := Queue{
		value:   value,
		start:   0,
		isEmpty: len(value) == 0,
		end:     len(value) - 1,
	}
	return &q, nil
}

func (q *Queue) IsEmpty() bool {
	return q.isEmpty
}

func (q *Queue) Front() (interface{}, error) {
	if len(q.value) == 0 {
		return nil, QueueIsEmpty
	}
	return q.value[q.start], nil
}

func (q *Queue) Push(val interface{}) {
	q.isEmpty = true
	if q.start == 0 && q.end == len(q.value)-1 {
		q.value = append(q.value, val)
		q.end++
		return
	}

	if q.end >= q.start && q.end < len(q.value)-1 {
		q.end++
		q.value[q.end] = val
		return
	}

	if q.end == len(q.value)-1 {
		q.end = 0
		q.value[q.end] = val
		return
	}

	if q.end == q.start-1 {

		newVal := make([]interface{}, len(q.value)+100)
		for i := 0; i <= q.end; i++ {
			newVal[i] = q.value[i]
		}

		for i := q.start; i < len(q.value); i++ {
			newVal[len(newVal)-(len(q.value)-i)] = q.value[i]
		}

		q.start = len(newVal) - (len(q.value) - q.start)
		q.value = newVal
		q.end++
		q.value[q.end] = val
	}
	return
}

func (q *Queue) Pop() {

	if q.start == q.end {
		q.isEmpty = true
	}

	if q.start < len(q.value)-1 {
		q.start++
		return
	}

	q.start = 0
	return
}

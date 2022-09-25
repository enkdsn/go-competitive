package queue_test

import (
	"testing"

	"github.com/enkdsn/go-competitive/queue"
	"github.com/stretchr/testify/assert"
)

func TestCircularQueue(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		assert := assert.New(t)

		q := queue.NewQueue[int](3)

		assert.True(q.EnQueue(1))
		assert.True(q.EnQueue(2))
		assert.True(q.EnQueue(3))

		assert.False(q.EnQueue(3))

		r, _ := q.Rear()
		assert.Equal(3, r)

		assert.True(q.IsFull())
		_, ok := q.DeQueue()

		assert.True(ok)
		assert.True(q.EnQueue(4))
		r, _ = q.Rear()
		assert.Equal(4, r)
	})

	t.Run("case2", func(t *testing.T) {
		assert := assert.New(t)

		q := queue.NewQueue[int](6)

		assert.True(q.EnQueue(6))

		r, _ := q.Rear()
		assert.Equal(6, r)

		r, _ = q.Rear()
		assert.Equal(6, r)

		_, ok := q.DeQueue()
		assert.True(ok)

		assert.True(q.EnQueue(5))

		r, _ = q.Rear()
		assert.Equal(5, r)

		q.DeQueue()

		r, _ = q.Front()
		assert.Equal(-1, r)
	})
}

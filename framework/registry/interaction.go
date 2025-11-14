package registry

type InteractionItem interface {
	GetId() string
}

type InteractionSystem[T InteractionItem] struct {
	Queue         map[string][]T
	QueueSnapshot map[string][]T
}

func NewInteractionSystem[T InteractionItem]() *InteractionSystem[T] {
	return &InteractionSystem[T]{
		Queue:         make(map[string][]T),
		QueueSnapshot: make(map[string][]T),
	}
}

func (i *InteractionSystem[T]) UpdateInteractionSystemSnapshot() {
	for id, interactions := range i.Queue {
		i.QueueSnapshot[id] = interactions
	}
	i.Queue = make(map[string][]T)
}

func (i *InteractionSystem[T]) AddInteraction(id string, item T) {
	i.Queue[id] = append(i.Queue[id], item)
}

func (i *InteractionSystem[T]) GetInteractions(id string) ([]T, bool) {
	if _, ok := i.QueueSnapshot[id]; ok {
		return i.QueueSnapshot[id], true
	}
	return make([]T, 0), false
}

func (i *InteractionSystem[T]) Reset() {
	i.QueueSnapshot = make(map[string][]T)
}

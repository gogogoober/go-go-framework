package interactionsystem

type InteractionItem interface {
	GetId() string
}

type InteractionSystem[T InteractionItem] struct {
	Queue map[string][]T
}

func NewInteractionSystem[T InteractionItem]() *InteractionSystem[T] {
	return &InteractionSystem[T]{
		Queue: make(map[string][]T),
	}
}

func (i *InteractionSystem[T]) AddInteraction(id string, item T) {
	i.Queue[id] = append(i.Queue[id], item)
}

func (i *InteractionSystem[T]) GetInteraction(id string) ([]T, bool) {
	if _, ok := i.Queue[id]; ok {
		return i.Queue[id], true
	}
	return make([]T, 0), false
}

func (i *InteractionSystem[T]) Reset() {
	i.Queue = make(map[string][]T)
}

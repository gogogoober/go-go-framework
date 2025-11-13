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

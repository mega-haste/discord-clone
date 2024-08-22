package SocketIO

import (
	"fmt"
	"slices"
)

type option[T any] struct {
	has_value bool
	value     T
}

func Some[T any](value T) option[T] {
	return option[T]{
		has_value: true,
		value:     value,
	}
}

func None[T any]() option[T] {
	return option[T]{
		has_value: false,
	}
}

func (s option[T]) IsSome() bool {
	return s.has_value
}

func (s option[T]) Unwrap() T {
	return s.value
}

type SocketEmiter struct {
	except []ID_t
	to     option[ID_t]
	server *SocketIO
}

func NewSocketEmiter(server *SocketIO) SocketEmiter {
	return SocketEmiter{
		except: make([]ID_t, 0),
		to:     None[ID_t](),
		server: server,
	}
}

func (s *SocketEmiter) To(ID ID_t) *SocketEmiter {
	s.to = Some(ID)
	return s
}

func (s *SocketEmiter) Except(ID ID_t) *SocketEmiter {
	s.except = append(s.except, ID)
	return s
}

func (s SocketEmiter) Emit(eventId string, message string) error {
	if s.to.IsSome() {
		ID := s.to.Unwrap()
		client, exists := s.server.clients[ID]
		if !exists {
			return fmt.Errorf("Client with id `%uz` doesn't exist.", ID)
		}
		client.Conn.WriteJSON(SocketMessageEvent {
			EventKey: eventId,
			Message: message,
		})
		return nil
	}
	for k, v := range s.server.clients {
		if slices.Contains(s.except, k) {
			continue
		}
		
		v.Conn.WriteJSON(SocketMessageEvent{
			EventKey: eventId,
			Message:  message,
		})
	}
	return nil
}


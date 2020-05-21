/*
 * Author: fasion
 * Created time: 2020-04-23 16:07:17
 * Last Modified by: fasion
 * Last Modified time: 2020-04-24 09:18:17
 */

package job

import "context"

type TicketAllocator struct {
	tickets    int
	ticketChan chan struct{}
}

func NewTicketAllocator(tickets int) *TicketAllocator {
	return &TicketAllocator{
		tickets:    tickets,
		ticketChan: make(chan struct{}, tickets),
	}
}

func (allocator *TicketAllocator) Total() int {
	return allocator.tickets
}

func (allocator *TicketAllocator) Allocateds() int {
	return len(allocator.ticketChan)
}

func (allocator *TicketAllocator) Acquire(ctx context.Context) {
	select {
	case allocator.ticketChan <- struct{}{}:
	case <-ctx.Done():
	}
}

func (allocator *TicketAllocator) Release() {
	select {
	case <-allocator.ticketChan:
	}
}

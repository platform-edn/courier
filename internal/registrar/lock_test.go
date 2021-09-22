package registrar

import "testing"

func TestNewTicketLock(t *testing.T) {
	tl := newTicketLock()

	if tl.ticket != 0 || tl.next != 0 {
		t.Fatal("ticket and next need to equal zero when instantiated")
	}
}

func TestTicketLock_Lock(t *testing.T) {
	tl := newTicketLock()
	defer tl.unlock()

	startTicket := tl.ticket
	startNext := tl.next
	tl.lock()
	lockTicket := tl.ticket
	lockNext := tl.next

	if lockTicket-startTicket != 1 {
		t.Fatal("ticket count did not increase by 1")
	}

	if startNext != lockNext {
		t.Fatalf("expected next to be %v but got %v", startNext, lockNext)
	}
}

func TestTicketLock_Unlock(t *testing.T) {
	tl := newTicketLock()

	tl.lock()
	lockTicket := tl.ticket
	lockNext := tl.next

	tl.unlock()
	unlockTicket := tl.ticket
	unlockNext := tl.next

	if lockTicket != unlockTicket {
		t.Fatalf("ticket should not increase after unlock. lockTicket: %v unlockTicket: %v", lockTicket, unlockTicket)
	}

	if unlockNext-lockNext != 1 {
		t.Fatalf("next should have increased by 1 but did not. lockNext: %v, unlockNext: %v", lockNext, unlockNext)
	}
}

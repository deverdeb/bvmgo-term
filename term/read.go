package term

import (
	"github.com/eiannone/keyboard"
)

type EventProcessor struct {
	started bool
}

type KeyboardEventHandler = func(event keyboard.KeyEvent, processor *EventProcessor)

func Read(handler KeyboardEventHandler) error {
	keyEventsChannel, err := keyboard.GetKeys(100)
	if err != nil {
		return err
	}
	defer func() {
		go keyboard.Close()
	}()

	var processor = EventProcessor{started: true}

	for {
		event := <-keyEventsChannel
		if event.Err != nil {
			return event.Err
		}
		handler(event, &processor)
		if !processor.started {
			return nil
		}
	}
}

func (processor *EventProcessor) Stop() {
	processor.started = false
}

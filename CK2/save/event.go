package save

type Event struct {
	Event string      `paradox_field:"event" json:"event,omitempty"`
	Days  int         `paradox_field:"days" json:"days,omitempty"`
	Scope *EventScope `paradox_field:"scope" json:"scope,omitempty"`
}

type EventScope struct {
	SavedEventTargets []*SavedEventTarget `paradox_field:"saved_event_target" paradox_type:"list" json:"saved_event_target,omitempty"`
	Seed              int                 `paradox_field:"seed" json:"seed,omitempty"`
	Char              int                 `paradox_field:"char" json:"char,omitempty"`
	From              *EventScope         `paradox_field:"from" json:"from,omitempty"`
	Root              *EventScope         `paradox_field:"root" json:"root,omitempty"`
}

type SavedEventTarget struct {
	Name string `paradox_field:"name" json:"name,omitempty"`
	Char int    `paradox_field:"char" json:"char,omitempty"`
}

type ActiveEvent struct {
	Type  string      `paradox_field:"type" json:"type,omitempty"`
	Scope *EventScope `paradox_field:"scope" json:"scope,omitempty"`
}

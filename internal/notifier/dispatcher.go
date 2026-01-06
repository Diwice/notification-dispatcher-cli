package notifier

import "not-dis/internal/templates"

type Dispatcher interface {
	Appoint(obj templates.Template) error
}

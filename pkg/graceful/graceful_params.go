package graceful

import "context"

type Params struct {
	OnStart    func()
	OnTimeout  func()
	OnShutdown func(timeoutCtx context.Context)
}

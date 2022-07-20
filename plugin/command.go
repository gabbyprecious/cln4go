package plugin

type rpcMethod[T any] struct {
	name            string 
	usage           string 
	description     string
	LongDescription string `json:"long_description"`
	callback        RPCCommand[T]
}

func (instance *rpcMethod[T]) Call(plugin *Plugin[T], request map[string]any) (map[string]any, error) {
	return instance.callback.Call(plugin, request)
}

type rpcNotification[T any] struct {
	onEvent  string
	callback RPCEvent[T]
}

func (instance *rpcNotification[T]) Call(plugin *Plugin[T], request map[string]any) {
	instance.callback.Call(plugin, request)
}

type rpcHook[T any] struct {
	name 	  	string 
	before 	   	[]string 
	after 	   	[]string 
	callback  	RPCCommand[T] 
}

func (instance *rpcHook[T]) Call(plugin *Plugin[T], request map[string]any) {
	instance.callback.Call(plugin, request)
}
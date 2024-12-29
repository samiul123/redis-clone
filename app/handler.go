package main

import "sync"

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}
var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{dataType: "string", str: "PONG"}
	}

	return Value{dataType: "string", str: args[0].bulk}
}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{dataType: "string", str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].bulk
	val := args[1].bulk

	SETsMu.Lock()
	SETs[key] = val
	SETsMu.Unlock()

	return Value{dataType: "string", str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{dataType: "string", str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].bulk
	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()

	if !ok {
		return Value{dataType: "null"}
	}

	return Value{dataType: "bulk", bulk: value}
}

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{dataType: "string", str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk
	val := args[2].bulk

	HSETsMu.Lock()
	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = val
	HSETsMu.Unlock()
	return Value{dataType: "string", str: "OK"}
}

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{dataType: "string", str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk

	HSETsMu.RLock()
	val, ok := HSETs[hash][key]
	HSETsMu.RUnlock()

	if !ok {
		return Value{dataType: "null"}
	}

	return Value{dataType: "bulk", bulk: val}
}

func hgetall(args []Value) Value {
	if len(args) != 1 {
		return Value{dataType: "string", str: "ERR wrong number of arguments for 'hgetall' command"}
	}

	hash := args[0].bulk

	HSETsMu.RLock()
	innerMap, ok := HSETs[hash]
	HSETsMu.RUnlock()

	if !ok {
		return Value{dataType: "null"}
	}

	items := make([]Value, 0, 2*len(innerMap))
	for key, val := range innerMap {
		items = append(items, Value{dataType: "bulk", bulk: key},
			Value{dataType: "bulk", bulk: val})
	}

	return Value{dataType: "array", array: items}
}

var handlers = map[string]func([]Value) Value{
	"PING":    ping,
	"SET":     set,
	"GET":     get,
	"HSET":    hset,
	"HGET":    hget,
	"HGETALL": hgetall,
}

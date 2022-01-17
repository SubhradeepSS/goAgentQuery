module example.com/main

go 1.17

replace example.com/agent => ../agent

replace example.com/issue => ../issue

require (
	example.com/agent v0.0.0-00010101000000-000000000000
	example.com/issue v0.0.0-00010101000000-000000000000
)

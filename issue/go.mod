module example.com/issue

go 1.17

replace example.com/agent => ../agent

replace example.com/queries => ../queries

require (
	example.com/agent v0.0.0-00010101000000-000000000000
	example.com/queries v0.0.0-00010101000000-000000000000
)

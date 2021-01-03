module run

go 1.15

replace example.com/rootmethods => ../rootmethods

replace example.com/optimization => ../optimization

require (
	example.com/optimization v0.0.0-00010101000000-000000000000 // indirect
	example.com/rootmethods v0.0.0-00010101000000-000000000000
)

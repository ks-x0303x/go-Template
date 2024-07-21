module TemplateApp

go 1.20

replace (
	Common/ExpansionString => ./Common/ExpansionString // パッケージ名 => project から pkg_orge までのパス
	Common/Logger => ./Common/Logger // パッケージ名 => project から pkg_orge までのパス
	DatabaseManager => ./DatabaseManager
	Models/users => ./Models/users
	Server => ./Server
)

require (
	Common/ExpansionString v1.0.0
	Common/Logger v0.0.0-00010101000000-000000000000
	DatabaseManager v1.0.0
	Server v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.8.1
	Models/users v1.8.0
)

require filippo.io/edwards25519 v1.1.0 // indirect

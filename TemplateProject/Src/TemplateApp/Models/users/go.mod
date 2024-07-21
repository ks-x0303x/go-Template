module Models/users

go 1.20

replace (
	Common/Logger => ./../../Common/Logger
	DatabaseManager => ./../../DatabaseManager
)

require (
	Common/Logger v0.0.0-00010101000000-000000000000
	DatabaseManager v1.0.0
)

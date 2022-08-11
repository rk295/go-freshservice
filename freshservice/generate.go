package freshservice

//go:generate -command gen go run ./generator/main.go

//go:generate gen --endpoint agent agent.gen.go
//go:generate gen --endpoint announcement announcement.gen.go
//go:generate gen --endpoint application application.gen.go
//go:generate gen --endpoint asset asset.gen.go
//go:generate gen --endpoint assetType --api-endpoint asset_types --json-key asset_types asset_type.gen.go
//go:generate gen --endpoint businessHour --api-endpoint business_hours --json-key business_hours business_hour.gen.go
//go:generate gen --endpoint department department.gen.go
//go:generate gen --endpoint location location.gen.go
//go:generate gen --endpoint requester requester.gen.go
//go:generate gen --endpoint serviceCatalog --api-endpoint service_catalog --json-key service_items service_catalog.gen.go
//go:generate gen --endpoint ticket ticket.gen.go
//go:generate gen --endpoint task task.gen.go

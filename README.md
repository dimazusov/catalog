## Simple catalog

Сatalog includes categories, organizations and buildings:
- category can include many organizations
- buildings can include many organizations

```bigquery
type Category struct {
    ID            uint
    ParentID      uint
    Name          string
}

type Building struct {
    ID            uint
    Address       string
    Coords        coords.Coords
}

type Organization struct {
    ID     uint                                  
    Name   string                                
    Phones []organization_phone.OrganizationPhone
}
```

## Docs
    api/swagger.json - swagger api

## Make commands
    make up - run application
    make test - run tests
    make lint - run linters
    make migrate - run migrations
    make swagger - update swagger docs to directory api
    make test-data - fill db tests cases
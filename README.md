##Simple catalog

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

## docs
    api/swagger.json - swagger api

## make commands
    make up - run application
    make test - run tests <br/>
    make lint - run linters <br/>
    make migrate - run migrations <br/>
    make swagger - update swagger docs to directory api <br/>
    make test-data - fill db tests cases <br/>
    Skill Test Saham Rakyat

        Migrations
            Command to run:
            goose -dir migrations/ mysql "<username>:<password>@tcp(<url>:<port>)/<db-name>?parseTime=true" up

            For rollback:
            goose -dir migrations/ mysql "<username>:<password>@tcp(<url>:<port>)/<db-name>?parseTime=true" down
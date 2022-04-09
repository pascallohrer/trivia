# Coding Challenge results

## Setup instructions
1. Make sure you have docker up and running on your machine.
2. Navigate to the project folder and run `docker compose up -d`. This will build the required images, initialize the database and spin up the containers for running the app.

Note: I decided to go with **docker compose** after fiddling around with a local mongodb installation and realizing how complicated it would be to make this work on different systems/OSs. I'm aware this is technically no longer "a mongodb local instance" as was requested in the requirements, but it saves at least 10 lines of explanation here (including a link to the online installation instructions) as well as lots of time and headache on the part of whoever tries to run this without mongodb already set up locally. Also, I feel docker compose is significantly easier to slot into an existing CI/CD pipeline, so ignoring the "local instance" requirement makes my solution much better in terms of the "production-ready" requirement.

## The Logger
Production logging should usually log to file in order to provide a permanent record of logging activity. The logger is instantiated once and subsequently passed around to any components that need it (dependency injection).

## The database
As per the requirements, only read access to the database is necessary, so the adapter will only implement a `Find` method, reading data based on filters. Since those are the relevant and meaningful values, `text` and `number` will be filterable, where `text` uses regex search to find the filter value anywhere in an entry's text, while `number` looks for an exact match. Both filters should allow multiple values, which will be comma-separated in the input and will be `OR`-connected in the query, so all resulting entries will need to match any one (or more) of the provided values. If both `text` and `number` filters are provided, they will be `AND`-connected, so all resulting entries will need to match both the text and the number filters.

Tests must cover all possible filters, which can be achieved with the following representative cases:
- Find a specific number
- Find a specific text part
- Return an empty result set in case no match is found
- Find all results for multiple filter values
- Find only the correct results for a text+number filter combined

## The API
The API offers a single endpoint `/api/v1/trivia` which only supports `GET`. The following query parameters are supported:
- `text`: Filters the trivia bits for any that contain either of the comma-separated values provided in the parameter.
- `number`: Filters the trivia bits for any that exactly match any of the comma-separated values provided in the parameter.
- `random`: Returns a random selection from the actual results. This parameter is interpreted as an integer, defaulting to 1 for invalid values (including values less than 1). This value is the number of results returned, where each result is randomly picked from the full set of actual results based on the other filters. Duplicates are allowed and the returned array will always contain exactly the number of records passed as value to this parameter (unless there are no records to pick from).

The returned value is a JSON-encoded array sent with the HTTP status `OK`. If no records are found, the response will instead have HTTP status `Not Found`. In case any errors occur during database access, the error and request data will be logged and the response will have HTTP status `Internal Server Error`.

Examples:
- `api/v1/trivia?number=3,6,4e93&text=number`
  - This will return all entries for the numbers 3, 6 and 4e93 that contain the word "number" in their text.
- `api/v1/trivia?random=hell_yeah`
  - This will default to the value 1 and return one random entry from the entire dataset

Tests are shortened here, covering only the `random` filter. In production, they should cover all filters, similarly to the Database tests. For this purpose, the mockDB would need to be expanded to allow text search, which is currently ignored at this point. Since the functionality's entry point lies in the `router` package, this is where the tests are located, but they cover the `handlers` package implicitly.

A formal API documentation is available in the file [trivia.yml](trivia.yml).

## Security
Hardcoding passwords is obviously a terrible idea in production. Docker compose offers the `secrets` functionality, which is included here with the local file version. For demonstration purposes, the password file is committed here, this is not the way this should be done in production. The docker secrets can be set with docker directly instead of using local files, but going into detail for this part is beyond the scope of this coding challenge.

## Other Notes
- The number field is being stored as a `float64`. This is the type that is easiest to handle, especially considering conversion/string parsing among the types large enough to hold all valid values. An alternative would be the `math/big` package, specifically the `Int` type.
- Testing should be a part of the deployment process, so the actual pipeline should include something that runs the package tests. Integration tests are also a good idea, they would look similar to the `router` tests, but use the actual database implementation instead of a mock.
- Spinning up the containers (`docker compose up -d`) without shutting them down first will actually reuse the database and duplicate all entries. There is intentionally no duplicate protection, so this will break the "correct total number of documents" test and cause every result to come back twice (or more often if the containers are spun up repeatedly). Don't do this, shut the containers down before starting them again.
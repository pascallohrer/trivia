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
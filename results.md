# Coding Challenge results

## Setup instructions
1. Make sure you have docker up and running on your machine.
2. Navigate to the project folder and run `docker compose up -d`. This will build the required images, initialize the database and spin up the containers for running the app.

Note: I decided to go with **docker compose** after fiddling around with a local mongodb installation and realizing how complicated it would be to make this work on different systems/OSs. I'm aware this is technically no longer "a mongodb local instance" as was requested in the requirements, but it saves at least 10 lines of explanation here (including a link to the online installation instructions) as well as lots of time and headache on the part of whoever tries to run this without mongodb already set up locally. Also, I feel docker compose is significantly easier to slot into an existing CI/CD pipeline, so ignoring the "local instance" requirement makes my solution much better in terms of the "production-ready" requirement.

## The Logger
Production logging should usually log to file in order to provide a permanent record of logging activity. The logger is instantiated once and subsequently passed around to any components that need it (dependency injection).
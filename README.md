# go-exercise

## Getting Started

### Running the app
```bash
$ go run cmd/server/main.go
```

### Tests
```bash
$ go test -v ./...
```

## Tasks

Each task should be completed in order and with test coverage.

### Add a zipcode field

Add a new field `ZipCode` to `domain.User` that's persisted by `memory.UserStore` and returned in the response of the `/v1/users` route.

Example output of `/v1/users`:

```json
{
    "users": [
        {
            "nameFirst": "John",
            "nameLast": "Smith",
            "email": "john.smith@gmail.com",
            "zipcode": "02210"
        }
    ]
}
```

You can make a POST request to `/v1/users` to create a new user, but will need to modify it to support a zipcode.

### User location

Modify the `/v1/user/{id}` route to incude a new field `location`. This field should be populated with the city and state code of the user's zipcode.

Use the `/search` endpoint from [Zipcodebase](https://app.zipcodebase.com/documentation) API to get a city and state code for a zipcode. You will receive an API key for Zipcodebase from Eleanor Health.

Example output for the Zipcodebase `/search` endpoint:

```json
{"query":{"codes":["02210"],"country":"us"},"results":{"02210":[{"postal_code":"02210","country_code":"US","latitude":"42.34890000","longitude":"-71.04650000","city":"Boston","state":"Massachusetts","city_en":"Boston","state_en":"Massachusetts","state_code":"MA","province":"Suffolk","province_code":"025"}]}}
```

Example output of `/v1/user/{id}`:

```json
{
    "nameFirst": "John",
    "nameLast": "Smith",
    "email": "john.smith@gmail.com",
    "location": "Boston, MA"
}
```

### Locations count

Add a new route `/v1/locations` to return a list of distinct locations. A location is an object made up of a distinct zip code, city, state code, and total count of users for the location.

Example output of `/locations` given three users with the zipcode `02210`, two users with the zipcode `06040`, and one user with the zipcode `92648`:

```json
{
    "zipcode": "02210",
    "location": "Boston, MA",
    "userCount": 3
},
{
    "zipcode": "06040",
    "location": "Manchester, CT",
    "userCount": 2
},
{
    "zipcode": "92648",
    "location": "Huntington Beach, CA",
    "userCount": 1
}
```

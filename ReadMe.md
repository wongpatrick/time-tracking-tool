# Time-Tracking-Tool

Time tracking tool for shift based data

## Description

This program reads a JSON file containing shift data and calculates various metrics for employees, such as regular hours, overtime hours, and identifies invalid shifts. The challenge involves handling shifts that cross midnight and span multiple weeks.

## Getting Started

### Dependencies

* [Go 1.21.0+ installed](https://go.dev/doc/install)

### Installing

1. Clone the repository:
    ```
    https://github.com/wongpatrick/time-tracking-tool.git
    ```
2. Change to the project directory:
    ```bash
    cd time-tracking-tool
    ```
3. Install packages:
    ```bash
    go mod vendor
    ```
4. Build the application:
    ```bash
    go build main.go
    ```

### Executing program

* After going through the installation, run the following command to execute the program
    ```bash
    go run main.go
    ```

## Input Data

The application expects a JSON file named `dataset.json` in the `dataset` directory containing shift data. The following JSON structure is required:

```JSON
[
    {
        "ShiftID": 123,
        "EmployeeID": 456,
        "StartTime": "1985-04-12T23:20:50.52Z",
        "EndTime": "1985-04-13T07:19:14.03Z"
    }
]
```

## Output

The application will generate a JSON file named `output.json` containing the summaries for each employee, their start of week, their regular hours, overtime hours for that week and any invalid shifts that week.

```JSON
[
    {
        "EmployeeID": 456,
        "StartOfWeek": "2021-08-22",
        "RegularHours": 20.56,
        "OvertimeHours": 0,
        "InvalidShifts": [
            123,
            234
        ]
    }
]
```

## Testing
* To run all the tests, run the following command
    ```bash
    go test ./...
    ```
    
## Wish List
Since this take-home challenge was time boxed for 3 hours or less, there are things I would have loved to spend more time on. They are as followed:

* Put a config file indicating where the dataset.json is located
* Write more tests cases to cover some edge cases, especially the one in the bonus
    * Cases like if values in JSON were missing
* Refactor the tests models since a few seems repetitive
* Refactor logic
    * After getting most of the code done, I realized that I could have optimized the sorting done in processshift.go could be more optimized for later. I could be adding an extra bit to the key with start of the week so I could reduce the number of time I filter in calculateemployeesummary.go
        * this would essentially eliminate the sorting on /calculateemployyeesummary.go#L17
        * this would reduce the if statement on /calculateemployyeesummary.go#L26 & /calculateemployyeesummary.go#L33
    * I feel lilke findInvalidShift function could be optimized but currently it is 
* Refactor some of the code for readability such as
    * Grouping shifts by employee ID into it's own function and writing tests
    * Part of the CalculateEmployeeSummary could be factored out into it's own function for readiability
        * i.e. Checking if we need a new entry for EmployeeSummary if it's a start of a new week

## Extra Notes
* Normally I would add `go vendor` files into the repo if this was not a challenge, but for file readability it was left out.
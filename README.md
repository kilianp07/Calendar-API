# Calendar API

The Calendar API is a Golang-based web service that provides calendar data in both JSON and HTML formats. It allows users to retrieve calendar information for a specific month and year.

## Usage

To use the Calendar API, you can make HTTP GET requests to the following endpoints:

### JSON Calendar Data

Retrieve calendar data in JSON format:

```
GET /json-calendar?month=<month>&year=<year>
```

- `month` (optional): The numeric representation of the month (1-12). If not provided, the current month will be used.
- `year` (optional): The year (e.g., 2023). If not provided, the current year will be used.

### HTML Calendar

Retrieve calendar data in HTML format:

```
GET /html-calendar?month=<month>&year=<year>
```

- `month` (optional): The numeric representation of the month (1-12). If not provided, the current month will be used.
- `year` (optional): The year (e.g., 2023). If not provided, the current year will be used.

Certainly! To indicate to users that they should download the project from the releases section, you can update the "Usage" section of your README as follows:


## Installation

To use the Calendar API, follow these steps:

1. Go to the [Releases](https://github.com/kilianp07/Calendar-API/releases) section of this repository.

2. Find the release version you want to use.

3. Download the executable for your operating system:
   - For Linux: Download `calendarapi-linux`
   - For Windows: Download `calendarapi-windows.exe`
   - For macOS: Download `calendarapi-macos`

4. Run the downloaded executable in your terminal or command prompt.

You can also build the project from the source code by cloning the repository and following the installation instructions in the "Installation" section below.

## Installation (from source)

If you prefer to build the project from source code, follow these steps:

1. Clone the repository to your local machine:

```
git clone https://github.com/kilianp07/Calendar-API.git
```
2. Run the API:

```
go run main.go
```

The API will start and be accessible at `http://localhost:8080`.


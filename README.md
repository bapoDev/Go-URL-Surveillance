# Go URL Sureillance

This project was created as in intro to Go. I used Gemini to correct some synchronisation errors (it basically told me about WaitGroups).
This project uses the `net/http` package as well as goroutines, all available with a base Go install.
I tried to make the output as readable as possible but if you have line wrapping on on your terminal, the error messages might wrap onto a new line since they are rather long.

## Usage

There are releases for most used architectures, otherwise go to the Build section.
Simply execute it in the terminal.

```bash
./service-monitor <text file with URLs>
```

## Build

```bash
git clone https://github.com/bapoDev/Go-URL-Surveillance
go build
./service-monitor <text file with URLs>
```

## What's next ?

- Flags for more customisation
- Worker pool to not fry your RAM and CPU
- Formatting in other files and outputs

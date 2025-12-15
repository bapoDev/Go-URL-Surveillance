# Go URL Sureillance

This project was created as in intro to Go. I used Gemini to correct some synchronisation errors (it basically told me about WaitGroups).
This project uses the `net/http` package as well as goroutines, all available with a base Go install.
I tried to make the output as readable as possible but if you have line wrapping on on your terminal, the error messages might wrap onto a new line since they are rather long.

## Installation & usage

> [!IMPORTANT]
> You need to create a file called `urls.txt` that contains exactly **ONE** url per line.

```bash
git clone https://github.com/bapoDev/Go-URL-Surveillance
go build
./service-monitor
```

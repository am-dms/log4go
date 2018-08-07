# log4go
a thin wrapper around the logrus for adding styles as per DMS (Daimler Mobility Services) guideline.
also it adds WithObject Sugar to easily add context with objects

## controling behaviour
there are two environment variable that control the behaviour
* **loglevel** the value of which should be one of DEBUG, INFO, WARN, ERROR. the defualt is INFO
* **logfmt** if you specify this as **text** then it format ourput as plain text, otherwise it formats output as json

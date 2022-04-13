# go-crud-k8s-demo
A very basic demo for a CRUD API built in golang and interfacing with a Postgresql DB

Originally, the intended purpose of this small demo app was to showcase a simple way to provision k8s resources on-demand for feature branches,
however the demo may be interesting beyond that use-case, so I've decided to make it available here.

Please note that the example is still WIP and the application currently only exposes two very simple GET endpoints and does not cache anything.

## Configuration Options
The easiest way to configure the application is to set the appropriate environment variables.

Available options:  
`APPLICATION_PORT` Port the application server binds to (default: 8080)  
`ENABLE_DEV` Set dev environment  
`ENABLE_TRACE` If set, logger is configured to print file and line it is called from
`DB_TIMEOUT` Database timeout in seconds  
`DB_HOST` Database Hostname
`DB_PORT` Database Port  
`DB_USER` Database User
`DB_PASS` Database Password  
`DB_NAME` Database Name
`DB_SSL_DISABLE` disable ssl mode 

## Testing locally
You can get started by navigating to the migrations folder and running the `run_migration.sh` script. It will download the test data, launch a postgres container based on the official image and load the test data to the database  

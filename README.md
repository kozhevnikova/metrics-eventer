# Metrics Eventer

Simple solution of metrics eventer.

## Task description

There is a Postgresql database containes tables: users, devices, device_metrics, device_alerts. 

* users: users' information
* devices: devices' information linked to users' data
* device_metrics: devices' metrics
* device_alerts: messages about errors from devices

There are about 10000 devices. 
Devices send metrics every 5 seconds and they are able to send several count of metrics at once.

Need to create program, which will monitor devices' metrics.
If metric value is more than critical value (need to write this value in config) then send email about that.
Also, keep in device_alerts table and Redis this message. 

Redis should keep only last message for device.
If message exists, need to rewrite to newest. 

## Solution

There are two folders in repository:
* **Generation** containes code for generation of metrics.
* **Notification** containes code for eventer. 
You need to configure all of them.

### Environment
Use lxd containers. [For more information](https://linuxcontainers.org/lxd/getting-started-cli/) how to install and use lxd. 
Install [postgresql](https://wiki.postgresql.org/wiki/Detailed_installation_guides) and [redis](https://redis.io/topics/quickstart) there. 

### Database
Create postgresql database. 
```
CREATE DATABASE metrics
```
The Notification folder containes script.sql file. There are 4 tables.
You need to create tables and insert data in users and device. 
Use:
```
INSERT INTO users(id,name,email) VALUES (1,'','');
```
For generation devices use function from 41 line to 46 line of file script.sql.

### Config
First, go to *Generation* folder. Add info in config.toml. 
```
[database]
      user="user of database"
      password="some password"
      name="name of database"
      host="ip"
```
*Notice: ip address of container.*

Second, go to *Notification* folder. Add info in config.toml.

* **[database]** - data for sql database
* **[metrics]** - critical value for metrics (must be integer)

* **[redis]** - redis info 
```
address="ip:port"
database=(number)
```
*Notice: ip address of container and port of redis.*

* **[mail]** - info for sending email
```
addressFrom="email address"
password="password for email"
servername="smtp server of email"
port=(port number of smtp server for outgoing messages)
```
*Notice: servername of mail service (for example, smtp.yandex.ru).*

## Build project
For building and starting code use:
```
go build && ./generation
go build && ./notification
```

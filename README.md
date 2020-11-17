<img src='https://user-images.githubusercontent.com/1423657/50455638-a8c41580-094f-11e9-8b43-dd0a9ae0f622.png' width=100>

# cLoki-go

### like Loki, but for Clickhouse.

Super experimental, fully functional [Loki](https://github.com/grafana/loki) API emulator in GO for [Clickhouse](https://clickhouse.yandex/)<br/>
APIs are compatible with [Grafana Explore](http://docs.grafana.org/features/explore/) and [paStash](https://github.com/sipcapture/paStash/wiki/Example:-Loki) for logs ingestion

:fire: *Beta Stage, Contributions Welcome! :octocat: Do not use this for anything serious.. yet!*

![ezgif com-optimize 15](https://user-images.githubusercontent.com/1423657/50496835-404e6480-0a33-11e9-87a4-aebb71a668a7.gif)

### Project Background

The *Loki API* and its Grafana native integration are brilliant, simple and appealing - but we just love **Clickhouse**. 

**cLoki** implements the same API functionality as Loki in GO, sitting on top of **Clickhouse** tables and relying on its *columnar search and insert performance alongside solid distribuion and clustering capabilities* for stored data. Just like Loki, cLoki does not parse or index incoming logs, but rather groups log streams using the same label system as Prometheus. 

<img src="https://user-images.githubusercontent.com/1423657/54091852-5ce91000-4385-11e9-849d-998c1e5d3243.png" width=700 />

*The current purpose of this project is to research and understand inner aspects of the original implementation.*


### Instructions

#### Requirements
* golang 1.13+

#### Installation
To get dependencies and compile the latest cloki-webapp from source, use the following commands:
```
make modules
make all
```

### Configuration
Before using the application, configure all database parameters using the example configuration file:
```
clokiapp_config.json
```

#### Usage
##### Command Help
```
./cloki-go -h
```
##### Custom Config in `/etc`
```
./cloki-go -config-path /etc
```

##### Initialization
To initialize the database and tables required by the application use the following commands:
```
./cloki-go --initialize_db
```

###### Create User
```
./cloki-webapp -create-cloki-user -database-root=default -database-host=localhost -database-root-password=postgres
```

###### Create CLOKI DBs
```
./cloki-go -create-cloki-db -database-root=default -database-host=localhost -database-root-password=password -database-cloki-user=cloki_user

```

<!--
###### Save it or edit the webapp_config.json manualy
```
./cloki-go -save-cloki-db-settings -database-host=localhost -database-cloki-config=cloki_config -database-cloki-user=cloki_user -database-cloki-password=cloki_password
```
-->

###### Create Table / Migration - connection data will be read from `webapp_config.json`
```
./cloki-go -create-table-db-config 
```

------------

#### Usage ENV
```
WEBAPPENV = config file extension "local" 
WEBAPPPATH - path for config
WEBAPPLOGPATH - path to the log dir
WEBAPPLOGNAME - prefix name of the log
```


----

#### License & Copyright
This project is released under the Apache License 2.0

#### Made by Humans
This Open-Source project is made possible by actual Humans without corporate sponsors, angels or patreons.<br>
If you use this software in production, please consider supporting its development with contributions or [donations](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=donation%40sipcapture%2eorg&lc=US&item_name=SIPCAPTURE&no_note=0&currency_code=EUR&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHostedGuest)

[![Donate](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=donation%40sipcapture%2eorg&lc=US&item_name=SIPCAPTURE&no_note=0&currency_code=EUR&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHostedGuest) 

###### (C) 2020 QXIP BV

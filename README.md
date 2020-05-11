# optresume_scraper

optresume scraper

## Installation

### Download Repository

```bash
git clone git@github.com:mavensingh/optresume_scraper.git
```

### Configuration

```bash
mv sample.env.sh env.sh
```

```bash
export DBHOST="DBHOST"
export DBUSER="DBUSER"
export DBPASS="DBPASSWORD"
export DBPORT="DBPORT"
```

### Build Software

```bash
go build -o app

OR

sudo go build -o app
```

### Run Software

```bash
source env.sh
```

```bash
./app
```

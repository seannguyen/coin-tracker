# coinbase-tracker
Tracker for cryptocurrencies on various exchanges

# Getting Started:
## System Depnedencies
```bash
sudo apt-get install python-dev python-pip gcc
```

## Python dependencies
```bash
pip install virtualenv
cd <REPO_DIR>
virtualenv env
source env/bin/active
pip install -r requirements.txt
```

## Setup crontab file
```
crontab -u seannguyen /var/www/coin-tracker/cron/crontab
```

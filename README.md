# gogcmd

Golang API to fetch GCMD Keywords and format them in a particular way.
Applicable for a specific use case only.

GCMD Keywords: https://earthdata.nasa.gov/earth-observation-data/find-data/idn/gcmd-keywords

## How to use

### If you have Go installed on your system
You can run the script as is. This also gives you the flexibility to generate only some categories by editing the 'concepts' variable.
```
git clone git@github.com:HarrySng/gogcmd.git
cd gogcmd
mkdir files terms
go run gcmd.go
cd terms
ls -l
```

### If you do not have Go installed on your system
Execute the binary based on your OS and system architecture.
Windows 64 bit: gogcmd_windows_amd64.exe
Linux 64 bit: gogcmd_linux_amd64
```
git clone git@github.com:HarrySng/gogcmd.git
cd gogcmd
mkdir files terms
gogcmd_linux_amd64 # Change binary based on your system
cd terms
ls -l
```

## How It Works
* Keywords are first downloaded in csv format from the static page here: https://gcmd.earthdata.nasa.gov/static/kms/
* Each category is stores as a separate file in ./files/ by its name.
* Each record in the csv file is then converted into this specific format:
```
Raw record: 
"EARTH SCIENCE SERVICES","DATA ANALYSIS AND VISUALIZATION","CALIBRATION/VALIDATION","CALIBRATION","","","","ecf29317-bd5e-447b-b911-f8bfb153c83b"

Formatted record: 
ecf29317-bd5e-447b-b911-f8bfb153c83b
EARTH SCIENCE SERVICES > DATA ANALYSIS AND VISUALIZATION > CALIBRATION/VALIDATION > CALIBRATION
```
* The formatted records are written to txt files in ./terms/ by category names.

## Contact
For any queries, please reach out to me at harrysng@outlook.com
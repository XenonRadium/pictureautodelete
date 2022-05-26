# pictureautodelete

# To test, run Autodelete.exe (Windows)|Autodelete (Linnux), or use command go build main.go, go run main.go
# To deploy, please delete the kiplebox folder containing dummy pictures, and change path in config.yml according to system structure.

## Function
The program takes in a user input for date in the format of YYYY-MM-DD.
All pictures contained in directory specified in the path specified in config.yml is deleted if 2 requirements are met:
      -   Their current folder date is before the date specified by the user.
      -   The pictures' name starts with 'http' (uploaded to cloud database.)
The directories processed by the program is shown in the command line, as well as whether they were preserved or deleted.

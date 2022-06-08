# minnie-test

**This work depends on PostgreSQL being installed, PostgreSQL server being turned on, Go lang and Node to be installed on your Mac / Linux machine.**

Now run below command to clone this project into your directory
*git clone https://github.com/soberservicesguy/minnie-tests*

Now rename .env-example file into .env and fill your own environment variables into it before proceeding. 

To install project dependencies, create database, create tables into database, import bulk data from csv to database, run frontend and backend tests, just run the following command in project root directory
*make*

To create a fresh build and run it, perform below command
*make run-fresh-build*

To perform only tests again, run below command
*make test*
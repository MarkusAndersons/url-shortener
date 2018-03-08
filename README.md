# url-shortener
[![Build Status](https://travis-ci.org/MarkusAndersons/url-shortener.svg?branch=master)](https://travis-ci.org/MarkusAndersons/url-shortener)

A simple web app to produce and manage shortened URLs

This app is just a fun way for me to learn some Go.

## Configuration
The app now requires a Postgres database with the connection url stored in the environment variable DATABASE_URL (as provided by Heroku)

**NOTE** Currently there is not sufficient validation so doing a "; drop tables" will break things!

## TODO
- Add further validation to given urls
- Add authentication to create url
- Add concurrency safety to accessing the data store
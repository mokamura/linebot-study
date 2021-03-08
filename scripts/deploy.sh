#!/bin/bash

set -e

cd $(dirname $(readlink $0 || echo $0)); cd ../

heroku container:login
heroku container:push web
heroku container:release web
heroku open
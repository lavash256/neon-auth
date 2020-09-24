#!/bin/sh

echo "Hello, it is neon auth"
#It is using only docker container
set -e

if [ "$1" = 'migrate' ]; then
    echo "Run migrate command"
    exec neon-migrate
fi
if [ "$1" = 'run' ]; then
    echo "Run neon-auth service"
    exec neon-auth
fi
echo "Parameters were not passed"
exec "$@"
